package main

import (
	"context"
	"fmt"
	"time"
)

func handelRequestValue(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handelRequestValue Done.")
			return
		default:
			fmt.Println("handelRequestValue running,parameter:", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "parameter", "1")

	go handelRequestValue(ctx)

	time.Sleep(10 * time.Second)
}
