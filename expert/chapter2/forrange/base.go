package main

import (
	"fmt"
	"time"
)

func RangeArray() {
	a := [3]int{1, 2, 3}
	//返回下标和值 可作用于数组指针 效果等同
	for i, v := range a {
		fmt.Printf("index:%d,value:%d\n", i, v)
	}
}

func RangeSlice() {
	s := []int{1, 2, 3}
	//返回下标和值 不可作用于切片指针
	for i, v := range s {
		fmt.Printf("index:%d,value:%d\n", i, v)
	}
}

func RangeString() {
	s := "Hello"
	for i, v := range s {
		fmt.Printf("index:%d,value:%c\n", i, v)
	}
	//遍历 非纯ascii码字符串 下标不连续
	s = "中国"
	//v是rune类型代表一个unicode编码的一个字节
	for i, v := range s {
		fmt.Printf("index:%d,value:%c\n", i, v)
	}
}

func RangeMap() {
	m := map[string]string{"animal": "monkey", "fruit": "apple"}
	//map本身没有顺序概念，多次遍历可能顺序不一样
	for key, value := range m {
		fmt.Printf("key is %s,value is %s\n", key, value)
	}
}

func RangeChannel() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	time.AfterFunc(time.Second, func() {
		close(c)
	})
	//阻塞直到channel关闭
	//只会返回一个值
	for e := range c {
		fmt.Printf("element is %s\n", e)
	}
}
func main() {
	RangeArray()
	RangeSlice()
	RangeString()
	RangeMap()
	RangeChannel()
}
