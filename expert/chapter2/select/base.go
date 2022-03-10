package main

import (
	"fmt"
)

func SelectForChan(c chan string) {
	var recv string
	send := "Hello"
	select {
	case recv = <-c:
		fmt.Printf("received %s\n", recv)
	case c <- send:
		fmt.Printf("send %s\n", send)
	}
}

//SelectAssign 读取管道时最多可以给两个变量赋值
func SelectAssign(c chan string) {
	select {
	case <-c: //0个变量
		fmt.Printf("0\n")
	case d := <-c: //1个变量
		fmt.Printf("1: received %s\n", d)
	case d, ok := <-c:
		if !ok { //ok代表是否成功读出数据 channel关闭也可能为true
			fmt.Printf("no data found")
			break
		}
		fmt.Printf("2: received %s,%t\n", d, ok)

	}
}

func SelectDefault() {
	c := make(chan string)
	select { //阻塞会走default分支
	default: //default可以在任何位置出现 每个select只能出现一次
		fmt.Printf("no data found in default\n")
	case <-c:
		fmt.Printf("received\n")

	}
}
func main() {
	c := make(chan string)
	go SelectForChan(c) //阻塞
	c = make(chan string, 1)
	//只有一个缓冲区 只能先写
	SelectForChan(c)
	c = make(chan string, 1)
	c <- "Hello"
	//缓冲区满 只能先读
	SelectForChan(c)
	c = make(chan string, 3)
	c <- "Hi"
	//缓冲区有数据且未满 可读可写 随机挑选一个执行
	SelectForChan(c)
	ch := make(chan string, 4)
	ch <- "hello "
	ch <- " world"
	ch <- "!"
	SelectAssign(ch)
	close(ch)
	SelectAssign(ch)

	//default
	SelectDefault()
}
