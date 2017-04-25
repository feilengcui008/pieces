package main

import (
	"fmt"
)

func TestComposite() {
	const (
		Enone = iota
		Eio
		Einval
	)

	a := [...]string{Enone: "EnoneValue", Eio: "EioValue", Einval: "EinvalValue"}
	s := []string{Enone: "EnoneValue", Eio: "EioValue", Einval: "EinvalValue"}
	m := map[int]string{Enone: "EnoneValue", Eio: "EioValue", Einval: "EinvalValue"}

	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println(a[Enone])
	fmt.Println("====")

	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println(s[Enone])
	fmt.Println("====")

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m[Enone])
	fmt.Println(m[Eio])
}

func main() {
	TestComposite()
}
