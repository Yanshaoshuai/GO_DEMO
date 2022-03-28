package main

import (
	"time"
)

func main() {
	//time
	timeLayout := time.RFC3339
	now := time.Now()
	println(now.String())
	formatTime := now.Format(timeLayout)
	println(formatTime)
	println(now.Format("2006-01-02 15:04 05"))
}
