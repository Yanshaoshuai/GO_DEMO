package main

import "fmt"

type Animal interface {
	eat()
	sleep()
}
type Cat struct {
	name string
}

func (c Cat) eat() {
	println("Cat eat")
}

func (c Cat) sleep() {
	println("Cat sleep")
}
func (cat Cat) printName() {
	println(cat.name)
}
func main() {
	var cat = Cat{}
	fmt.Printf("%v \n", cat)
	cat.sleep()
	cat.name = "tom"
	cat.printName()
}
