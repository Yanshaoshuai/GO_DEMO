package main

import "fmt"

func RecoverDemo1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()
	panic("demo") //panic被recover之后无法回到当前位置继续执行
	fmt.Println("B")
}

func RecoverDemo2() {
	defer func() {
		fmt.Println("C")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()
	panic("demo")
	fmt.Println("B")
}

func RecoverDemo3() {
	defer func() {
		func() { //不被defer作用无法recover
			if err := recover(); err != nil {
				fmt.Println("A")
			}
		}()
	}()
	panic("demo")
	fmt.Println("B")
}

func RecoverDemo4() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("B")
		}
	}()
	panic("B")
	fmt.Println("C")
}

func RecoverDemo5() {
	foo := func() int {
		defer func() {
			recover()
		}()
		panic("demo")
		return 10
	}
	ret := foo() //匿名返回值返回零值 具名返回值返回已存在的值
	fmt.Println(ret)
}
func main() {
	println("RecoverDemo1")
	RecoverDemo1()
	println("RecoverDemo2")
	RecoverDemo2()
	println("RecoverDemo3")
	//RecoverDemo3()
	println("RecoverDemo4")
	RecoverDemo4()
	println("RecoverDemo5")
	RecoverDemo5()
}
