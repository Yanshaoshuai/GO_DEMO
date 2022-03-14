package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) //设置计数器，数值即goroutine的个数
	go func() {
		//Do some work
		time.Sleep(1 * time.Second)

		fmt.Println("Goroutine 1 finished!")
		wg.Done() //goroutine 执行结束后将计数器减一
	}()
	go func() {
		//Do some work
		time.Sleep(2 * time.Second)

		fmt.Println("Goroutine 2 finished!")
		wg.Done() //goroutine 执行结束后将计数器减一
	}()
	wg.Wait() //主 goroutine阻塞等待计数器变为0
	fmt.Println("All Goroutine finished!")
}
