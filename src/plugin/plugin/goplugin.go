// Plugin mode must be in main package

package main

import "fmt"

func ExportedForPlugin() {
	fmt.Println("in ExportedForPlugin func")
}
