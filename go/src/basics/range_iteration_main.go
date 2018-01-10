package main

import (
	"fmt"
)

func MapIterDel() {
	m := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
	}

	// delete during iteration
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
		if k <= 0 {
			delete(m, k+1)
		}
	}

	fmt.Println("====")

	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}

	fmt.Println("====")

	for k, _ := range m {
		if k >= 1 {
			delete(m, k-1)
		}
	}

	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}

}

func main() {
	MapIterDel()
}
