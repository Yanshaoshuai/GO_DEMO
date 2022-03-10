package main

import "fmt"

func main() {
	//simple
	//数组 0<=low<=high<=len(a)
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:4] //low=1 high=4 cap=cap(a)-low=4 len=high-low=3
	fmt.Printf("len=%d,cap=%d\n", len(b), cap(b))
	//切片 0<=low<=high<=cap(baseSlice)
	baseSlice := make([]int, 0, 10) //len=0 cap =10
	newSlice := baseSlice[2:5]      //low=2>len
	fmt.Printf("newSlice:%v\n", newSlice)

	//扩展表达式  0<=low <=high <=max <=cap(a)
	//cover 示例
	cover()
	//避免覆盖
	avoidCover()

	//省略写法
	slice := []int{1, 2, 3}
	slice2 := slice[:]
	fmt.Printf("%v\n", slice2)
	//拓展表达式只能省略low
	sli := []int{1, 2, 3, 4, 5}
	sli2 := sli[:2:4]
	fmt.Printf("%v", sli2)

}
func cover() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a=%v\n", a)
	b := a[1:4]
	b = append(b, 0) //覆盖了a[4]
	fmt.Printf("a=%v\n", a)
}
func avoidCover() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a=%v\n", a)
	b := a[1:4:4]    //cap=max-low=3
	b = append(b, 0) //重新扩容 没有覆盖a[4]
	fmt.Printf("a=%v\n", a)
}
