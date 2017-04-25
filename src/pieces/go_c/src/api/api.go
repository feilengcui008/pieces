package api

/*
#include <stdio.h>
#include <stdlib.h>
#include "api.h"
//#cgo LDFLAGS: -L. -lapi
*/
import "C"
import "unsafe"
import "fmt"

func ApiHello() {
	C.printHello()
}

func StringPrint(s string) {
	cs := C.CString(s)
	C.printString(cs)
	C.puts(cs)
	// C.printf(cs) // error, cgo不支持可变参数
	defer C.free(unsafe.Pointer(cs))
}

func TestCStruct() {
	var s C.struct_CStruct
	s.a = C.int(12)
	s._type = C.float(12.12)
	C.printCStruct(s)
}

func TestTestStruct() {
	var s C.TestStruct
	s.a = C.int(12)
	s._type = C.float(12.12)
	C.printTestStruct(s)
}

func AllocTestStruct() {
	var p *C.TestStruct
	// C运行时分配内存
	p = C.allocTestStruct()
	fmt.Println(p.a)
	// 记得释放
	C.free(unsafe.Pointer(p))
}

func TestVoidPointer() {
	//var buf [1]byte
	p := C.malloc(C.size_t(10))
	defer C.free(p)
	C.testVoidPointer(p, 10)
	// GoBytes和GoString会在Go运行时重新分配内存拷贝数据，gc负责回收
	fmt.Println(string(C.GoBytes(p, 10)))
	fmt.Println(C.GoString((*C.char)(p)))

}

func GetStructInfo() {
	var p unsafe.Pointer
	C.setStruct(&p)
	C.printStruct(p)
	C.freeStruct(p)
}
