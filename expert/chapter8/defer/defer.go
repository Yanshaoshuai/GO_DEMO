package main

import "fmt"

func DeferDemo1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	//4,3,2,1,0
}

func DeferDemo2() {
	var aInt = 1
	fmt.Println(&aInt)
	//函数参数值传递
	defer fmt.Println(aInt)
	aInt = 2
	return
	//1
}

func DeferDemo3() {
	var i = 0
	//延迟函数中使用外部变量执行时才绑定
	defer func() {
		fmt.Println(i)
	}()
	i++
	//1
}

func DeferDemo4() {
	var aArray = [3]int{1, 2, 3}
	defer func(array *[3]int) {
		for i := range array {
			fmt.Println(array[i])
		}
	}(&aArray)
	aArray[0] = 10
	//10,2,3
}

//DeferDemo5 return拆解:设置返回值->执行defer->ret
func DeferDemo5() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
	//result =2
}

func DeferDemo6() {
	defer func() {
		defer func() {
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	//A B
}
func DeferDemo7() {
	i := 0
	defer func(a int) {
		fmt.Println(a)
		fmt.Println(i)
	}(i)
	i++
}
func main() {
	println("DeferDemo1")
	DeferDemo1()
	println("DeferDemo2")
	DeferDemo2()
	println("DeferDemo3")
	DeferDemo3()
	println("DeferDemo4")
	DeferDemo4()
	println("DeferDemo5")
	result := DeferDemo5()
	println(result)
	println("DeferDemo6")
	DeferDemo6()
	println("DeferDemo7")
	DeferDemo7()
}
