package main

import (
	"errors"
	"fmt"
	"os"
)

type MyError struct {
	err string
}

// error接口只有一个 Error()方法
//任意实现了此方法的结构体都实现了error接口
//type error interface {
//	Error() string
//}
func (receiver *MyError) Error() string {
	return receiver.err
}

type errorString struct {
	s string
}

//error 内部的一个实现
func (e *errorString) Error() string {
	return e.s
}

//New 模拟errors.New 方法
func New(text string) error {
	return &errorString{text}
}

func Errorf(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a))
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (pathError PathError) Error() string {
	return pathError.Op + " " + pathError.Path +
		": " + pathError.Err.Error()
}

func AssertError(err error) {
	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("it's an os.PathError,operation:%s,path:%s,msg:%v\n", e.Op, e.Path, e.Err)
	}
}

func ExampleAssertError() {
	err1 := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}
	AssertError(err1)
}
func main() {
	//检查error
	var err error
	if err != nil {
		//something went wrong
	}
	//比较异常
	err1 := errors.New("permission denied")
	if err1 == os.ErrPermission {
		fmt.Printf("permission denied")
	}
	//类型断言
	AssertError(err1)
	ExampleAssertError()
	//传递error
	//添加上下文到error中传递
	//不推荐
	ExampleWriteFile()
	//自定义error上下文信息和原error分开存放 见PathError
}

func ExampleWriteFile() {
	err := writeFile("a.txt")
	if err == os.ErrPermission { //无法判断
		fmt.Printf("permission denied\n")
	}
}
func writeFile(fileName string) error {
	if fileName == "a.txt" {
		return fmt.Errorf("write file error:%v", os.ErrPermission)
	}
	return nil
}
