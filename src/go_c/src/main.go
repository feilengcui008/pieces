package main

import "fmt"
import "rand"
import "api"

func main() {
	fmt.Println(rand.Random())
	api.ApiHello()
	api.StringPrint("ffffff")
	api.TestCStruct()
	api.TestTestStruct()
	api.AllocTestStruct()
	api.TestVoidPointer()
	//api.GetStructInfo()
}
