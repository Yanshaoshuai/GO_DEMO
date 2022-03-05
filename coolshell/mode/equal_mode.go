package main

import (
	"fmt"
	"reflect"
)

type data struct {
}

func main() {
	//需要比较两个结构体中的数据是否相同时，我们需要使用深度比较，而不是只是简单地做浅度比较。这里需要使用到反射 reflect.DeepEqual()
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true
}
