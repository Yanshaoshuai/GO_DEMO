package main

import (
	"fmt"
	"time"
)

func foo() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
	panic("demo")
	defer fmt.Println("D")
}
func PanicDemo1() {
	defer func() {
		recover()
	}()
	foo()
}

func PanicDemo2() {
	defer func() {
		recover()
	}()
	defer func() {
		fmt.Println("1")
	}()
	foo()
}

func PanicDemo3() {
	defer func() {
		fmt.Println("demo")
	}()
	go foo()
}

func PanicDemo4() {
	defer func() {
		recover()
	}()
	defer fmt.Println("A")
	defer func() {
		fmt.Println("B")
		panic("panic in defer")
		fmt.Println("C")
	}()
	panic("panic")
	fmt.Println("D")
}
func main() {
	println("PanicDemo1")
	PanicDemo1()
	println("PanicDemo2")
	PanicDemo2()
	println("PanicDemo3")
	PanicDemo3()
	println("PanicDemo4")
	PanicDemo4()
	time.Sleep(2 * time.Second)
}
