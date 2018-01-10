package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

// demo for testing struct none pointer receiver
// data copy and mutex losing effect due to this.
// try to use pointer receiver for struct with
// mutex and shared data all the time.

type DataAnonymousFieldMutex struct {
	sync.Mutex
	x int
}

// data will be copied
// mutex will be copied and lose effect
func (d DataAnonymousFieldMutex) AddAndPrint(i int) {
	//func (d *DataAnonymousFieldMutex) AddAndPrint(i int) {
	d.Lock()
	defer d.Unlock()
	d.x += 1
	fmt.Printf("goroutine index %v\n", i)
	fmt.Printf("addr of d: %p, content of d: %#v\n", &d, d)
	fmt.Printf("addr of d.x: %p, content of d.x: %v\n", &d.x, d.x)
}

func DataAnonymousFieldMutexTest() {
	d := DataAnonymousFieldMutex{x: 0}
	for i := 1; i <= 5; i++ {
		go func(i int) {
			d.AddAndPrint(i)
		}(i)
	}
	select {
	case <-time.After(time.Second):
		return
	}
}

type DataNamedFieldMutex struct {
	mu sync.Mutex
	x  int
}

// data will be copied and mutex will lose effect
func (d DataNamedFieldMutex) AddAndPrint(i int) {
	//func (d *Data) AddAndPrint(i int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.x += 1
	fmt.Printf("goroutine index %v\n", i)
	fmt.Printf("addr of d: %p, content of d: %#v\n", &d, d)
	fmt.Printf("addr of d.x: %p, content of d.x: %v\n", &d.x, d.x)
	fmt.Printf("addr of d.mu: %p, content of d.mu: %p, %v\n", &d.mu, d.mu, d.mu)
}

func DataNamedFieldMutexTest() {
	d := DataNamedFieldMutex{x: 0, mu: sync.Mutex{}}
	for i := 1; i <= 5; i++ {
		go func(i int) {
			d.AddAndPrint(i)
		}(i)
	}
	select {
	case <-time.After(time.Second):
		return
	}
}

type DataPointerMutex struct {
	mu *sync.Mutex
	x  int
}

// will work, but data still copied
func (d DataPointerMutex) AddAndPrint(i int) {
	//func (d *DataMutex) AddAndPrint(i int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.x += 1
	fmt.Printf("goroutine index %v\n", i)
	fmt.Printf("addr of d: %p, content of d: %#v\n", &d, d)
	fmt.Printf("addr of d.x: %p, content of d.x: %v\n", &d.x, d.x)
	fmt.Printf("addr of d.mu: %p, content of d.mu: %p, %v\n", &d.mu, d.mu, d.mu)

}

func DataPointerMutexTest() {
	d := DataPointerMutex{x: 0, mu: &sync.Mutex{}}
	for i := 1; i <= 5; i++ {
		go func(i int) {
			d.AddAndPrint(i)
		}(i)
	}
	select {
	case <-time.After(time.Second):
		fmt.Printf("=== %v\n", d.x)
		return
	}
}
func main() {
	var mu sync.Mutex
	fmt.Printf("sizeof mutex: %v\n", unsafe.Sizeof(mu)) // mutex is just a pointer
	fmt.Println("=== DataAnonymousFieldMutexTest ===")
	DataAnonymousFieldMutexTest()
	fmt.Println("=== DataNamedFieldMutexTest ===")
	DataNamedFieldMutexTest()
	fmt.Println("=== DataPointerMutexTest ===")
	DataPointerMutexTest()
}
