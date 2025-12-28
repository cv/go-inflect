//go:build ignore

// gen-reference.go generates a CSV reference of all public functions.
// Usage: go run tools/gen-reference.go
package main

import (
	"encoding/csv"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

// funcInfo holds parsed function info.
type funcInfo struct {
	name      string
	file      string
	line      int
	signature string
	doc       string
}

// fileToGroup maps source files to semantic groups.
var fileToGroup = map[string]string{
	"plural.go":        "nouns",
	"singular.go":      "nouns",
	"article.go":       "articles",
	"adjective.go":     "adjectives",
	"adverb.go":        "adverbs",
	"verbs.go":         "verbs",
	"participle.go":    "verbs",
	"past_tense.go":    "verbs",
	"number.go":        "numbers",
	"ordinal.go":       "numbers",
	"fraction.go":      "numbers",
	"currency.go":      "numbers",
	"counting.go":      "numbers",
	"join.go":          "formatting",
	"case.go":          "formatting",
	"possessive.go":    "formatting",
	"compare.go":       "comparison",
	"classical.go":     "classical",
	"custom.go":        "customization",
	"gender.go":        "gender",
	"rails.go":         "rails",
	"util.go":          "utility",
	"inflect_funcs.go": "inflection",
	"inflect.go":       "inflection",
	"pronouns.go":      "pronouns",
	"engine.go":        "engine",
}

func main() {
	fset := token.NewFileSet()

	// Parse internal/inflect to find real implementations
	internalDir := "internal/inflect"
	pkgs, err := parser.ParseDir(fset, internalDir, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", internalDir, err)
		os.Exit(1)
	}

	// Collect Engine methods and package-level functions with their locations
	implementations := make(map[string]funcInfo)
	var tests, examples, fuzzTests, benchmarks []funcInfo

	for _, pkg := range pkgs {
		for filename, file := range pkg.Files {
			isTest := strings.HasSuffix(filename, "_test.go")

			ast.Inspect(file, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if !ok {
					return true
				}

				name := fn.Name.Name
				if !unicode.IsUpper(rune(name[0])) {
					return true // Skip unexported functions
				}

				pos := fset.Position(fn.Pos())
				baseFile := filepath.Base(pos.Filename)
				info := funcInfo{
					name: name,
					file: baseFile,
					line: pos.Line,
				}

				if isTest {
					switch {
					case strings.HasPrefix(name, "Example"):
						examples = append(examples, info)
					case strings.HasPrefix(name, "Fuzz"):
						fuzzTests = append(fuzzTests, info)
					case strings.HasPrefix(name, "Benchmark"):
						benchmarks = append(benchmarks, info)
					case strings.HasPrefix(name, "Test"):
						tests = append(tests, info)
					}
				} else {
					// Check if this is a method on *Engine
					if fn.Recv != nil && len(fn.Recv.List) > 0 {
						recvType := exprToString(fn.Recv.List[0].Type)
						if recvType == "*Engine" {
							info.signature = formatSignature(fn)
							info.doc = extractFirstSentence(fn.Doc)
							implementations[name] = info
						}
					} else if fn.Recv == nil {
						// Package-level function
						info.signature = formatSignature(fn)
						info.doc = extractFirstSentence(fn.Doc)
						// Only add if not already present (prefer Engine methods)
						if _, exists := implementations[name]; !exists {
							implementations[name] = info
						}
					}
				}

				return true
			})
		}
	}

	// Now parse the root package to get public API functions
	rootPkgs, err := parser.ParseDir(fset, ".", func(fi os.FileInfo) bool {
		// Skip test files
		name := fi.Name()
		return !strings.HasSuffix(name, "_test.go")
	}, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing root: %v\n", err)
		os.Exit(1)
	}

	var publicFuncs []funcInfo

	for _, pkg := range rootPkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if !ok || fn.Recv != nil { // Skip methods
					return true
				}

				name := fn.Name.Name
				if !unicode.IsUpper(rune(name[0])) {
					return true // Skip unexported functions
				}

				// Skip special functions that don't map to implementations
				if name == "DefaultEngine" || name == "NewEngine" || name == "GetPossessiveStyle" {
					return true
				}

				// Look up the implementation to get real location
				if impl, ok := implementations[name]; ok {
					info := funcInfo{
						name:      name,
						file:      impl.file,
						line:      impl.line,
						signature: impl.signature,
						doc:       impl.doc,
					}
					publicFuncs = append(publicFuncs, info)
				}

				return true
			})
		}
	}

	// Sort public functions by name
	sort.Slice(publicFuncs, func(i, j int) bool {
		return publicFuncs[i].name < publicFuncs[j].name
	})

	// Build lookup maps for tests (use internal test files)
	testMap := buildMap(tests, "Test")
	exampleMap := buildMap(examples, "Example")
	fuzzMap := buildMap(fuzzTests, "Fuzz")
	benchmarkMap := buildMap(benchmarks, "Benchmark")

	// Output CSV
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"name", "sig", "desc", "loc", "group", "example", "test", "fuzz", "bench"})

	for _, f := range publicFuncs {
		w.Write([]string{
			f.name,
			f.signature,
			f.doc,
			fmt.Sprintf("internal/inflect/%s:%d", f.file, f.line),
			fileToGroup[f.file],
			findLoc(f.name, exampleMap, internalDir),
			findLoc(f.name, testMap, internalDir),
			findLoc(f.name, fuzzMap, internalDir),
			findLoc(f.name, benchmarkMap, internalDir),
		})
	}
	w.Flush()
}

