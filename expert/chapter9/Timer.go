package main

import (
	"fmt"
	"log"
	"time"
)

// WaitChannel
//等待一段事件，如果仍旧没有数据可操作 则判定为超时 转而去处理其他逻辑
func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-conn:
		timer.Stop() //返回值为true代表提前结束 false表示超时后结束
		return true
	case <-timer.C: //超时
		println("WaitChannel timeout!")
		return false
	}
}

//DelayFunction 延迟执行每个方法
func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		log.Println("Delayed 5s,start to do something")
	}
}

func AfterDemo() {
	log.Println(time.Now())
	<-time.After(1 * time.Second) //返回C
	log.Println(time.Now())
}

func AfterFuncDemo() {
	log.Println("AfterFuncDemo start:", time.Now())
	time.AfterFunc(1*time.Second, func() { //1s后执行回调函数 是异步执行的
		log.Println("AfterFuncDemo end:", time.Now())
	})
	time.Sleep(2 * time.Second)
}

func ResetDemo() {
	fmt.Printf("%v\n", time.Now())
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("%v\n", <-timer.C)
	timer.Reset(time.Second * 1) //重置定时器 一般是用于已经超时的Timer
	//先Stop再启动 返回值是Stop的返回值
	fmt.Printf("%v\n", <-timer.C)

}

func main() {
	ResetDemo()
}
