package main

import (
	"fmt"
	"reflect"
)

func main() {
	baseStr := "Hello World!"
	fmt.Printf("baseStr:%s\n", baseStr)
	fmt.Printf("baseStr type: %s\n", reflect.TypeOf(baseStr))

	newStr := baseStr[0:5] //字符串切片返回的仍然是字符串
	fmt.Printf("baseStr:%s\n", newStr)
	fmt.Printf("baseStr type: %s\n", reflect.TypeOf(newStr))

	fmt.Printf("%c", newStr[0])
	//newStr[1]='a'//不能单独给字符串某个元素赋值
}
