package main

import (
	"fmt"
)

func main() {
	arr := [10]int{} //this is an array
	slice := []int{} //this is an array
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", slice)
	fmt.Println("arr:", arr)
	fmt.Println("slice:", slice)
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
	//make创建
	slice1 := make([]int, 12)    //len=cap=12
	slice2 := make([]int, 3, 12) //len=3 cap=12
	fmt.Printf("slice1 len=%d,slice1 cap=%d\n", len(slice1), cap(slice1))
	fmt.Printf("slice2 len=%d,slice2 cap=%d\n", len(slice2), cap(slice2))
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	//切取
	array := [5]int{1, 2, 3, 4, 5}
	sliceFromArr := array[0:2] //cap=5 len=high-low=2
	fmt.Printf("sliceFromArr=%v len=%d cap=%d\n", sliceFromArr, len(sliceFromArr), cap(sliceFromArr))
	sliceFromSlice := sliceFromArr[0:1] //cap=5 len=high-low=1
	fmt.Printf("sliceFromSlice=%v len=%d cap=%d\n", sliceFromSlice, len(sliceFromSlice), cap(sliceFromSlice))
	//new
	sliceFromNew := *new([]int) //nil
	fmt.Printf("sliceFromNew=%v len=%d cap=%d\n", sliceFromNew, len(sliceFromNew), cap(sliceFromNew))
	println(sliceFromNew == nil)
}
