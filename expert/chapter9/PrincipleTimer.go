package main

import "time"

//Timer 代表一次定时 时间到来后仅发生一个事件
type Timer struct {
	C <-chan time.Time //管道 上层应用据此接收事件
	r runtimeTimer     //runtime 定时器 系统管理的定时器 对上层不可见
}

//runtimeTimer 存放在数组中 按照when字段堆排序
type runtimeTimer struct {
	tb     uintptr                    //存储当前定时器的数组地址
	i      int                        //存储当前定时器的数组下标
	when   int64                      //当前定时器触发时间
	period int64                      //当前定时器周期性触发间隔
	f      func(interface{}, uintptr) //定时器触发时执行的回调函数
	arg    interface{}                //定时器触发时执行回调函数传递的参数一
	seq    uintptr                    //定时器触发时执行回调函数传递的参数二
	//(该参数只在网络收发场景下使用)
}

func NewTimer(d time.Duration) *Timer {
	c := make(chan time.Time, 1)
	t := &Timer{
		C: c,
		r: runtimeTimer{
			when: when(d),
			f:    sendTime,
			arg:  c,
		},
	}
	startTimer(&t.r)
	return t
}

func startTimer(r *runtimeTimer) {
	//todo 把runtimeTimer写入系统协程的数组中并启动系统协程
}

func when(d time.Duration) int64 {
	//todo 计算下一次触发绝对时间
	panic("")
}

// sendTime 定时器触发时的动作
func sendTime(c interface{}, seq uintptr) {
	select {
	case c.(chan time.Time) <- time.Now():
	default: //Ticker 无法保证之前的数据已被取走 所以定义一个空default 避免阻塞
	}
}
func (t *Timer) Stop() bool {
	return stopTimer(&t.r)
}

func stopTimer(r *runtimeTimer) bool {
	//todo 通知系统协程把该Timer移除 并不会关闭管道
	//Timer已触发返回false 未触发返回true
	panic("")
}

func (t *Timer) Reset(d time.Duration) bool {
	w := when(d)
	active := stopTimer(&t.r)
	t.r.when = w //when变化 在系统协程 runtimeTimer数组中的位置可能变化
	startTimer(&t.r)
	return active
}
