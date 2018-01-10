package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// LimitReader读取固定长度的数据，并返回io.EOF
func LimitReaderFunc() {
	lr := io.LimitReader(os.Stdin, 16)
	buf := make([]byte, 5)
	n, err := lr.Read(buf)
	fmt.Printf("%v, %v, %v\n", n, strconv.Quote(string(buf)), err)
	for {
		n, err = lr.Read(buf)
		if n > 0 {
			fmt.Printf("%v, %v, %v\n", n, string(buf), err)
		}
		if err != nil {
			fmt.Printf("got err %v\n", err)
			break
		}
	}
}

// SectionReader从某个段[begin, begin+n]的某个offset起始读取，超过范围返回io.EOF，可seek
func SectionReaderFunc() {
	r, err := os.Open("/etc/resolv.conf")
	if err != nil {
		return
	}
	b, e := int64(3), int64(5)
	sr := io.NewSectionReader(r, b, e)
	buf := make([]byte, 5)
	n, err := sr.Read(buf)
	fmt.Printf("read begin at offset %v len %v: %v, %v, %v\n", b, e, n, string(buf), err)
	// 再次读取触发io.EOF
	n, err = sr.Read(buf)
	fmt.Printf("%v, %v\n", n, err)
}

// TeeReader同时读取并写入
func TeeReaderFunc() {
	r, w := os.Stdin, os.Stdout
	tr := io.TeeReader(r, w)
	buf := make([]byte, 16)
	for {
		n, err := tr.Read(buf)
		if n > 0 {
			fmt.Printf("from direct read: %s\n", string(buf))
		}
		if err != nil {
			break
		}
	}
}

func main() {
	//LimitReaderFunc()
	SectionReaderFunc()
}
