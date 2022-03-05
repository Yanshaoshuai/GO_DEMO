package main

import "fmt"

type Person struct {
	Name   string
	Sexual string
	Age    int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

// Print 成员函数 接收者是*Person
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func main() {
	//在 Go 语言中，使用“成员函数”的方式叫“Receiver”，这种方式是一种封装，因为 PrintPerson()本来就是和 Person强耦合的，所以，理应放在一起
	var p = Person{
		Name:   "Hao Chen",
		Sexual: "Male",
		Age:    44,
	}

	PrintPerson(&p)
	p.Print()
}
