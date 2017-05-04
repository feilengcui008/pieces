// Go1.8 only support Go style api plugin(.so compiled from go build -buildmode=plugin)
// do not support C stype api plugin(.so compiled from gcc)

package main

import (
	"fmt"
	"plugin"
)

func testUseGoPlugin() {
	p, err := plugin.Open("plugin/goplugin.so")
	fmt.Printf("%#v with err %v\n", p, err)
	fn, err := p.Lookup("ExportedForPlugin")
	fmt.Printf("%T %v\n", fn, fn)
	fn.(func())()
}

// will fail to load c shared library
func testUseCPlugin() {
	p, err := plugin.Open("plugin/cplugin.so")
	fmt.Printf("%#v with err %v\n", p, err)
	fn, err := p.Lookup("Myprintf")
	fmt.Printf("%T %v\n", fn, fn)
}

func main() {
	//testUseGoPlugin()
	fmt.Println("=======")
	testUseCPlugin() // fail don't know why
}
