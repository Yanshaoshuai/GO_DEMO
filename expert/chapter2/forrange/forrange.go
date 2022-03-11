package main

import (
	"fmt"
	"sync"
	"time"
)

//ForExpression 传统for循环
func ForExpression() {
	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		fmt.Printf("index=%d,value=%d\n", i, s[i])
	}
}

// ForRangeExpression for range
func ForRangeExpression() {
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}
}

//FindMonkey 字符串赋值会有内存拷贝 不如下标快
func FindMonkey(s []string) bool {
	for _, v := range s {
		if v == "monkey" {
			return true
		}
	}
	return false
}

func PrintSlice() {
	s := []int{1, 2, 3}
	var wg sync.WaitGroup
	wg.Add(len(s))
	for _, v := range s {
		go func() { //开启线程花费时间比循环长,而每次循环都会给v重新赋值 所以最后输出三个3
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
func RangeNilChannel() {
	var c chan string
	//读nil channel会死锁
	for v := range c {
		println(v)
	}
}

func RangeTimer() {
	t := time.NewTimer(time.Second)

	//1 s后 打印 hi 然后陷入阻塞
	for _ = range t.C {
		println("hi")
	}
}

func RangeDemo() {
	s := []int{1, 2, 3}
	//会正常结束 循环次数在遍历前已经确定
	for i := range s {
		s = append(s, i)
	}
	fmt.Printf("%v", s)
}
func main() {
	ForExpression()
	PrintSlice()
	//RangeNilChannel()
	//RangeTimer()
	RangeDemo()
}
