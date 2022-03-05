package main

import "encoding/json"

type Student struct {
	Name string
	Age  int32
}

func (receiver Student) GetName() string {
	return receiver.Name
}
func (receiver *Student) SetName(newName string) {
	receiver.Name = newName
}
func (receiver Student) GetAge() int32 {
	return receiver.Age
}
func (receiver *Student) SetAge(newAge int32) {
	receiver.Age = newAge
}
func main() {
	var student = Student{}
	studentStr, err := json.Marshal(student)
	if err != nil {
		println(err)
	}
	println(string(studentStr))
	student.SetName("yanshaoshuai")
	studentStr, err = json.Marshal(student)
	if err != nil {
		println(err)
	}
	println(string(studentStr))
	student.SetAge(25)
	studentStr, err = json.Marshal(student)
	if err != nil {
		println(err)
	}
	println(string(studentStr))
}
