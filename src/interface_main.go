package main

import "fmt"

type Data struct {
	n int
}

func main() {
	d := Data{10}
	fmt.Printf("address of d: %p\n", &d)
	// assign not interface type variable to interface variable
	// d will be copied
	var i1 interface{} = d
	// assign interface type variable to interface variable
	// the data of i1 will directly assigned to i2.data and will not be copied
	var i2 interface{} = i1

	fmt.Println(d)
	fmt.Println(i1)
	fmt.Println(i2)
}
