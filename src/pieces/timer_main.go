package main

import (
	"fmt"
	"time"
)

func testTimer() {
	t := time.NewTimer(time.Second * 2)
	select {
	case tm := <-t.C:
		fmt.Printf("time is up: %v\n", tm)
		break
	}
	t1 := time.AfterFunc(time.Second, func() {
		fmt.Println("call from timer goroutine")
	})
	// can not be select anymore
	// but can stop
	//select {
	//case tm := <-t1.C:
	//	fmt.Printf("t1 is up: %v", tm)
	//	break
	//}
	t1.Stop() // Stop is used to cancel the call of f inner timer goroutine

	select {
	case <-time.After(time.Second * 2):
		break
	}
}

func main() {
	testTimer()
}