// formatSignature returns a function signature like "(word string) string".
func formatSignature(fn *ast.FuncDecl) string {
	var params []string
	if fn.Type.Params != nil {
		for _, field := range fn.Type.Params.List {
			typeStr := exprToString(field.Type)
			if len(field.Names) == 0 {
				params = append(params, typeStr)
			} else {
				for _, name := range field.Names {
					params = append(params, name.Name+" "+typeStr)
				}
			}
		}
	}

	var results []string
	if fn.Type.Results != nil {
		for _, field := range fn.Type.Results.List {
			typeStr := exprToString(field.Type)
			if len(field.Names) == 0 {
				results = append(results, typeStr)
			} else {
				for _, name := range field.Names {
					results = append(results, name.Name+" "+typeStr)
				}
			}
		}
	}

	sig := "(" + strings.Join(params, ", ") + ")"
	if len(results) == 1 {
		sig += " " + results[0]
	} else if len(results) > 1 {
		sig += " (" + strings.Join(results, ", ") + ")"
	}
	return sig
}

// exprToString converts an AST expression to a string representation.
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + exprToString(t.Elt)
		}
		return "[...]" + exprToString(t.Elt)
	case *ast.MapType:
		return "map[" + exprToString(t.Key) + "]" + exprToString(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.Ellipsis:
		return "..." + exprToString(t.Elt)
	case *ast.FuncType:
		return "func(...)"
	default:
		return fmt.Sprintf("%T", expr)
	}
}

// extractFirstSentence returns the first sentence from a doc comment.
func extractFirstSentence(doc *ast.CommentGroup) string {
	if doc == nil {
		return ""
	}

	text := doc.Text()
	text = strings.TrimSpace(text)

	// Find first sentence ending
	for i, r := range text {
		if r == '.' || r == '\n' {
			sentence := strings.TrimSpace(text[:i+1])
			// Remove trailing period for cleaner output
			sentence = strings.TrimSuffix(sentence, ".")
			return sentence
		}
	}
	return text
}

// buildMap creates a mapping from function names to their info.
func buildMap(funcs []funcInfo, prefix string) map[string]funcInfo {
	m := make(map[string]funcInfo)
	for _, f := range funcs {
		baseName := strings.TrimPrefix(f.name, prefix)
		if _, exists := m[baseName]; !exists {
			m[baseName] = f
		}
	}
	return m
}

// findLoc returns "file:line" for a test, or empty string if not found.
func findLoc(funcName string, testMap map[string]funcInfo, baseDir string) string {
	if t, ok := testMap[funcName]; ok {
		return fmt.Sprintf("%s/%s:%d", baseDir, t.file, t.line)
	}
	if t, ok := testMap["Inflect"+funcName]; ok {
		return fmt.Sprintf("%s/%s:%d", baseDir, t.file, t.line)
	}
	return ""
}
