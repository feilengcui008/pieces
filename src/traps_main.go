package main

import (
	"fmt"
	"time"
)

// 一些Go语言使用中容易出现的问题，目前添加
// 1. 使用defer和go语句时闭包相关问题
// 2. point receiver和非pointer receiver相关的问题
// 3. nil相关问题

/* 闭包相关，在使用defer和go语句时需要注意 */
func deferFunc1() {
	fn := func(i int) {
		fmt.Printf("deferFunc1: %d\n", i)
	}
	// 输出0-9
	for i := 0; i < 10; i++ {
		// 相当于 defer func(i int){fn(i)}(i)，所以立即传值
		defer fn(i)
	}
}

func deferFunc2() {
	// 输出0-9
	for i := 0; i < 10; i++ {
		defer func(i int) {
			fmt.Printf("deferFunc2: %d\n", i)
		}(i)
	}
}

func deferFunc3() {
	// 全部输出10，由于直接引用了外面作用域的变量，而不是传值进去
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("deferFunc3: %d\n", i)
		}()
	}
}

func goFunc1() {
	// 输出0-9
	for i := 0; i < 10; i++ {
		go fmt.Printf("goFunc1: %d\n", i)
	}
}

func goFunc2() {
	// 输出0-9
	for i := 0; i < 10; i++ {
		go func(arg int) {
			fmt.Printf("goFunc2: %d\n", arg)
		}(i)
	}
}

func goFunc3() {
	// 全部输出10
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("goFunc3: %d\n", i)
		}()
	}
}

/* 结构体的receiver，非指针类型的receiver会触发结构体的拷贝 */
type MyInt int

func (m *MyInt) Add(n int) *MyInt {
	*m += MyInt(n)
	return m
}

func (m MyInt) Minus(n int) *MyInt {
	m -= MyInt(n)
	return &m
}

func ReceiverTest() {
	m := MyInt(5)
	mp := &m
	fmt.Printf("%p\n", mp)
	// a == mp
	a := mp.Add(1)
	fmt.Printf("%p\n", a)
	// b point to another address diff from a and mp
	// so copy happened
	b := mp.Minus(1)
	fmt.Printf("%p\n", b)
}

/* nil相关，尤其是在返回定制化的error时。
接口包含三个字段: _type、pkgname、[]imethod接口方法数组
结构体也包含三个字段: _type、pkgname、[]structfield

1. 结构体指针和接口可以赋予nil值，当把结构体指针转换为接口时，
接口的_type字段不会为nil，因此接口不会为nil, 只有当接口的类型
信息也为nil时，此接口才==nil
2. 给接口赋值nil时其类型信息_type也为nil，所以整体==nil
3. 将其他类型强制转换为接口时，由于在其定义的时候类型信息即存在
了，所以无论其值是不是nil，其类型信息始终导致转换后的接口永远
不为nil，除非其本身就没类型即nil
4. 不同类型nil比较的语义不同。只有slice、map、chan、interface、pointer
可以直接和nil比较，array、string、struct不行，但是所有类型都可以转换
为接口类型比如interface{}
*/
type Cat interface {
	Meow()
}
type Tabby struct{}

func (*Tabby) Meow() {
	fmt.Println("meow")
}

func (Tabby) Error() string {
	return ""
}

type MyError struct{}

func (e *MyError) Error() string {
	return ""
}

// 接口组合，这种情况下会返回nil
type CatChild interface {
	Cat
	Say()
}

// 结构体转换为返回的接口类型Cat时会保留类型信息，所以不为nil
func returnStructPointer() Cat {
	var myTabby *Tabby = nil
	return myTabby
}

func returnStruct() error {
	// 结构体不能直接和nil比较，但是转换为接口后同样不为nil
	var myTabby Tabby
	return myTabby
}

func returnString() interface{} {
	// 字符串不能直接和nil比较
	var s string
	return s
}

func returnArray() interface{} {
	var array [3]int
	//数组不能直接和nil比较
	//fmt.Printf("array is nil: %v\n", interface{}(array) == nil)
	return array
}

// error接口同理
func returnsError() error {
	var p *MyError = nil
	return p // Will always return a non-nil error.
}

func returnChildInterface() Cat {
	var cc CatChild = nil
	return cc
}

func returnIntPointer() interface{} {
	var a *int = nil
	return a
}

func returnChan() interface{} {
	var ch <-chan int = nil
	return ch
}

func returnSlice() interface{} {
	var slice []int = nil
	return slice
}

func returnMap() interface{} {
	var m map[int]string = nil
	return m
}

func NilTest() {
	err := returnsError()
	if err != nil {
		fmt.Println("err not nil")
	}

	st := returnStruct()
	if st != nil {
		fmt.Println("struct not nil")
	}

	s := returnString()
	if s != nil {
		fmt.Println("string not nil")
	}

	array := returnArray()
	if array != nil {
		fmt.Println("array not nil")
	}

	stp := returnStructPointer()
	if stp != nil {
		fmt.Println("struct* not nil")
	}

	a := returnIntPointer()
	if a != nil {
		fmt.Println("int* not nil")
	}

	ch := returnChan()
	if ch != nil {
		fmt.Println("chan not nil")
	}

	slice := returnSlice()
	if slice != nil {
		fmt.Println("slice not nil")
	}

	m := returnMap()
	if m != nil {
		fmt.Println("map not nil")
	}

	// 返回接口类型比较特殊
	c := returnChildInterface()
	if c == nil {
		fmt.Println("interface is special: nil")
	} else {
		fmt.Println("interface cat not nil")
	}
}

func main() {
	//deferFunc1()
	//deferFunc2()
	//deferFunc3()
	//goFunc1()
	//goFunc2()
	//goFunc3()

	ReceiverTest()
	NilTest()

	select {
	case <-time.After(time.Second * 2):
		break
	}
}
