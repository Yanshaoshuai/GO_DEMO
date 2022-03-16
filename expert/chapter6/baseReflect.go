package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Speak() string
}
type Cat struct {
}

func (c Cat) Speak() string {
	return "cat"
}

//实现了接口所有方法的结构体也就实现了接口
//实现了接口的结构体示例可以存储在接口类型变量中
var animal Animal = Cat{}

//接口类型的变量在存储某个变量时会同时保存变量类型和变量值
//这是interface变量可以存储任意实现了该接口类型的变量的原因
//Go的反射就是在运行时操作interface中的值和类型的特性
//type iface struct {
//	tab *tab//保存变量类型
//	data unsafe.Pointer//变量值位于堆栈的指针
//}

//空interface interface{}没有方法
//所以此类型的变量可以存储任何值

func foo() {
	var A interface{}
	A = 100
	v := reflect.ValueOf(A) //获取值
	B := v.Interface()      //获取interface类型变量
	if A == B {
		fmt.Printf("They are same!\n")
	}
}
func main() {
	//变量==>反射对象
	var x float64 = 3.4
	t := reflect.TypeOf(x) //获取变量类型
	fmt.Printf("type:%v\n", t)

	v := reflect.ValueOf(x) //获取变量值
	fmt.Printf("value:%v\n", v)

	//反射对象==>变量
	foo()

	//通过反射设置value
	//errorWay()
	rightWay()
}

func rightWay() {
	var x float64 = 3.4
	//传地址
	v := reflect.ValueOf(&x)
	//ValueOf返回的是unaddressable value 不可修改
	//v.SetFloat(7.1)
	var elem reflect.Value = v.Elem() //获取Value
	fmt.Printf("has address value=%v\n", elem)
	elem.SetFloat(7.1)
	fmt.Printf("x=%f\n", x)
	fmt.Printf("elem.Interface{}=%v\n", elem.Interface())
}

func errorWay() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1)
	fmt.Printf("x=%f\n", x)
}
