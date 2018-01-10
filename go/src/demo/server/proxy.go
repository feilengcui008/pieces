package main

// simple reverse proxy server with packet delay support

import (
	"context"
	"flag"
	"fmt"
	//"io"
	"math/rand"
	"net"
	"time"
)

var (
	lport  string
	raddr  string
	rport  string
	delay  int
	min    int
	max    int
	random bool
)

var (
	rd    = rand.New(rand.NewSource(time.Now().UnixNano()))
	count int
)

func init() {
	flag.StringVar(&lport, "lport", "7776", "local port to listen at")
	flag.StringVar(&raddr, "raddr", "127.0.0.1", "remote address")
	flag.StringVar(&rport, "rport", "7777", "remote port")
	flag.IntVar(&delay, "delay", 30, "delay ms")
	flag.IntVar(&min, "min", 10, "min ms for random delay")
	flag.IntVar(&max, "max", 200, "max ms for random delay")
	flag.BoolVar(&random, "random", false, "switch on random delay")
}

func delayRW(ctx context.Context, src net.Conn, dst net.Conn, closeC chan struct{}) {
	defer func() {
		closeC <- struct{}{}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf := make([]byte, 4096)
			n, err := src.Read(buf)
			if err != nil {
				fmt.Printf("local %s -> remote %s closed\n", src.LocalAddr(), src.RemoteAddr())
				return
			}
			dt := delay
			if random {
				dt = min + rd.Intn(max-min)
			}
			if dt < 0 {
				// drop packet
				fmt.Printf("delay infinitely, drop packet\n")
				return
			}
			if dt > 0 {
				<-time.After(time.Millisecond * time.Duration(dt))
				fmt.Printf("delayed %d ms\n", dt)
			}
			n, err = dst.Write(buf[:n])
			if err != nil {
				fmt.Printf("local %s -> remote %s closed\n", dst.LocalAddr(), dst.RemoteAddr())
				return
			}
		}
	}
}
func handleConn(ctx context.Context, conn net.Conn, countC chan int) {
	var (
		rconn net.Conn
		err   error
	)
	if rconn, err = net.Dial("tcp", fmt.Sprintf("%s:%s", raddr, rport)); err != nil {
		fmt.Printf("connect to remote %s:%s faild %v\n", raddr, rport, err)
		return
	}
	closeC := make(chan struct{})
	go delayRW(ctx, conn, rconn, closeC)
	go delayRW(ctx, rconn, conn, closeC)
	<-closeC
	conn.Close()
	rconn.Close()
	countC <- int(-1)
}

func handleCount(ctx context.Context, c chan int) {
	tm := time.After(time.Second * 1)
	for {
		select {
		case <-ctx.Done():
			return
		case op := <-c:
			if op == 1 {
				count++
			} else {
				count--
			}
		case <-tm:
			fmt.Printf("total connection: %d\n", count)
			tm = time.After(time.Second * 1)
		}
	}
}

func main() {
	flag.Parse()
	var (
		l           net.Listener
		err         error
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()

	if l, err = net.Listen("tcp", fmt.Sprintf(":%s", lport)); err != nil {
		fmt.Printf("listen err %v\n", err)
		return
	}
	defer l.Close()

	countC := make(chan int)
	go handleCount(ctx, countC)

	for {
		var conn net.Conn
		if conn, err = l.Accept(); err != nil {
			fmt.Printf("accept err %v\n", err)
			return
		}
		fmt.Printf("got a new conn from %s, current connection count %d\n", conn.RemoteAddr(), count)
		countC <- int(1)
		go handleConn(ctx, conn, countC)
	}
}
