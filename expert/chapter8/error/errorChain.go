package main

import (
	"errors"
	"fmt"
)

type wrapError struct {
	msg string //存储上下文信息和err.Error()
	err error  //存储原error
}

func (e wrapError) Error() string {
	return e.msg
}

//Unwrap 返回原error
func (e *wrapError) Unwrap() error {
	return e.err
}

func ExampleCreateBasicError() {
	err := errors.New("this is a error")
	basicErr := fmt.Errorf("some context:%v", err)
	if _, ok := basicErr.(interface{ Unwrap() error }); !ok { //如果没有实现Unwrap接口
		fmt.Println("basicError is a errorString")
	}
}

func ExampleCreateWrapError() {
	err := errors.New("this is demo error")
	wrapError := fmt.Errorf("some context:%w", err) //有%w的Errorf返回的是wrap error
	//%w只能匹配error类型
	//每次只能接受一个%w
	if _, ok := wrapError.(interface{ Unwrap() error }); ok { //如果实现了Unwrap接口
		fmt.Println("wrapError is a wrapError")
	}
}

func Unwrap(err error) error {
	u, ok := err.(interface{ Unwrap() error }) //检查是否实现了Unwrap方法
	if !ok {
		return nil
	}
	return u.Unwrap()
}

//PathError113 自定义链式error 需要实现Unwrap方法
type PathError113 struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError113) Error() string {
	return e.Op + " " + e.Path +
		": " + e.Err.Error()
}

//Unwrap 返回原始error
func (e *PathError113) Unwrap() error {
	return e.Err
}

func main() {
	ExampleCreateBasicError()
	ExampleCreateWrapError()
}
