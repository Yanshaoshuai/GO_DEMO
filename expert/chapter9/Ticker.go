package main

import (
	"log"
	"time"
)

//type Ticker struct {
//	C <-chan time.Time
//	r runtimeTimer
//}

// TickerDemo 周期性任务
func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		log.Println(time.Now())
	}
}
func TickerLaunch() {
	ticker := time.NewTicker(5 * time.Minute)
	maxPassenger := 30 //每车最大装载人数
	passengers := make([]string, 0, maxPassenger)
	for {
		passenger := GetNewPassenger()
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}
		select {
		case <-ticker.C: //到点出发
			Launch(passengers)
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger { //人满出发
				Launch(passengers)
				passengers = []string{}
			}
		}
	}
}

func Launch(passengers []string) {
	//todo 出发
}

func GetNewPassenger() string {
	//todo 上乘客
	panic("")
}

//WrongTicker 每次循环创建一个ticker
//造成资源泄露
func WrongTicker() {
	for {
		select {
		case <-time.Tick(1 * time.Second):
			log.Printf("Resource leak!\n")
		}
	}
}
func main() {
	TickerDemo()
	//WrongTicker()
}
