package main

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}
func (c City) ToString() string {
	return "City = " + c.Name
}

func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}
func main() {
	//使用了一个叫Stringable 的接口，用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了
	//这就是面向对象编程方法的黄金法则——“Program to an interface not an implementation”
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}
