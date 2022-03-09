package main

import "fmt"

type Priority int

const (
	LOG_EMERG Priority = 1
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)
const (
	mutexLocked           = 1 << iota
	mutexWoken            //1<<1
	mutexStarving         //1<<2
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

//iota代表const声明块的行索引(指格式化后的,从0开始)
//如果为常量指定了一个表达式,后面没有表达式的常量继承上面的表达式(该表达式必须包含iota)
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1
	bit1, mask1 //1<<1,1<<1 -1
	_, _
	bit3, mask3 //1<<3,1<<3 -1
)
const (
	ZERO       = 0
	lineNumber = iota
)

func main() {
	fmt.Printf("LOG_EMERG=%d,LOG_ALERT=%d,LOG_CRIT=%d,LOG_ERR=%d,LOG_WARNING=%d,LOG_NOTICE=%d,LOG_INFO=%d,LOG_DEBUG=%d\n",
		LOG_EMERG, LOG_ALERT, LOG_CRIT, LOG_ERR, LOG_WARNING, LOG_NOTICE, LOG_INFO, LOG_DEBUG)

	fmt.Printf("mutexLocked=%d,mutexWoken=%d,mutexStarving=%d,mutexWaiterShift=%d,starvationThresholdNs=%f\n",
		mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs)

	fmt.Printf("bit0,mask0=%d,%d,bit1,mask1=%d,%d,bit3,mask3=%d,%d\n",
		bit0, mask0, bit1, mask1, bit3, mask3)

	fmt.Printf("lineNumber=%d\n", lineNumber)
}
