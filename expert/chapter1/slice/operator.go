package main

import "fmt"

func main() {
	//len 和 cap 均为结构体存储字段 len() cap()为O(1)操作
	s := make([]int, 0) //cap=len=0
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1) //扩容 cap=2 len=1
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 2, 3, 4) //扩容 cap=4 len=4
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, []int{5, 6}...) //扩容 cap=8 len=6
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	fmt.Println(s)
}
