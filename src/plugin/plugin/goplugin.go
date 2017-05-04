// Plugin mode must be in main package

package main

import "fmt"

//export ExportedForPlugin
func ExportedForPlugin() {
	fmt.Println("in ExportedForPlugin func")
}
