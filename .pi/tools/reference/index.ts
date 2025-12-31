import { Type } from "@sinclair/typebox";
import { StringEnum } from "@mariozechner/pi-ai";
import { Text } from "@mariozechner/pi-tui";
import type { CustomToolFactory } from "@mariozechner/pi-coding-agent";

interface FunctionRef {
  name: string;
  sig: string;
  desc: string;
  loc: string;
  group: string;
  example: string;
  test: string;
  fuzz: string;
  bench: string;
}

function parseCSV(csv: string): FunctionRef[] {
  const lines = csv.trim().split("\n");
  if (lines.length < 2) return [];

  const results: FunctionRef[] = [];

  for (let i = 1; i < lines.length; i++) {
    const line = lines[i];
    const fields: string[] = [];
    let current = "";
    let inQuotes = false;

    for (let j = 0; j < line.length; j++) {
      const char = line[j];
      if (char === '"') {
        inQuotes = !inQuotes;
      } else if (char === "," && !inQuotes) {
        fields.push(current);
        current = "";
      } else {
        current += char;
      }
    }
    fields.push(current);

    if (fields.length >= 9) {
      results.push({
        name: fields[0],
        sig: fields[1],
        desc: fields[2],
        loc: fields[3],
        group: fields[4],
        example: fields[5],
        test: fields[6],
        fuzz: fields[7],
        bench: fields[8],
      });
    }
  }

  return results;
}

function formatRef(ref: FunctionRef, verbose: boolean): string {
  if (verbose) {
    let out = `${ref.name}${ref.sig}\n`;
    out += `  ${ref.desc}\n`;
    out += `  Location: ${ref.loc}`;
    if (ref.group) out += ` | Group: ${ref.group}`;
    out += "\n";
    if (ref.test) out += `  Test: ${ref.test}\n`;
    if (ref.example) out += `  Example: ${ref.example}\n`;
    if (ref.fuzz) out += `  Fuzz: ${ref.fuzz}\n`;
    if (ref.bench) out += `  Bench: ${ref.bench}\n`;
    return out;
  }
  return `${ref.name}${ref.sig} - ${ref.desc.substring(0, 60)}${ref.desc.length > 60 ? "..." : ""} [${ref.loc}]`;
}

const factory: CustomToolFactory = (pi) => ({
  name: "reference",
  label: "Reference",
  description:
    "Get a map of the go-inflect codebase. Lists all public functions with signatures, descriptions, locations, and test coverage. Use 'filter' to search by function name or group. Use 'group' to list functions in a specific group. Use 'groups' format to see available groups.",
  parameters: Type.Object({
    filter: Type.Optional(
      Type.String({
        description:
          "Filter functions by name (case-insensitive substring match)",
      })
    ),
    group: Type.Optional(
      Type.String({
        description:
          "Filter by group name (e.g., 'nouns', 'verbs', 'numbers', 'formatting')",
      })
    ),
    format: Type.Optional(
      StringEnum(["brief", "verbose", "groups"] as const, {
        description:
          "Output format: 'brief' (default), 'verbose' (full details), 'groups' (list available groups)",
      })
    ),
  }),

  async execute(toolCallId, args, signal) {
    const params = args as {
      filter?: string;
      group?: string;
      format?: "brief" | "verbose" | "groups";
    };

    const result = await pi.exec("make", ["reference"], {
      signal,
      timeout: 30000,
    });

    if (result.code !== 0) {
      return {
        content: [
          {
            type: "text",
            text: `Error running 'make reference': ${result.stderr}`,
          },
        ],
        details: { error: result.stderr },
      };
    }

    // Extract CSV from output (skip the "go run..." line)
    const lines = result.stdout.split("\n");
    const csvStart = lines.findIndex((l) => l.startsWith("name,sig,"));
    if (csvStart === -1) {
      return {
        content: [{ type: "text", text: "Could not parse reference output" }],
        details: { error: "parse_error" },
      };
    }
    const csv = lines.slice(csvStart).join("\n");
    const refs = parseCSV(csv);

    // Handle 'groups' format
    if (params.format === "groups") {
      const groups = new Map<string, number>();
      for (const ref of refs) {
        const g = ref.group || "(none)";
        groups.set(g, (groups.get(g) || 0) + 1);
      }
      const sorted = [...groups.entries()].sort((a, b) =>
        a[0].localeCompare(b[0])
      );
      const groupList = sorted
        .map(([name, count]) => `${name}: ${count} functions`)
        .join("\n");
      return {
        content: [
          {
            type: "text",
            text: `Available groups:\n${groupList}\n\nTotal: ${refs.length} functions`,
          },
        ],
        details: { groups: Object.fromEntries(groups), total: refs.length },
      };
    }

    // Apply filters
    let filtered = refs;
    if (params.filter) {
      const f = params.filter.toLowerCase();
      filtered = filtered.filter(
        (r) =>
          r.name.toLowerCase().includes(f) ||
          r.desc.toLowerCase().includes(f)
      );
    }
    if (params.group) {
      const g = params.group.toLowerCase();
      filtered = filtered.filter((r) => r.group.toLowerCase() === g);
    }

    const verbose = params.format === "verbose";
    const output = filtered.map((r) => formatRef(r, verbose)).join("\n");

    const summary =
      filtered.length === refs.length
        ? `${refs.length} functions`
        : `${filtered.length} of ${refs.length} functions`;

    return {
      content: [{ type: "text", text: `${summary}\n\n${output}` }],
      details: { count: filtered.length, total: refs.length },
    };
  },

  renderCall(args, theme) {
    const params = args as {
      filter?: string;
      group?: string;
      format?: string;
    };
    let text = theme.fg("toolTitle", theme.bold("reference"));
    if (params.format) text += ` ${theme.fg("muted", params.format)}`;
    if (params.group) text += ` ${theme.fg("accent", `group:${params.group}`)}`;
    if (params.filter)
      text += ` ${theme.fg("dim", `filter:"${params.filter}"`)}`;
    return new Text(text, 0, 0);
  },

  renderResult(result, { expanded }, theme) {
    const details = result.details as {
      count?: number;
      total?: number;
      groups?: Record<string, number>;
      error?: string;
    };

    if (details?.error) {
      return new Text(theme.fg("error", `Error: ${details.error}`), 0, 0);
    }

    if (details?.groups) {
      let text = theme.fg("success", "✓ ") + theme.fg("muted", "Groups:");
      if (expanded) {
        for (const [name, count] of Object.entries(details.groups).sort()) {
          text += `\n  ${theme.fg("accent", name)}: ${count}`;
        }
      } else {
        text += ` ${Object.keys(details.groups).length} groups, ${details.total} total functions`;
      }
      return new Text(text, 0, 0);
    }

    let text =
      theme.fg("success", "✓ ") +
      theme.fg("muted", `${details?.count} of ${details?.total} functions`);
    return new Text(text, 0, 0);
  },
});

export default factory;
