// Test the go gc
// remember to set env GODEBUG=gctrace=1
//
// 字段分别表示:
// + 第几次
// + 开始sweep termination(gc)距离runtime启动以来的时间
// + 本次gc总共消耗的cpu百分比
// + sweep termination, mark, mark termination各阶段经过的时间
// + sweep termination, mark, mark termination各阶段消耗的cpu(注意是多个P即多个CPU的总时间，基本等于前面经过的时间*P数量)
// + sweep termination准备开始时(gc开始时)的可达内存heap_alive(即被标记对象的内存，<=分配内存), mark termination结束时的heap_alive，mark termination结束后总共被标记的内存
// + 下一次gc完成后的heap_alive目标值
// + P数量
// gc 34 @10.738s 5%: 0.013+424+0.091 ms clock, 0.053+4.4/83/333+0.36 ms cpu, 167->167->108 MB, 168 MB goal, 4 P
// gc 35 @11.235s 5%: 0.016+90+0.053 ms clock, 0.067+0.30/87/167+0.21 ms cpu, 182->188->147 MB, 216 MB goal, 4 P
//

package main

import (
	"time"
)

type Garbage struct {
	x int
	y int64
}

func main() {
	var lst []*Garbage
	for {
		for i := 0; i < 10000000; i++ {
			lst = append(lst, &Garbage{})
		}
		time.Sleep(time.Millisecond * 100)
		lst = nil
	}
	println(lst)
}
