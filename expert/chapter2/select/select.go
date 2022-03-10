package main

import "fmt"

//SelectExam1 既可能输出c1
//也可能输出c2
func SelectExam1() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	c1 <- 1
	c2 <- 2
	select {
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	}
}

//SelectExam2 阻塞
func SelectExam2() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("readable")
	case c <- 1:
		fmt.Println("writable")
	}
}
func SelectExam3() {
	c := make(chan int, 10)
	c <- 1
	close(c)
	select {
	case d := <-c:
		fmt.Println(d)
	}
}

func SelectExam4() {
	c := make(chan int, 10)
	c <- 1
	close(c)
	select {
	case d, ok := <-c:
		if !ok {
			fmt.Println("no data received")
			break
		}
		fmt.Println(d)
	}
}

//SelectExam5 阻塞
func SelectExam5() {
	select {}
}
func SelectExam6() {
	var c chan string
	select {
	case c <- "Hello": //case语句操作nil会被忽略
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}
func main() {
	SelectExam1()
	go func() {
		SelectExam2()
	}()
	SelectExam3()
	SelectExam4()
	go func() {
		SelectExam5()
	}()
	SelectExam6()
}
