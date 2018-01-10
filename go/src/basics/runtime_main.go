package main

import (
	"fmt"
	"runtime"
	//"runtime/debug"
)

func main() {
	fmt.Printf("gomaxprocs: %d\n", runtime.GOMAXPROCS)
	fmt.Printf("goos: %s\n", runtime.GOOS)
	fmt.Printf("goarch: %s\n", runtime.GOARCH)
	fmt.Printf("goroot: %s\n", runtime.GOROOT())
	fmt.Printf("compiler: %s\n", runtime.Compiler)
}
