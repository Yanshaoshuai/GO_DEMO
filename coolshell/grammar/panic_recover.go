package main

import "fmt"

func g(i int) {
	if i > 1 {
		fmt.Println("Panic!")
		panic(fmt.Sprintf("%v", i))
	}

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("Calling g with ", i)
		g(i)
		fmt.Println("Returned normally from g.")
	}
}

//当panic被调用时，它将立即停止当前函数的执行并开始逐级解开函数堆栈，同时运行所有被defer的函数。如果这种解开达到堆栈的顶端，程序就死亡了。
//但是，也可以使用内建的recover函数来重新获得Go程的控制权并恢复正常的执行。
//对recover的调用会通知解开堆栈并返回传递到panic的参量。由于在解开期间运行的代码仅处在被defer的函数之内，recover仅在被延期的函数内部才是有用的。

func main() {
	f()
	fmt.Println("Returned normally from f.")
}
