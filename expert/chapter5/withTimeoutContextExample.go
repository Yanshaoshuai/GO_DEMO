package main

import (
	"context"
	"fmt"
	"time"
)

func handelRequest(ctx context.Context) {
	go writeRedis(ctx)
	go writeDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}

	}
}
func writeRedis(ctx context.Context) {
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
func writeDatabase(ctx context.Context) {
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go handelRequest(ctx)
	time.Sleep(10 * time.Second)
}
