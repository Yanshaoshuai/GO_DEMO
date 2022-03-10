package main

import "time"

func main() {
	ch := make(chan int)
	go func() {
		println(<-ch)
	}()
	go func() {
		println(<-ch)
	}()
	time.Sleep(time.Second * 2)
	ch <- 10
	close(ch)
	time.Sleep(time.Second * 2)
}
