package api

import "C"

//export Add
func Add(a, b int) int {
	return a + b
}

//export ReturnStruct
func ReturnStruct(a int, b string) (int, string) {
	// remember to deallocate in caller
	return a, b
}

//export ReturnStruct2
func ReturnStruct2(a int, b string) (int, *C.char) {
	return a, C.CString(b)
}
