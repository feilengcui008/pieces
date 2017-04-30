// this is the Doc node
//

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

// func ParseFile commment
//
func ParseFile(filename string) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	return f, fset, err
}

// func main comment
//
func main() {
	f, fset, _ := ParseFile("parser_main.go")

	fmt.Printf("===== ast.File ======\n")
	fmt.Printf("%+v\n", f)

	fmt.Printf("===== fset ======\n")
	fmt.Printf("%+v\n", fset)

	fmt.Printf("===== whole file ======\n")
	printer.Fprint(os.Stdout, fset, f)
	fmt.Println()

	// 文档节点列表
	//
	fmt.Printf("==== Doc ======\n")
	printer.Fprint(os.Stdout, fset, f.Doc)

	// package节点
	//
	fmt.Printf("===== package node ====\n")
	printer.Fprint(os.Stdout, fset, f.Package)
	fmt.Println()

	// package name节点
	fmt.Printf("===== package name node ====\n")
	printer.Fprint(os.Stdout, fset, f.Name)
	fmt.Println()

	// decls节点列表
	fmt.Printf("===== decls of file ======\n")
	for i, n := range f.Decls {
		fmt.Printf("--- decl %d ---\n", i)
		printer.Fprint(os.Stdout, fset, n)
		fmt.Println()
	}

	// imports节点
	fmt.Printf("===== imports of file ======\n")
	for i, n := range f.Imports {
		fmt.Printf("--- import %d ---\n", i)
		printer.Fprint(os.Stdout, fset, n)
		fmt.Println()
	}

	// 未解析的符号节点列表
	fmt.Printf("===== unresolved identifiers ======\n")
	for i, n := range f.Unresolved {
		fmt.Printf("--- unresolved %d ---\n", i)
		printer.Fprint(os.Stdout, fset, n)
		fmt.Println()
	}

	// 注释节点列表
	fmt.Printf("===== comment group ======\n")
	for i, n := range f.Comments {
		fmt.Printf("--- comment %d ---\n", i)
		printer.Fprint(os.Stdout, fset, n)
		fmt.Println()
	}

}
