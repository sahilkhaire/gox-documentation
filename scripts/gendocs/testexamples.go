package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// loadTestExamples scans gox *_test.go files and extracts usage snippets per pkg.Symbol.
func loadTestExamples(goxRoot string) map[string]string {
	out := make(map[string]string)
	for _, group := range packageGroups {
		for _, pkgName := range group.packages {
			dir := filepath.Join(goxRoot, pkgName)
			matches, _ := filepath.Glob(filepath.Join(dir, "*_test.go"))
			for _, file := range matches {
				extractTestFileExamples(file, pkgName, out)
			}
		}
	}
	return out
}

func extractTestFileExamples(file, pkgName string, out map[string]string) {
	src, err := os.ReadFile(file)
	if err != nil {
		return
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, src, 0)
	if err != nil {
		return
	}
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil || !strings.HasPrefix(fn.Name.Name, "Test") {
			continue
		}
		symbol := testNameToSymbol(fn.Name.Name)
		if symbol == "" {
			symbol = guessSymbolFromBody(fn.Body, pkgName)
		}
		if symbol == "" {
			continue
		}
		key := pkgName + "." + symbol
		if _, exists := out[key]; exists {
			continue
		}
		snippet := extractSnippet(fset, fn.Body, pkgName)
		if snippet == "" {
			continue
		}
		out[key] = snippet
	}
}

func testNameToSymbol(testName string) string {
	name := strings.TrimPrefix(testName, "Test")
	if name == "" {
		return ""
	}
	if strings.Contains(name, "And") {
		parts := strings.Split(name, "And")
		name = parts[0]
	}
	runes := []rune(name)
	if len(runes) == 0 {
		return ""
	}
	return string(unicode.ToUpper(runes[0])) + string(runes[1:])
}

func guessSymbolFromBody(body *ast.BlockStmt, pkgName string) string {
	var found string
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		switch fun := call.Fun.(type) {
		case *ast.Ident:
			if fun.IsExported() {
				found = fun.Name
				return false
			}
		case *ast.SelectorExpr:
			if id, ok := fun.X.(*ast.Ident); ok && id.Name == pkgName && fun.Sel.IsExported() {
				found = fun.Sel.Name
				return false
			}
		}
		return true
	})
	return found
}

func extractSnippet(fset *token.FileSet, body *ast.BlockStmt, pkgName string) string {
	var lines []string
	for _, stmt := range body.List {
		if isTestBoilerplate(stmt) || isTestAssertion(stmt) {
			continue
		}
		var buf strings.Builder
		if err := printer.Fprint(&buf, fset, stmt); err != nil {
			continue
		}
		line := strings.TrimSpace(buf.String())
		if line == "" || strings.Contains(line, "t.Fatal") || strings.Contains(line, "t.Error") {
			continue
		}
		lines = append(lines, line)
		if len(lines) >= 6 {
			break
		}
	}
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\n")
}

func isTestAssertion(stmt ast.Stmt) bool {
	if _, ok := stmt.(*ast.IfStmt); ok {
		return true
	}
	if _, ok := stmt.(*ast.ForStmt); ok {
		return true
	}
	return false
}

func isTestBoilerplate(stmt ast.Stmt) bool {
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return false
	}
	call, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := sel.X.(*ast.Ident)
	return ok && id.Name == "t" && (sel.Sel.Name == "Fatal" || sel.Sel.Name == "Fatalf" || sel.Sel.Name == "Error")
}
