// Some concurrent patterns using goroutines and channels for fun
//SimpleSynchronizationPattern1()
//SimpleSynchronizationPattern2()
//BatchSynchronizationPattern()
//BatchNotificationPattern()
//ProducerWorkerConsumerPattern()
//PipelineProcessingPattern()
//PrimeSievePipeline()
//PingPongPattern()

package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Simple-Synchronization-Pattern
func SimpleSynchronizationPattern1() {
	c := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 2)
		// or close(c)
		c <- 1
	}()
	<-c
	log.Printf("waiting another goroutine finished\n")
}

func SimpleSynchronizationPattern2() {
	fn := func() chan interface{} {
		c := make(chan interface{})
		go func() {
			time.Sleep(time.Second * 2)
			c <- 1
		}()
		return c
	}
	c := fn()
	<-c
	log.Printf("waiting another goroutine finished\n")
}

// Batch-Synchronization-Pattern
func BatchSynchronizationPattern() {
	var wg sync.WaitGroup
	fn := func(wg *sync.WaitGroup, i int) {
		go func() {
			time.AfterFunc(time.Second*2, func() {
				go func() {
					log.Printf("gopher %d done\n", i)
					wg.Done()
				}()
			})
		}()
	}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		fn(&wg, i)
	}
	wg.Wait()
}

// Batch-Notification-Pattern
func BatchNotificationPattern() {
	c := make(chan interface{})
	fn := func(i int, ch chan interface{}) {
		go func() {
			<-ch
			log.Printf("gopher %d received close chan signal\n", i)
		}()
	}
	for i := 1; i <= 10; i++ {
		fn(i, c)
	}
	// or use context, see the following examples
	close(c)
	<-time.After(time.Second * 2)
}

// Producer-Worker-Consumer Pattern
type Task struct {
	Id int
}

type Result struct {
	Value int
}

func NewWorker(ctx context.Context, c <-chan Task, ch chan<- Result, id int) {
	go func() {
		for {
			select {
			case v, ok := <-c:
				if !ok {
					log.Printf("task chan closed\n")
					return
				}
				ch <- Result{Value: v.Id}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func NewProducer(ctx context.Context, c chan<- Task, n int) {
	go func() {
		i := 1
		for {
			select {
			case c <- Task{Id: i}:
				i += 1
				if i > n {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func NewConsumer(ctx context.Context, ch <-chan Result) {
	go func() {
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					log.Printf("receive chan closed in master\n")
					return
				}
				log.Printf("result %v\n", v)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func ProducerWorkerConsumerPattern() {
	c := make(chan Task)
	ch := make(chan Result)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		NewWorker(ctx, c, ch, i)
	}
	NewConsumer(ctx, ch)
	NewProducer(ctx, c, 100)
	select {
	case <-time.After(time.Second * 3):
		cancel()
	}
}

// PipelineProcessingPattern
func PipelineProcessingPattern() {
	type Pipe struct {
		Data chan int
	}
	var chans []Pipe
	for i := 1; i <= 10; i++ {
		p := Pipe{Data: make(chan int)}
		chans = append(chans, p)
	}
	ctx, cancel := context.WithCancel(context.Background())
	fn := func(ctx context.Context, c <-chan int, ch chan<- int, i int) {
		go func() {
			for {
				select {
				case v := <-c:
					//log.Printf("gopher %d received value %d\n", i, int(v))
					ch <- v + 1
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	for i := 1; i < 10; i++ {
		fn(ctx, chans[i-1].Data, chans[i].Data, i)
	}
	// root sender
	go func(ctx context.Context, n int) {
		i := 0
		for {
			select {
			case chans[0].Data <- i:
				i += 1
				if i > n {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx, 1)
	// root receiver
	go func(ctx context.Context) {
		for {
			select {
			case v := <-chans[9].Data:
				log.Printf("root receiver gopher receved valued: %d", int(v))
				if v > 100 {
					return
				}
				// send back to form a circle
				chans[0].Data <- v + 1
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	<-time.After(time.Second * 2)
	cancel()
}

// prime sieve, use PipelineProcessingPattern
func PrimeSievePipeline() {
	const (
		N      = 100
		FirstK = 20
	)
	var chans []chan int
	for i := 0; i < FirstK+1; i++ {
		chans = append(chans, make(chan int))
	}
	gen := func(n int) {
		go func() {
			for i := 2; i <= n; i++ {
				chans[0] <- i
			}
		}()
	}
	filter := func(c <-chan int, ch chan<- int, id int) {
		go func() {
			// got the first prime first
			prime := <-c
			//if prime > N {
			//	return
			//}
			log.Printf("filter gopher %d got one prime: %d", id, prime)
			for v := range c {
				if v%prime != 0 {
					ch <- v
				}
			}
		}()
	}
	// find first 9 primes
	for i := 1; i <= FirstK; i++ {
		filter(chans[i-1], chans[i], i)
	}
	gen(N)
	time.Sleep(time.Second)
}

// PingPongPattern
func PingPongPattern() {
	c := make(chan int)
	ch := make(chan int)
	gopher := func(ctx context.Context, c <-chan int, ch chan<- int, id int) {
		go func() {
			for {
				select {
				case v := <-c:
					log.Printf("gopher %d got: %d\n", id, v)
					ch <- v + 1
					time.Sleep(time.Second)
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	ctx, cancel := context.WithCancel(context.Background())
	gopher(ctx, c, ch, 0)
	gopher(ctx, ch, c, 1)
	c <- 0
	time.Sleep(time.Second * 4)
	cancel()
}

func main() {
	//SimpleSynchronizationPattern1()
	//SimpleSynchronizationPattern2()
	//BatchSynchronizationPattern()
	//BatchNotificationPattern()
	//ProducerWorkerConsumerPattern()
	//PipelineProcessingPattern()
	PrimeSievePipeline()
	//PingPongPattern()
}
