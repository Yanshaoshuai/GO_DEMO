package main

//type Ticker struct {
//	C <-chan time.Time //上层据此管道接收事件
//	r runtimeTimer//交由系统管理的定时器
//}

//func NewTicker(d time.Duration) *time.Ticker {
//	if d<=0{
//		panic(errors.New("non-positive interval for NewTicker"))
//	}
//	c:=make(chan time.Time,1)
//	t:=&time.Ticker{
//		C:c,
//		r:runtimeTimer{
//			when:when(d),
//			period:int64(d),//Ticker和Timer的重要区别就是提供了period参数,
//							//据此决定Timer是一次性的还是周期性的
//			f:sendTime,
//			arg:c,
//		},
//	}
//	startTimer(&t.r)
//	return t
//}

//func sendTime(c interface{}, seq uintptr) {
//	select {
//	case c.(chan time.Time) <- time.Now():
//	default: //Ticker 无法保证之前的数据已被取走 所以定义一个空default 避免阻塞
//如果管道中有数据没有被取走，那么走default分支，本次事件丢失
//	}
//}
