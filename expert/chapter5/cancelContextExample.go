package main

import (
	"context"
	"fmt"
	"time"
)

func HandleRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandlRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandleRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()
	time.Sleep(5 * time.Second)
}
