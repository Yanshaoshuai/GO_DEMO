package main

import "fmt"

func main() {
	var ch chan int
	println(ch == nil)
	chWithoutBuffer := make(chan int)
	println(chWithoutBuffer == nil)
	chWithBuffer := make(chan int, 10)
	chWithBuffer <- 10
	println(<-chWithBuffer)
	chWithBuffer <- 10
	readCh(chWithBuffer)
	writeCh(chWithBuffer)
	println(<-chWithBuffer)
	readWriteCh(chWithBuffer)

	//ok==true means  has value
	chWithBuffer <- 10
	value, ok := <-chWithBuffer
	fmt.Printf("value=%d,ok=%v\n", value, ok)
	close(chWithBuffer)
	value, ok = <-chWithBuffer
	fmt.Printf("value=%d,ok=%v\n", value, ok)
	chCloseWithValue := make(chan int, 3)
	chCloseWithValue <- 10
	chCloseWithValue <- 9
	close(chCloseWithValue)
	value, ok = <-chCloseWithValue
	fmt.Printf("value=%d,ok=%v\n", value, ok)
	println(len(chCloseWithValue))
	println(cap(chCloseWithValue))
}

func readCh(ch <-chan int) {
	//ch<-10//read only chan
	println(<-ch)
}
func writeCh(ch chan<- int) {
	//println(<-ch)//write only
	ch <- 10
}
func readWriteCh(ch chan int) {
	ch <- 10
	println(<-ch)
}
