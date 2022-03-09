package main

import "fmt"

//StringExam1 字符串长度表示Unicode编码字节数,两个中文字符长度大于2
func StringExam1() {
	var s string
	s = "中国"
	fmt.Printf("len(`中国`)=%d\n", len(s))
}
func main() {
	StringExam1()
	var s string
	fmt.Printf("len(%v)=%d\n", s, len(s))
	//字符串可能为空但不会为nil
	//println(s==nil)//编译报错
	a := "123"
	b := "123"
	//比较的是值
	fmt.Printf("a==b is %v,&a==&b is %v\n", a == b, &a == &b)
}
