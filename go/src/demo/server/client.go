package main

// simple client for simulating concurrent long or short tcp connections

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	raddr       string
	rport       string
	last        int
	concurrency int
	long        bool
	count       int
	interval    int
)

func init() {
	flag.StringVar(&raddr, "raddr", "127.0.0.1", "remote address to connect")
	flag.StringVar(&rport, "rport", "7777", "remote port to connect")
	flag.IntVar(&concurrency, "concurrency", 10, "concurrent workers to issue requests to remote server")
	flag.BoolVar(&long, "long", false, "issue long connections")
	flag.IntVar(&last, "last", 1000, "milliseconds to hold for long connections")
	flag.IntVar(&count, "count", 1000, "number of requests for each short connection")
	flag.IntVar(&interval, "interval", 200, "waiting milliseconds between requests")

}

func makeLongConn(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("connection %d\n", id)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", raddr, rport))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer conn.Close()
	buf := []byte("hello, world!\n")
	cnt := 0
	tm := time.After(time.Millisecond * time.Duration(last))
Loop:
	for {
		select {
		case <-tm:
			break Loop
		default:
			time.Sleep(time.Millisecond*100)
			wn, err := conn.Write(buf)
			if err != nil {
				fmt.Printf("write buf err %v\n", err)
				return
			}
			fmt.Printf("send data len %d\n", wn)

			rb := make([]byte, 4096)
			rn, err := conn.Read(rb)
			if err != nil {
				fmt.Printf("read conn err %v\n", err)
				return
			}
			fmt.Printf("recv data len %d\n", rn)
			cnt++
		}
	}
	fmt.Printf("connection %d: after %d second, total ping-pong count %d\n", id, last/1000, cnt)
}

func makeShortConn(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// issue short connections(requests) sequentially
	for i:=0; i<count; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", raddr, rport))
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		buf := []byte("hello, world!\n")
		wn, err := conn.Write(buf)
		if err != nil {
			fmt.Printf("write buf err %v\n", err)
		}
		fmt.Printf("send data len %d\n", wn)
		rb := make([]byte, 4096)
		rn, err := conn.Read(rb)
		if err != nil {
			fmt.Printf("read conn err %v\n", err)
		}
		fmt.Printf("recv data len %d\n", rn)
		conn.Close()
		// sleep a little bit
		time.Sleep(time.Millisecond*time.Duration(interval))
	}
}

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		if long {
			go makeLongConn(context.TODO(), i, &wg)
		} else {
			go makeShortConn(context.TODO(), i, &wg)
		}
	}
	wg.Wait()
}
