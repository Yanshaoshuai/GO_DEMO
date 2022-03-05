package main

import (
	"fmt"
)

func main() {
	arr := [10]int{} //this is an array
	slice := []int{} //this is an array
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", slice)
	println()
	//变量声明
	var s []int //nil 不分配内存
	//字面值声明
	s1 := []int{} //not nil
	s2 := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s == nil)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
}
