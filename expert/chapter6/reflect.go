package main

import "fmt"

type Foo struct {
	A int
	B string
	C interface{}
}

func testEqual() {
	//当结构体仅包含简单类型可以用==比较
	foo1 := Foo{A: 1, B: "string", C: 1}
	foo2 := Foo{A: 1, B: "string", C: 1}
	fmt.Printf("foo1=%v,foo2=%v\n", foo1, foo2)
	println(foo1 == foo2)
	//包含interface{}类型 且实际类型不可比较时不能比较
	foo1 = Foo{A: 1, B: "string", C: make([]int, 10)}
	foo2 = Foo{A: 1, B: "string", C: make([]int, 10)}
	fmt.Printf("foo1=%v,foo2=%v\n", foo1, foo2)
	println(foo1 == foo2)
}

//IsEqual interface{}类型 仅当底层类型一致且是可比较类型时可以用==比较
//go 提供了 布尔 数值 字符串等基础类型 它们可以比较
func IsEqual(a, b interface{}) bool {
	return a == b
}
func main() {
	//testEqual()
	println(IsEqual(1, 1))
	println(IsEqual(1, "1"))
	println(IsEqual(make([]int, 1), make([]int, 1)))
}
