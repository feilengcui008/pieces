// string byte conversion testing(alloc and copy)
// string reslice to substring testing(not alloc)
/*
Breakpoint 1, main.main () at /tmp/string.go:30
30		s, b = StringToByteArrayTest()
(gdb) s
main.StringToByteArrayTest (~r0=..., ~r1=...) at /tmp/string.go:18
18	func StringToByteArrayTest() (string, []byte) {
(gdb) n
19		s := "1234567890"
(gdb) n
20		fmt.Printf("%#v\n", s)
(gdb) info locals
s = 0x4a7756 "1234567890"       <=====
(gdb) n
"1234567890"
23		return s[0:3], ([]byte)(s[0:3])
(gdb) n
main.main () at /tmp/string.go:31
31		fmt.Printf("%#v\n", s)
(gdb) info locals
s = 0x4a7756 "123"              <=====
b = {array = 0xc42000e1c8 "123", len = 3, cap = 8}


*/
package main

import (
	"fmt"
)

func ByteArrayToStringTest() (string, []byte) {
	buf := make([]byte, 10)
	for i := 0; i < 10; i++ {
		buf[i] = 1
	}
	fmt.Printf("%#v\n", buf)
	// string will alloc memory
	// buf just point to the same underlining array
	return string(buf), buf
}

func StringToByteArrayTest() (string, []byte) {
	s := "1234567890"
	fmt.Printf("%#v\n", s)
	// reslice to substring will use the same underlining memory
	// reslice to byte array will alloc new memory
	return s[0:3], ([]byte)(s[0:3])
}

func main() {
	s, b := ByteArrayToStringTest()
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", b)
	s, b = StringToByteArrayTest()
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", b)
}
