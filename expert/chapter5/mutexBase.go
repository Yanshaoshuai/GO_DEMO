package main

type Mutex struct {
	//内部实现把该变量分成四份
	//Waiter 29bit 阻塞等待锁的协程个数
	//Starving 1bit 表示是否处于饥饿状态 0 没有饥饿 1 饥饿 说明有协程阻塞超过1ms
	//Woken 1bit 表示是否有协程已被唤醒 0表示没有协程唤醒 1表示已有协程唤醒 正在加锁过程中
	//Locked 1bit 表示该Mutex是否已被锁定,0表示没有锁定 1表示已被锁定
	state int32  //是否被锁定
	sema  uint32 //信号量
}

type RWMutex struct {
	w           Mutex  //用于控制多个写锁 获得写锁首先要获取该锁
	writerSem   uint32 //写阻塞等待的信号量 最后一个读者释放锁时会释放该信号量
	readerSem   uint32 //读阻塞的协程等待的信号量 持有写锁的协程释放锁后会释放该信号量
	readerCount int32  //记录读者的个数
	readerWait  int32  //记录写阻塞时读者的个数
}
