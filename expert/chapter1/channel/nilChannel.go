package main

import "time"

// not work
func main() {
	var ch chan int
	go func() {
		time.Sleep(time.Second * 3) //deadlock
		ch = make(chan int, 10)
	}()
	go func() {
		ch <- 10
		println(<-ch)
	}()
	time.Sleep(time.Second * 10)
}
