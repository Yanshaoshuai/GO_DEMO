package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string //public
	age  int    //仅包内可访问
}

func (p *People) SetName(name string) {
	p.Name = name
}
func (p *People) SetAge(age int) {
	p.age = age
}
func (p *People) GatAges() int {
	return p.age
}

type Kid struct {
	Name string
	Age  int
}

// SetName 非指针接收者改变不了对象内容
func (k Kid) SetName(name string) {
	k.Name = name
}
func (k *Kid) SetAge(age int) {
	k.Age = age
}

type Fruit struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"` //omitempty空则忽略
	From  string //不指定 用字段名序列/反序列化
	Kind  string `json:"kind,omitempty"`
}

func main() {
	p := People{"yan", 20}
	fmt.Printf("%v\n", p)

	kid := Kid{"xiaoming", 22}
	kid.SetName("daming")
	kid.SetAge(64)
	fmt.Printf("%v\n", kid)

	var f = Fruit{Name: "apple", Color: "red"}
	var s = `{"Name":"banana","Weight":100}`
	fstr, err := json.Marshal(f)
	if err == nil {
		fmt.Printf("%s\n", fstr)
	}

	err = json.Unmarshal([]byte(s), &f)
	if err == nil {
		fmt.Printf("%v", f)
	}
}
