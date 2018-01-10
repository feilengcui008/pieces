package main

import (
	"fmt"
	"unsafe"
)

type ZeroStructTest struct {
}

// 32 byte len
type ContentStructTest struct {
	n int    // 8
	s string // 16
	p *int   // 8
}

// go has only pass-by-value
// all types are passed by value

// string is a struct with two words
func stringFunc(s string) {
	fmt.Printf("%p\n", &s)
	fmt.Println(unsafe.Sizeof(s)) // 16
}

// array is a block memory with len(arr) * sizeof(ele)
func arrayFunc(arr [4]int) {
	fmt.Printf("%p\n", &arr)
	fmt.Println(unsafe.Sizeof(arr)) // 32
}

// slice is a struct with size 3 word
func sliceFunc(sl []int) {
	fmt.Printf("%p\n", &sl)
	fmt.Println(unsafe.Sizeof(sl)) // 24
}

// map is a pointer with size a word
func mapFunc(m map[int]string) {
	fmt.Printf("%p\n", &m)
	fmt.Println(unsafe.Sizeof(m)) // 8
}

// chan is a pointer with size a word
func chanFunc(ch chan int) {
	fmt.Printf("%p\n", &ch)
	fmt.Println(unsafe.Sizeof(ch)) // 8
}

// pointer is a pointer with size a word
func pointerFunc(p *int) {
	fmt.Printf("%p\n", &p)
	fmt.Println(unsafe.Sizeof(p)) // 8
}

// func is a pointer with size a word
func funcFunc(fn func()) {
	fmt.Printf("%p\n", &fn)
	fmt.Println(unsafe.Sizeof(fn)) // 8
}

// struct is a struct with size of sum of all elem size
func contentStructFunc(s ContentStructTest) {
	fmt.Printf("%p\n", &s)
	fmt.Println(unsafe.Sizeof(s)) // sum(len(elem))
}

// zero struct
func zeroStructFunc(s ZeroStructTest) {
	fmt.Printf("%p\n", &s)
	fmt.Println(unsafe.Sizeof(s))
}

// interface is a struct with size 2 word
func interfaceFunc(i interface{}) {
	fmt.Printf("%p\n", &i)
	fmt.Println(unsafe.Sizeof(i)) // 8
}

func main() {
	fmt.Println("==== string ====")
	s := "str str str"
	fmt.Printf("%p\n", &s)
	stringFunc(s)

	fmt.Println("==== array ====")
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("%p\n", &arr)
	arrayFunc(arr)

	fmt.Println("==== slice ====")
	sl := make([]int, 2)
	fmt.Printf("%p\n", &sl)
	sliceFunc(sl)

	fmt.Println("==== map ====")
	m := make(map[int]string)
	fmt.Printf("%p\n", &m)
	mapFunc(m)

	fmt.Println("==== chan ====")
	ch := make(chan int)
	fmt.Printf("%p\n", &ch)
	chanFunc(ch)

	fmt.Println("==== pointer ====")
	var n int = 2
	p := &n
	fmt.Printf("%p\n", &p)
	pointerFunc(p)

	fmt.Println("==== func ====")
	fn := func() {}
	fmt.Printf("%p\n", &fn)
	funcFunc(fn)

	fmt.Println("==== struct ====")
	// notice the address of zero struct
	st := ZeroStructTest{}
	fmt.Printf("%p\n", &st)
	zeroStructFunc(st)
	var c int = 2
	stc := ContentStructTest{1, "string string", &c}
	fmt.Printf("%p\n", &stc)
	contentStructFunc(stc)

	fmt.Println("==== interface ====")
	var i interface{} = 1
	fmt.Printf("%p\n", &i)
	interfaceFunc(i)
}
