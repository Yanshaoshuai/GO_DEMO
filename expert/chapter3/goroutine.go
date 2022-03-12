package main

import (
	"runtime"
	"time"
)

//1.14之前会陷入无限循环，协程永远无法被抢占
func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for {

		}
	}()
	time.Sleep(1 * time.Second)
	println("Done")
}
