package main

import "fmt"

func fun1() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i=%d,j=%d\n", i, j)
}

//func fun2(i int) {
//	i:=0//变量i已存在 左边没有新变量
//	fmt.Println(i)
//}
func fun3() {
	i, j := 0, 0
	if true {
		j, k := 1, 1
		fmt.Printf("j=%d,k=%d\n", j, k)
	}
	fmt.Printf("i=%d,j=%d\n", i, j)
}

//简短变量声明只能用于函数中
//rule:="Short variable declarations"
func main() {
	fun1()
	fun3()
	Greeting("nobody")
	Greeting("hello:", "Joe", "Anna", "Eileen")
	guest := []string{"Joe", "Anna", "Eileen"}
	//传入切片 不会生成新切片,函数内部使用的切片与传入的切片共享内存
	Greeting("hello:", guest...)
}

//Println 可变参数
func Println(a ...interface{}) {

}

//可变参数必须在函数列表的最后
//可变参数在函数内部是作为切片来解析的
//可变参数可以不填,不填函数内部当作nil切片处理
//可变参数必须是相同的类型(需要不同类型时可以用interface{}类型)

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.\n")
		return
	}
	for _, people := range who {
		fmt.Printf("%s,%s\n", prefix, people)
	}
}
