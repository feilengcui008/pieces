package main

import (
	"fmt"
	"time"
)

// 用channel和goroutine模拟生成器
func yield(i int) chan<- int {
	fmt.Printf("make a generator %v\n", i)
	c := make(chan int)
	go func(i int) {
		value := <-c
		fmt.Printf("routine %v: get value %d\n", i, value)
	}(i)
	return c
}

func main() {
	const n = 10
	var gens [n]chan<- int
	fmt.Printf("make %v generators\n", n)
	for i := 0; i < n; i++ {
		gens[i] = yield(i)
	}
	fmt.Printf("feed them content\n")
	for i := 0; i < n; i++ {
		gens[i] <- i
	}
	// 等输出
	select {
	case <-time.After(time.Second * 2):
		break
	}
}
