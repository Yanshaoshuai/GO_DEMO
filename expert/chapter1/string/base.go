package main

import "fmt"

func main() {
	//声明一个空字符串变量再赋值
	var s1 string
	s1 = "Hello World"
	_ = s1
	//使用简短变量声明
	s2 := "Hello World"
	_ = s2

	//转义
	str1 := "hi!\nHallo~\nYou mean is \"Hello?\"\nYes"
	println(str1)
	str2 := `hi!
Hallo~
You mean is "Hello?"
Yes`
	println(str2)
	fmt.Printf("str1==str2 is %v\n", str1 == str2)

	//字符串拼接
	s := "s"
	//字符串拼接时会触发内存分配及内存拷贝,单行语句拼接多个字符串只分配一次内存
	s = s + "a" + "b"

	//字节切片string互转
	//互转会发生一次内存拷贝，有一定开销
	b := []byte{'H', 'e', 'l', 'l', 'o'}
	str := string(b)
	fmt.Printf("str=%s,b=%s\n", str, b)
	b = []byte(str)
	fmt.Printf("str=%s,b=%s\n", str, b)

	//中文遍历下标是不连续的 字符串下标标记的是字节位置不是字符位置
	scn := "中国"
	//遍历中文每次返回的index是首字节下标
	for index, value := range scn {
		fmt.Printf("%d,%c\n", index, value)
	}

	//不可以通过下标修改字符串
	stringVar := "Hello"
	println(&stringVar)
	//println(&stringVar[0])//不能取元素地址
	//&stringVar[0]=byte(104)//非法
	stringVar = "hello" //赋予新值
	println(&stringVar)
}
