package main

//
//type timer struct {
//	tb *timersBucket//当前定时器寄存于系统timer堆的地址
//	i int//当前定时器寄存于系统timer堆的下标
//	when int64//当前定时器下次触发时间
//	period int64//周期性触发间隔(如果是Timer则为0,表示不重复触发)
//	f func(interface{},uintptr)//定时器触发时执行的函数
//	arg interface{}//定时器触发时执行函数传递的参数一
//	seg uintptr//定时器触发时执行回调函数传递的参数二
//				//(该参数只在网络收发场景下使用)
//}

//系统预留了64个timersBucket(最多64个,GOMAXPROCS小于64则只有GOMAXPROCS个)
//协程ProcessID%64=timersBucket号
//协程数大于64才会出现多个协程的timer在同一个timersBucket中
//同意timersBucket中的timer按照触发时间堆排序形成小头堆(四叉堆)
//据此顺序获取timer在timersBucket.t中的下标
//type timersBucket struct {
//	lock mutex
//	gp *g  //处理堆中事件的协程
//	created bool //事件处理协程是否已创建，默认为false,添加首个定时器时置为true
//	sleeping bool//事件处理协程(gp)是否在睡眠(如果t中有定时器,那么还未到触发的时间，gp会进入睡眠)
//	rescheduling bool//事件处理协程(gp)是否已暂停(如果t中定时器均已删除,那么gp会暂停)
//	sleepUntil int64//系统协程睡眠到指定的时间(如果有新的定时任务则可能会提前唤醒)
//	waitnote note//提前唤醒时使用的通知
//	t []*timer//定时器切片
//}

//上述原理适用于go1.0~go1.13版本
//go1.14取消了timersBucket桶,直接把runtimeTimer的堆保存在处理放到了处理器P中
//type P struct {
//	runq [256]gunintptr
//	//...
//	timersLock mutex
//	//...
//	timers []*timer
//}
//取消了timerproc，每次协程调度时检查定时器是否需要触发
//当有定时器需要触发时先处理定时器再调度协程
