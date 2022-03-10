package main

// case数据结构
//type scase struct {
//	c *hchan //操作的管道 决定了每个case语句只能处理一个管道,case不处理管道会编译出错
//	kind uint16//case 类型
//	elem unsafe.Pointer//data element 从管道读取的数据的地址或者将写入管道数据的地址
//}
//kind取值
const (
	caseNil     = iota //管道值为nil
	caseRecv           //读管道的case
	caseSend           //写管道的case
	caseDefault        //default是特殊的case 不会操作管道
)

//selectgo()实现要点
//1.通过随机函数fastrandn()将原始的case顺序打乱,在遍历各个case时使用打乱后的顺序就会表现出随机性
//2.循环遍历各个case时，如果发现某个case已就绪,则直接跳出循环进行管道操作并返回
//3.循环遍历各个case时,循环能正常结束(没有跳转),说明所有case都没有就绪，如果有default语句则命中default
//4.如果所有case都未命中且没有default,selectgo()将阻塞等待所有管道,任一管道就绪后，都将开始新的循环
