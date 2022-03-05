package main

import "fmt"

//new 是一个分配内存的内建函数，但不同于其他语言中同名的new所作的工作，它只是将内存清零，而不是初始化内存。
//new(T)为一个类型为T的新项目分配了值为零的存储空间并返回其地址，也就是一个类型为*T的值。用Go的术语来说，就是它返回了一个指向新分配的类型为T的零值的指针。
//make(T, args)函数的目的与new(T)不同。它仅用于创建切片、map和chan（消息管道），并返回类型T（不是*T）的一个被初始化了的（不是零）实例。

func main() {
	var p *[]int = new([]int)     // 为切片结构分配内存；*p == nil；很少使用
	var v []int = make([]int, 10) // 切片v现在是对一个新的有10个整数的数组的引用

	// 不必要地使问题复杂化：
	p = new([]int)
	fmt.Println(p) //输出：&[]
	*p = make([]int, 10, 10)
	fmt.Println(p)       //输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) //输出： 0

	// 习惯用法:
	v = make([]int, 10)
	fmt.Println(v) //输出：[0 0 0 0 0 0 0 0 0 0]
}
