// Go1.8 only support Go style api plugin(.so compiled from go build -buildmode=plugin)
// do not support C stype api plugin(.so compiled from gcc)

package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("plugin/goplugin.so")
	fmt.Printf("%#v with err %v\n", p, err)
	fn, err := p.Lookup("ExportedForPlugin")
	fmt.Printf("%T %v\n", fn, fn)
	fn.(func())()
}
