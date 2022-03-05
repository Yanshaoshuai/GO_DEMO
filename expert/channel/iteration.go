package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ch = make(chan int, 10)
	go func() {
		for value := range ch {
			fmt.Printf("value=%d\n", value)
		}
		println("iteration finish")
	}()
	for true {
		i := rand.Intn(10000)
		ch <- i
		if i == 4396 {
			time.Sleep(time.Second * 3)
			close(ch)
			break
		}
	}
}
