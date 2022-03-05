package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var count int = 0
var mutexCh = make(chan int, 1)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			mutex()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			mutex()
		}
	}()
	var ch chan int
	println(ch == nil)
	//无缓冲
	var chNotNil = make(chan int)
	go func() {
		//block
		i := <-ch
		fmt.Printf("read nil channel = %d", i)
	}()
	go func() {
		//block
		ch <- 0
		fmt.Printf("write nil channel = %d\n", <-ch)
	}()
	go func() {
		//阻塞至写入成功
		fmt.Printf("read not nil channel = %d\n", <-chNotNil)
	}()
	go func() {
		chNotNil <- 10
		fmt.Printf("write not nil channel\n")
	}()
	time.Sleep(time.Second * 2)
	close(chNotNil)
	go func() {
		//0
		fmt.Printf("read not nil channel after close channel=%d\n", <-chNotNil)
	}()
	time.Sleep(time.Second * 10)
	go func() {
		//panic: send on closed channel
		chNotNil <- 10
		fmt.Printf("write not nil channel after close\n")
	}()
	time.Sleep(time.Second * 2)

}

func mutex() {
	mutexCh <- 1
	count++
	fmt.Printf("count=%d goruntine  id=%d\n", count, GetGID())
	<-mutexCh
}
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
