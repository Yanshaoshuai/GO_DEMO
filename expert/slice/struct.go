package main

import (
	"fmt"
	"unsafe"
)

type sliceStruct struct {
	array unsafe.Pointer //指向底层数组的指针
	len   int            //长度
	cap   int            //容量
}

func main() {
	slice := make([]int, 5, 10)
	println(slice[0])
	println(slice[4])
	//println(slice[5])//cap=10 len=5 只能使用slice[0]~slice[4]的元素
	arr := [10]int{}
	sliceFromArr := arr[5:7]
	println(sliceFromArr[0])
	println(sliceFromArr[1])
	//println(sliceFromArr[2])//cap=10-5=5 len=7-5=2 只能使用sliceFromArr[0]~sliceFromArr[1]之间的元素
	fmt.Printf("sliceFromArr cap=%d len=%d\n", cap(sliceFromArr), len(sliceFromArr))

	//扩容
	sliceScale := make([]int, 10)      //cap=10 len =10
	sliceScale = append(sliceScale, 1) //capOld=10<1024 cap=2*capOld=20
	fmt.Printf("sliceScale cap=%d len=%d\n", cap(sliceScale), len(sliceScale))
	sliceScaleLarge := make([]int, 1024)         //cap=1024 len=1024
	sliceScaleLarge = append(sliceScaleLarge, 1) //capOld=1024>=1024 cap=1.25*1024=1280
	fmt.Printf("sliceScaleLarge cap=%d len=%d\n", cap(sliceScaleLarge), len(sliceScaleLarge))

	//拷贝
	sliceShort := []int{1, 2, 3, 4, 5} //cap=len=5
	fmt.Printf("sliceShort cap=%d len=%d\n", cap(sliceShort), len(sliceShort))
	sliceLang := make([]int, 10) //cap=len=10
	fmt.Printf("sliceLang cap=%d len=%d\n", cap(sliceLang), len(sliceLang))
	copy(sliceShort, sliceLang) //拷贝长度取决于较短的 不会发生扩容
	fmt.Printf("sliceShort cap=%d len=%d\n", cap(sliceShort), len(sliceShort))
	fmt.Println(sliceShort)
}
