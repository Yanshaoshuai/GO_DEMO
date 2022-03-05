package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}
type rect struct {
	width, height int
}

//注意：Go语言中没有public, protected, private的关键字，所以，如果你想让一个方法可以被别的包访问的话，你需要把这个方法的第一个字母大写。这是一种约定。

func (r *rect) area() int { //求面积
	return r.width * r.height
}
func (r *rect) perimeter() int { //求周长
	return 2 * (r.width + r.height)
}
func main() {

	//初始化
	person := Person{"Tom", 30, "tom@gmail.com"}
	//带名字的初始化
	person = Person{name: "Tom", age: 30, email: "tom@gmail.com"}

	fmt.Println(person) //输出 {Tom 30 tom@gmail.com}

	pPerson := &person

	fmt.Println(pPerson) //输出 &{Tom 30 tom@gmail.com}

	pPerson.age = 40
	person.name = "Jerry"
	fmt.Println(person) //输出 {Jerry 40 tom@gmail.com}

	//结构体方法
	r := rect{width: 10, height: 15}
	fmt.Println("面积: ", r.area())
	fmt.Println("周长: ", r.perimeter())
	rp := &r
	fmt.Println("面积: ", rp.area())
	fmt.Println("周长: ", rp.perimeter())
}
