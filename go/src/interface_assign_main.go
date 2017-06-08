package main

import "fmt"

type data struct {
	x int
}

type dataLarge struct {
	a int
	b int
	c int
}

// data copied: assign not interface data to interface
// data not copied: assign interface to interface
func main() {
	v1 := data{10}
	v2 := dataLarge{10, 11, 12}
	fmt.Printf("%p\n", &v1)
	fmt.Printf("%p\n", &v2)
	var i1 interface{} = v1
	var i2 interface{} = v2
	var it int = 2
	fmt.Printf("%p\n", &it)
	var i3 interface{} = it

	var i4 interface{} = i3
	var i5 interface{} = uint8(1)

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)
	fmt.Println(i4)
	fmt.Println(i5)

	var d int = i4.(int)
	var di dataLarge = i2.(dataLarge)
	fmt.Printf("%p\n", &d)
	fmt.Printf("%p\n", &di)
	fmt.Printf("%p\n", &d)
}
