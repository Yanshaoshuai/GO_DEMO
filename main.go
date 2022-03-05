package main

import (
	"fmt"
	varible "learn_go/var"
	"unicode/utf8"
)

func main() {
	println(varible.Astr)
	//println(varible.b)
	println("hello world")
	println(`hello " world`)
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))
	sum, plus := varible.GetSumAndPlus(3, 2)
	println(sum, plus)
	_, plus2 := varible.GetSumAndPlus(3, 2)
	println(plus2)
	//数组
	arr := [5]int{9, 8, 7}
	//slice
	sli := []int{9, 8, 7}
	fmt.Printf("arr = %v,len=%v,capcity=%v\n", arr, len(arr), cap(arr))
	fmt.Printf("sli = %v,len=%v,capcity=%v\n", sli, len(sli), cap(sli))
	sli = append(sli, 8)
	fmt.Printf("sli = %v,len=%v,capcity=%v\n", sli, len(sli), cap(sli))
	//slice
	sli2 := make([]int, 0, 10)
	fmt.Printf("sli2 = %v,len=%v,capcity=%v\n", sli2, len(sli2), cap(sli2))
	index := 0
	//控制语句
	for {
		if index == 3 {
			break
		}
		fmt.Printf("index=%d=>%d\n", index, sli[index])
		index++
	}
	index = 0
	for index != 3 {
		fmt.Printf("index=%d=>%d\n", index, sli[index])
		index++
	}
	for i := 0; i < 10; i++ {
		println(i)
	}
	for index, value := range sli {
		println(index, value)
	}
	fruit := "苹果"
	switch fruit {
	case "苹果":
		println("fruit是苹果")
		//break 默认会break
	case "草莓":
		println("fruit是草莓")
	default:
		println("fruit是啥 不知道")
	}

}
