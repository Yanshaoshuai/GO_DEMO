package main

import (
	"fmt"
	"io"
	"os"
)

//Go的defer语句预设一个函数调用（延期的函数），该调用在函数执行defer返回时立刻运行。该方法显得不同常规，但却是处理上述情况很有效，无论函数怎样返回，都必须进行资源释放。

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	//被延期的函数以后进先出（LIFO）的顺行执行，因此以上代码在返回时将打印4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
