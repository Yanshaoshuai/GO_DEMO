package main

import (
	"errors"
	"fmt"
	"os"
)

//ExampleUnwrapLoop 逐层unwrap 获取原始error
func ExampleUnwrapLoop() {
	err1 := fmt.Errorf("write file error:%w", os.ErrPermission)
	err2 := fmt.Errorf("write file error:%w", err1)
	err := err2
	for {
		if err == os.ErrPermission {
			fmt.Printf("permission denied\n")
			break
		}
		if err = errors.Unwrap(err); err == nil {
			break
		}
	}
}
func Is(err, target error) bool {
	for {
		if err == target {
			return true
		}
		//优先调用自己的Is方法
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}
		//没有自己的Is方法会循环调用unwrap
		if err = errors.Unwrap(err); err == nil {
			return false
		}
	}
}

//ExampleIs 使用Is判断error chain是否有某种类型error
func ExampleIs() {
	err1 := fmt.Errorf("write file error:%w", os.ErrPermission)
	err2 := fmt.Errorf("write file error:%w", err1)
	if errors.Is(err2, os.ErrPermission) {
		fmt.Printf("permission denied\n")
	}
}

func ExampleAssertChainErrorWithoutAs() {
	err := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}
	err2 := fmt.Errorf("some context: %w", err)
	if e, ok := err2.(*os.PathError); ok { //多层wrap 转换失败
		fmt.Printf("it is an os.PathError,operation:%s,"+
			"path:%s,msg:%v\n", e.Op, e.Path, e.Err)
	}
}

//ExampleAssertChainWithAs As会递归调用unwrap 并尝试把值赋给target
func ExampleAssertChainWithAs() {
	err := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}
	err2 := fmt.Errorf("some context: %w", err)
	var target *os.PathError
	if errors.As(err2, target) {
		fmt.Printf("it is an os.PathError,operation:%s,"+
			"path:%s,msg:%v\n", target.Op, target.Path, target.Err)
	}
}
func main() {
	ExampleUnwrapLoop()
	ExampleIs()
	ExampleAssertChainErrorWithoutAs()
}
