//go:build ignore

// generate_reference.go generates a CSV reference of all public functions.
// Usage: go run tools/generate_reference.go
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
}

func main() {
	fset := token.NewFileSet()

	// Parse all Go files with comments
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing directory: %v\n", err)
		os.Exit(1)
	}

	var publicFuncs []funcInfo
	var tests, examples, fuzzTests, benchmarks []funcInfo

	for _, pkg := range pkgs {
		for filename, file := range pkg.Files {
			isTest := strings.HasSuffix(filename, "_test.go")

			ast.Inspect(file, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if !ok || fn.Recv != nil { // Skip methods
					return true
				}

				name := fn.Name.Name
				if !unicode.IsUpper(rune(name[0])) {
					return true // Skip unexported functions
				}

				pos := fset.Position(fn.Pos())
				info := funcInfo{
					name: name,
					file: filepath.Base(pos.Filename),
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
					info.signature = formatSignature(fn)
					info.doc = extractFirstSentence(fn.Doc)
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

	// Build lookup maps
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
			fmt.Sprintf("%s:%d", f.file, f.line),
			fileToGroup[f.file],
			findLoc(f.name, exampleMap),
			findLoc(f.name, testMap),
			findLoc(f.name, fuzzMap),
			findLoc(f.name, benchmarkMap),
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
func findLoc(funcName string, testMap map[string]funcInfo) string {
	if t, ok := testMap[funcName]; ok {
		return fmt.Sprintf("%s:%d", t.file, t.line)
	}
	if t, ok := testMap["Inflect"+funcName]; ok {
		return fmt.Sprintf("%s:%d", t.file, t.line)
	}
	return ""
}
