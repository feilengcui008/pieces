package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

/////////////////////////////////////////////////
//  基本语法
////////////////////////////////////////////////

//////////////////  变量与类型  ////////////////
// 包级变量
var packageLevel int
var a, b, c bool

// 自动推导类型
var m, n = 1, "i am a string"

func varTest() {
	// 函数级局部变量
	var funcLevel string
	funcLevel = "sss"
	// 声明并初始化，自动推导类型
	var funcStr1, funcInt1 = "func string", 12
	// 不能在函数外使用
	tempStr := "temp string"
	// 覆盖外面同名变量
	a, b, c := "a", "b", "c"
	fmt.Println(funcStr1, funcInt1, funcLevel, tempStr)
	fmt.Println(a, b, c)
}

// 基本类型
// bool
// uint8, uint16, uint32, uint, uint64, uintptr
//  int8,  int16,  int32,  int,  int64
// byte, rune, string
// float32, float64
// complex64, complex128
func typeTest() {
	var (
		a bool       = false
		b uint64     = 1<<64 - 1
		c complex128 = cmplx.Sqrt(-5 + 12i)
	)
	const f = "%T(%v)\n"
	fmt.Printf(f, a, a)
	fmt.Printf(f, b, b)
	fmt.Printf(f, c, c)

	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	// 类型转换，必须显示
	var z uint = uint(f1)
	//var k uint = f1  // error
	fmt.Println(x, y, z)

	// 常量可以是字符、字符串、布尔值、数值，不能使用:=
	const c1 = "const string"
	const c2 int = 12
	const c3 bool = false
	const (
		bigint   = 1 << 100
		smallint = 1 >> 20
	)
	fmt.Println(c1, c2, c3)
	fmt.Println(float64(bigint), smallint)
}

// 指针，不同于C/C++，没有指针运算
func pointerTest() {
	i := 12
	var p *int
	p = &i
	// 打印地址
	fmt.Println(p)
	// 取值
	fmt.Println(*p)
}

// 结构体struct与type
type MyInt int
type InnerStruct struct {
	name string
	p    *MyInt
}
type TestStruct struct {
	a     int
	b     string
	inner InnerStruct
}

func structTest() {
	var i MyInt = 12
	// 实例化结构体
	inner := InnerStruct{"name", &i}
	t := TestStruct{1, "aaa", inner}
	// 成员赋值
	t.a = 1111
	p := &t
	p.a = 2222
	fmt.Println(t)
	// 既可以使用结构体自身也可以使用其指针访问成员
	fmt.Println(t.a, t.b, p.inner.name)
	// 实例化结构体时只指定部分成员
	s1 := TestStruct{b: "sss", a: 1}
	fmt.Println(s1)
	s2 := TestStruct{}
	fmt.Println(s2)
}

// 数组，长度固定
func arrayTest() {
	// 定义与赋值
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr, arr[0])
	// 定义并直接初始化
	ar := [3]int{1, 2, 3}
	// 数组长度
	fmt.Println(ar, len(ar))
}

// 切片，长度可变
func sliceTest() {
	arr := []int{1, 2, 3, 4}
	var slice1 []int = arr[0:2]
	fmt.Println(slice1, len(slice1))
	slice1 = arr[2:]
	fmt.Println(slice1, len(slice1))
	// 切片只是浅拷贝，改变切片元素的值会改变对应数组元素的值
	slice1[0] = 111111
	fmt.Println(arr)
	// 切片初始化
	slice2 := []bool{true, false, true}
	// 切片长度和容量
	fmt.Println(slice2, len(slice2), cap(slice2))

	// make创建切片
	a := make([]int, 5) // len == 5
	a[0] = 1
	a[1] = 1
	//b := make([]int, 0, 5) // len == 0, cap == 5
	//b[0] = 1  // index out of range

	// 切片的切片
	c := [][]string{
		[]string{"s11", "s12"},
		[]string{"s21", "s22"},
	}
	fmt.Println(c)

	// append元素，当底层的支撑数组容量不够了时，会创建新的数组
	var slice3 []int
	slice3 = append(slice3, 0)
	slice3 = append(slice3, 1, 2)
	fmt.Println(slice3)
}

