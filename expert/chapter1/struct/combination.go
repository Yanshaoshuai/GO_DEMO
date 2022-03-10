package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

type Cat struct {
	Animal
}
type Dog struct {
	a Animal
}

func main() {
	cat := Cat{} //初始化不可以直接指定被组合对象内部字段
	cat.SetName("cat")
	fmt.Printf("cat.Animal.Name=%s\n", cat.Name) //隐式组合可以直接使用被组合对象的字段和方法

	dog := Dog{}
	dog.a.SetName("dog")
	fmt.Printf("dog.Animal.Name=%s\n", dog.a.Name)
}
