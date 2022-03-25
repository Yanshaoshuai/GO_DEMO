package main

//每个defer语句对应一个_defer实例
//多个实例使用link链接起来形成一个单链表
//保存到goroutine数据结构中
//type _defer struct {
//	sp unintptr //函数栈指针
//	pc unintptr //程序计数器
//	fn *funcval //函数地址
//	link *_defer //指向自身结构的指针 用于链接多个defer
//}

//heap-allocated 存储在堆上的defer
//替换成deferproc()
//stack-allocated 存储在栈上的defer
//编译器尽可能把defer编译成栈类型
//替换成 deferprocStack()
//open-coded 开放编码类型defer
//延迟函数被直接插入函数尾部

//defer不能被处理成开放编码类型的场景
//编译时禁用了编译器优化,即-gcflags="-N -l"
//defer出现在循环语句中
//单个函数中defer出现了8个以上,或者return语句的个数和defer语句的个数乘积超过了15个
