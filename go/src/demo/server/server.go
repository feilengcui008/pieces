package main

// simple echo server

import (
	"flag"
	"fmt"
	"net"
)

var lport string

func init() {
	flag.StringVar(&lport, "lport", "7777", "local port to listen at")
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 10240)
		rn, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read error %v\n", err)
			return
		}
		fmt.Printf("read data len %v, %v\n", rn, buf[:rn])
		wn, err := conn.Write(buf[:rn])
		if err != nil {
			fmt.Printf("write err %v\n", err)
			return
		}
		fmt.Printf("write data len %v\n", wn)
	}
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", lport))
	if err != nil {
		fmt.Printf("listen err %v\n", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept err %v\n", err)
			break
		}
		fmt.Printf("got a conn\n")
		go handleConn(conn)
	}

}
