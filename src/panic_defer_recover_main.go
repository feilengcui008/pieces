package main

import (
	"fmt"
)

// the defer flow is as following
// 1. once a panic called, if no defer recover, whole system exit
// 2. once a panic called, if some defer recovered, the control logic will be return at
// the end of the func where defer is declared, the logic behind panic will not be executed anymore
// 3. you can repanic in a defer block, the defer info is pushed into the linked list of each p

//
// the compiled code locgic like this:
//
// deferproc -> deferproc -> deferreturn    deferreturn    deferreturn
//           |            |                ^              ^
// 		       |            |                |              |
// 		       |            -----------------|              |
// 		       |                recovered                   |
//		       |                                            |
//		       |--------------------------------------------|
//					                  recovered
//

func OneDefer() {
	defer fmt.Println(1)
}

func TwoDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
}

func TwoDeferWithCodeLogic() {
	defer fmt.Println(1)
	var i int = 0x2222
	fmt.Println(i)
	defer fmt.Println(2)
}

func OnePanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var i int = 1
	fmt.Println(i)
	panic("paniced!")

	// not reach anymore
	i = 0x2222
	fmt.Println(i)
}

func TwoPanicRecover() {
	defer func() {
		fmt.Println("first defer")
		if r := recover(); r != nil {
			fmt.Printf("first defer got panic msg: %v\n", r)
		}
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("second defer got panic msg: %v\n", r)
		}
	}()
	var i int = 1
	fmt.Println(i)
	panic("paniced!")

	// not reach here anymore
	i = 2
	fmt.Println(i)
	panic("paniced again!")
	i = 3
	fmt.Println(i)
}

func PanicInDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("first defer got panic msg: %v\n", r)
		}
	}()

	defer func() {
		fmt.Println("second defer do nothing")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("third defer got panic msg: %v\n", r)
			panic("repanic in third defer")
		}
	}()
	panic("panic msg")
}

func PrintLine() {
	fmt.Println("======")
}

func main() {
	OneDefer()
	TwoDefer()
	TwoDeferWithCodeLogic()

	PrintLine()
	OnePanicRecover()
	PrintLine()
	TwoPanicRecover()
	PrintLine()
	PanicInDefer()
}
