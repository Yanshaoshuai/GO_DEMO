package main

import "fmt"

func main() {
	//Golang的切片是共享内存的，也就是说，没有数据的复制，只是记录从哪切到哪的信息
	//扩容后会重新分配内存
	a := [5]int{1, 2, 3, 4, 5}

	b := a[2:4] // a[2] 和 a[3]，但不包括a[4]
	fmt.Println(b)

	b = a[:4] // 从 a[0]到a[4]，但不包括a[4]
	fmt.Println(b)

	b = a[2:] // 从 a[2]到a[4]，且包括a[2]
	fmt.Println(b)
}