// 映射
func mapTest() {
	var m map[int]string
	//m[1] = "string"  // 为初始化不能赋值
	// 用make初始化
	m = make(map[int]string)
	m[12] = "12string"
	fmt.Println(m)
	m2 := map[string]int{
		"first":  1,
		"second": 2,
	}
	fmt.Println(m2)
	m2["first"] = 11
	delete(m2, "first")
	v, ok := m2["third"]
	v1, ok1 := m2["second"]
	fmt.Println(m2, v, ok)
	fmt.Println(m2, v1, ok1)
}

// 函数作为参数类型
func caller(fn func(int) int, n int) {
	fmt.Println(fn(n))
}
func funcTypeTest() {
	// 匿名函数
	fn := func(n int) int {
		return n + 1
	}
	// 高阶函数
	caller(fn, 11111)
}

//////////////////////  函数   ///////////////////////
// 简单函数与包的使用
func randomTest() {
	fmt.Println(rand.Intn(10))
}

// 同类型参数，多返回值
func swapString(a, b string) (string, string) {
	return b, a
}

// 返回值命名
func nameReturn(a int) (b, c int) {
	b = a + 1
	c = b * a
	return
}

//////////////////////  流程控制  ///////////////////////
// for
func forTest(n int) int {
	ret := 0
	for i := 0; i <= n; i++ {
		ret += i
	}
	// 初始化语句和后置语句可省略
	j := 0
	for j < 100 {
		j++
		ret += j
	}
	/*
		// infinite loop
		for {

		}
	*/
	// for range
	var slice1 []int = []int{1, 2, 3}
	var arr [3]int = [3]int{2, 3, 4}
	for k, v := range slice1 {
		fmt.Println(k, v)
	}
	for _, v := range arr {
		fmt.Println("v:", v)
	}
	fmt.Println(ret)
	return ret
}

// if
func ifTest() {
	i := 0
	// 简单用法
	if i == 0 {
		fmt.Println("==")
	}
	// 可以在条件表达式之前执行语句
	// j的作用域在if和else中
	if j := 12; j > 10 {
		fmt.Println("ok")
	} else {
		fmt.Println(j)
	}
}

// switch
func switchTest() {
	switch i := 1; i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default value")
	}
	j := 12
	switch true {
	case j < 10:
	case j >= 10:
		fmt.Println(j)
	default:
		fmt.Println("default value")
	}
	switch {
	}
}

// defer 通常可用在资源(文件/锁等)的析构上
func deferTest() {
	defer fmt.Println("defered print")
	fmt.Println("before defer print")

	// defer入栈，后进先出
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}

}

/////////////////////////////////////////////////
//  常用包的使用
////////////////////////////////////////////////

// IO

// Math

// String

// Date and Time

// Network and Socket

// Concurrency

// Decode and Encode

/////////////////////////////////////////////////
//  Go内部实现
////////////////////////////////////////////////

/////////////////////////////////////////////////
//  一些练习小程序
////////////////////////////////////////////////

// newton-method for sqrt
// x = x - (x^2 - y) / (2.0 * x)
func sqrt(x float64, delta float64) float64 {
	var ret float64 = 1
	for {
		temp := ret - (ret*ret-x)/(2.0*ret)
		if math.Abs(temp-ret) < delta {
			break
		}
		ret = temp
	}
	return ret
}

func main() {
	randomTest()
	fmt.Println(swapString("abc", "ttt"))
	fmt.Println(nameReturn(1))
	varTest()
	typeTest()
	forTest(100)
	ifTest()
	fmt.Println(sqrt(3, 0.000001), math.Sqrt(3))
	switchTest()
	deferTest()
	pointerTest()
	structTest()
	arrayTest()
	sliceTest()
	mapTest()
	funcTypeTest()
}
