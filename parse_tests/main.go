package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	// Flag to control new lines between test functions
	firstTest := true

	for _, decl := range f.Decls {
		if funDecl, ok := decl.(*ast.FuncDecl); ok {
			if len(funDecl.Name.Name) > 4 && funDecl.Name.Name[:4] == "Test" {
				// Add a new line for all but the first test function
				if !firstTest {
					fmt.Println()
				}
				firstTest = false

				fmt.Println("Test function:", funDecl.Name.Name)

				for _, stmt := range funDecl.Body.List {
					if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
						if callExpr, ok := exprStmt.X.(*ast.CallExpr); ok {
							if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
								if selExpr.Sel.Name == "Run" {
									if lit, ok := callExpr.Args[0].(*ast.BasicLit); ok {
										fmt.Println("  -> Sub-test:", lit.Value)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
