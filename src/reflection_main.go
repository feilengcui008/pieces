package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

type ReaderInter interface {
	Read([]byte) (int, error)
}

type MyReader struct {
}

func (r MyReader) Read(b []byte) (n int, err error) {
	return
}

type OneReader struct {
}

func (r OneReader) Read(b []byte) (n int, err error) {
	return
}

func PrintLine(n int) {
	fmt.Println(strings.Repeat("*", n))
}

func main() {
	var r io.Reader = MyReader{}

	typ := reflect.TypeOf(r)
	fmt.Printf("typeOf MyReader: %v\n", typ)
	fmt.Printf("valueOf MyReader: %v\n", reflect.ValueOf(r))
	fmt.Printf("value.Type: %v\n", reflect.ValueOf(r).Type())

	PrintLine(10)

	var r1 io.Reader
	fmt.Printf("typeOf nil interface: %v\n", reflect.TypeOf(r1))
	fmt.Printf("valueOf nil interface: %v\n", reflect.ValueOf(r1))
	//fmt.Printf("valueOf nil interface: %v\n", reflect.ValueOf(r1).Type())

	PrintLine(10)

	var r2 io.Reader = OneReader{}
	fmt.Printf("typeOf OneReader: %v\n", reflect.TypeOf(r2))
	fmt.Printf("valueOf OneReader: %v\n", reflect.ValueOf(r2))
	fmt.Printf("value.Type: %v\n", reflect.ValueOf(r2).Type())

	PrintLine(10)

	var r3 io.Reader = &OneReader{}
	fmt.Printf("typeOf &OneReader: %#v\n", reflect.TypeOf(r3))
	fmt.Printf("valueOf &OneReader: %v\n", reflect.ValueOf(r3))
	fmt.Printf("value.Type: %v\n", reflect.ValueOf(r3).Type())

	PrintLine(10)

	var r4 interface{} = OneReader{}
	fmt.Printf("typeOf interface{} by OneReader: %#v\n", reflect.TypeOf(r4))
	fmt.Printf("valueOf interface{} by OneReader: %v\n", reflect.ValueOf(r4))
	fmt.Printf("value.Type interface{} by OneReader: %v\n", reflect.ValueOf(r4).Type())

	PrintLine(10)

	sl := make([]int, 5)
	fmt.Printf("%v\n", reflect.TypeOf(sl))

	ch := make(chan int)
	fmt.Printf("%v\n", reflect.TypeOf(ch))

	m := make(map[int]string)
	fmt.Printf("%v\n", reflect.TypeOf(m))

	var a interface{} = m
	fmt.Printf("%v\n", reflect.TypeOf(a))

}
