package main

import (
	"fmt"
	"time"
)

func addNumberToChan(ch chan int) {
	for {
		ch <- 1
		time.Sleep(time.Second)
	}
}
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go addNumberToChan(ch1)
	go addNumberToChan(ch2)
	for {
		select {
		case e := <-ch1:
			fmt.Printf("Get element from ch1: %d\n", e)
		case e := <-ch2:
			fmt.Printf("Get element from ch2: %d\n", e)
		default:
			fmt.Printf("No element in ch1 and ch2 .\n")
			time.Sleep(time.Second)
		}
	}
}
