package main

import (
	"fmt"
	"os"
	"sync"
)

//RuleDefer defer规则
//defer只能作用于函数和函数调用
func RuleDefer() {
	//作用于函数
	defer fmt.Println("hello world")
	//作用于函数调用
	defer func() {
		fmt.Println("hello")
	}()
}

//DeferCloseFile close file
func DeferCloseFile() error {
	file, err := os.Open("./defer.go")
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
func DeferUnlock() {
	var m sync.Mutex
	m.Lock()
	defer m.Unlock()
}

//DeferWait 等待协程退出
func DeferWait() {
	var wg sync.WaitGroup
	defer wg.Wait()
}
func DeferRecover() {
	defer func() {
		//recover只能用在defer中
		recover()
	}()
	panic("")
}

//defer 无法改变匿名返回值
//返回前会把i赋值给一个匿名变量
//所以操作i不会改变匿名返回值
func foo() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

//规则
//延迟函数的参数在defer语句出现时已经确定
//延迟函数按照后进先出的规则执行
//延迟函数可能操作所在主函数的具名返回值
func main() {
	RuleDefer()
	i := foo()
	println(i)
}
