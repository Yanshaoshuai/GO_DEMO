package main

import (
	"errors"
	"fmt"
	"reflect"
)

func EmptyError() {
	err := errors.New("")
	if err != nil {
		fmt.Printf("empty error still is an error\n")
	}
}

func ErrorCompare() {
	err := errors.New("not found")
	//Errorf %v 会调用Error()方法获取err字符串
	//Errorf %w 会调生成wrapError实例
	err1 := fmt.Errorf("some context:%v", err)
	err2 := fmt.Errorf("some context:%w", err)
	//false
	fmt.Println(reflect.TypeOf(err1) == reflect.TypeOf(err2))
	fmt.Println(reflect.TypeOf(err1))
	fmt.Println(reflect.TypeOf(err2))
	fmt.Printf("err1.Error()=%v\n", err1.Error())
	fmt.Printf("err2.Error()=%v\n", err2.Error())
}

func UnwrapError() {
	err1 := errors.New("not found")
	//Unwrap遇到无法拆解的error时返回nil
	err2 := errors.Unwrap(err1)
	//false
	fmt.Println(err1 == err2)
	//err1=not found
	//err2=<nil>
	fmt.Printf("err1=%v\n", err1)
	fmt.Printf("err2=%v\n", err2)
}
func main() {
	EmptyError()
	ErrorCompare()
	UnwrapError()
}
