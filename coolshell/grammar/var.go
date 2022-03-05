package main

//常量
const s string = "hello world"
const pi float32 = 3.1415926

func main() {
	//声明初始化一个变量
	var x int = 100
	//x := 100 //等价于 var x int = 100;
	var str string = "hello world"
	//声明初始化多个变量
	var i, j, k int = 1, 2, 3

	//不用指明类型，通过初始化值来推导
	var b = true //bool型
	println(x, str, i, j, k, b)
}
