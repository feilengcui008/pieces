// dlv debug
// b flag
// c

/*
we can see what goroutines(G) and threads(M) the runtime started:


====

all: goroutines
change to some goroutine: goroutine number
see stack trace: goroutine

(dlv) goroutines
[10 goroutines]
* Goroutine 1 - User: ./gc_main.go:13 main.flag (0x44fff0) (thread 29837)
  Goroutine 2 - User: /home/tan/me/software/go/src/runtime/proc.go:272 runtime.gopark (0x424f2a)
  Goroutine 3 - User: /home/tan/me/software/go/src/runtime/proc.go:272 runtime.gopark (0x424f2a)
  Goroutine 4 - User: /home/tan/me/software/go/src/runtime/time.go:59 time.Sleep (0x43ac69)
  Goroutine 5 - User: /home/tan/me/software/go/src/runtime/time.go:59 time.Sleep (0x43ac69)
  Goroutine 6 - User: /home/tan/me/software/go/src/runtime/time.go:59 time.Sleep (0x43ac69)
  Goroutine 7 - User: /home/tan/me/software/go/src/runtime/time.go:59 time.Sleep (0x43ac69)
  Goroutine 8 - User: /home/tan/me/software/go/src/runtime/time.go:59 time.Sleep (0x43ac69)
  Goroutine 9 - User: /home/tan/me/software/go/src/runtime/time.go:59 time.Sleep (0x43ac69)
  Goroutine 17 - User: /home/tan/me/software/go/src/runtime/lock_futex.go:206 runtime.notetsleepg (0x40a7d2)

1. runtime.main goroutines
(dlv) goroutine
Thread 29837 at ./gc_main.go:13
Goroutine 1:
	Runtime: ./gc_main.go:13 main.flag (0x44fff0)
	User: ./gc_main.go:13 main.flag (0x44fff0)
	Go: /home/tan/me/software/go/src/runtime/asm_amd64.s:165 runtime.rt0_go (0x4468b2)

2. runtime gchelper goroutine
(dlv) goroutine
Thread 29837 at ./gc_main.go:13
Goroutine 2:
	Runtime: /home/tan/me/software/go/src/runtime/proc.go:272 runtime.gopark (0x424f2a)
	User: /home/tan/me/software/go/src/runtime/proc.go:272 runtime.gopark (0x424f2a)
	Go: /home/tan/me/software/go/src/runtime/proc.go:216 runtime.init.4 (0x424c75)

3. runtime bgsweeper goroutine
(dlv) goroutine
Thread 29837 at ./gc_main.go:13
Goroutine 3:
	Runtime: /home/tan/me/software/go/src/runtime/proc.go:272 runtime.gopark (0x424f2a)
	User: /home/tan/me/software/go/src/runtime/proc.go:272 runtime.gopark (0x424f2a)
	Go: /home/tan/me/software/go/src/runtime/mgc.go:211 runtime.gcenable (0x411691)

17. runtime timer proc goroutine
(dlv) goroutine
Thread 29837 at ./gc_main.go:13
Goroutine 17:
	Runtime: /home/tan/me/software/go/src/runtime/lock_futex.go:206 runtime.notetsleepg (0x40a7d2)
	User: /home/tan/me/software/go/src/runtime/lock_futex.go:206 runtime.notetsleepg (0x40a7d2)
	Go: /home/tan/me/software/go/src/runtime/time.go:118 runtime.addtimerLocked (0x43ae2d)

*/

package main

import (
	"time"
)

type Garbage struct {
	x int
	y int64
}

func flag() {
}

func test() {
	for i := 0; i < 6; i++ {
		go func() {
			println(1)
			for {
				time.Sleep(time.Second * 2)
			}
		}()
	}
	flag()
}

func main() {
	test()
	select {}
}
