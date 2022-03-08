package main

import "fmt"

type Student struct {
	Name string
}

// SetName 非指针接收者改变不了结构体内部值
func (s Student) SetName(name string) {
	s.Name = name
}
func (s *Student) UpdateName(name string) {
	s.Name = name
}
func main() {
	s := Student{}
	s.SetName("Rainbow")
	fmt.Printf("Name: %s\n", s.Name)

	s.UpdateName("Rainbow")
	fmt.Printf("Name: %s\n", s.Name)
}
