package main

import "unsafe"

type _panic struct {
	argp unsafe.Pointer //defer函数的参数
	arg  interface{}    //panic()的参数
	link *_panic        //嵌套panic时指向前一个panic
	//...
	recovered bool //标记当前panic是否已被recover()恢复
	aborted   bool //标记当前panic是否被中断 defer中的新panic会把原panic标记为aborted
	goexit    bool //标记当前是否为runtime.Goexit()产生的panic Goexit产生的panic不能被recover()恢复
}

// panic链表和defer链表都存储在协程数据结构中
type g struct {
	_panic *_panic
	//_defer *_defer
}
