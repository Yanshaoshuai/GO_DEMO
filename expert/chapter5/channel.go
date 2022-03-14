package main

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	//Do some work
	time.Sleep(time.Second)
	ch <- 1
}

// channel 用于并发控制
func main() {
	channels := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go Process(channels[i])
	}
	for i, ch := range channels { //遍历切片 等待子协程结束
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}
}
