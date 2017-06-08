package main

import (
	"context"
	"fmt"
	"time"
)

func OuterLogicWithContext(ctx context.Context, fn func(ctx context.Context) error) error {
	go fn(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("OuterLogicWithContext ended")
			return ctx.Err()
		}
	}
}

func InnerLogicWithContext(ctx context.Context) error {
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		}
	}
	fmt.Println("InnerLogicWithContext ended")
	return ctx.Err()
}

func main() {
	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	ctx, cancel = context.WithTimeout(ctx, time.Second)
	go OuterLogicWithContext(ctx, InnerLogicWithContext)
	time.Sleep(time.Second * 3)
	// has been canceled by timer
	cancel()
	fmt.Println("main ended")
}
