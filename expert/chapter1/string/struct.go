package main

import "unsafe"

type stringStruct struct {
	str unsafe.Pointer //字符串首地址
	len int            //字符串长度
}

func findnull(str *byte) int {
	return -1
}

//字符串构造过程
//比如:
//var str string
//str="Hello World"
func gostringnocopy(str *byte) string {
	//先构建stringStruct
	ss := stringStruct{str: unsafe.Pointer(str), len: findnull(str)}
	//再将stringStruct转成string
	s := *(*string)(unsafe.Pointer(&ss))
	return s
}

//字符串拼接过程
//多个字符串拼接只涉及一次内存分配
//性能主要在内存拷贝
func concatstrings(a []string) string {
	length := 0
	for _, str := range a {
		length += len(str)
	}
	//分配内存 返回一个string和切片 二者共享内存空间
	s, b := rawstring(length)
	//string 无法修改 只能通过切片修改
	for _, str := range a {
		copy(b, str) //多次拷贝
		b = b[len(str):]
	}
	return s
}

//生成一个新的string 返回string和与其共享内存的切片
func rawstring(size int) (s string, b []byte) {
	//p:=malloocgc(unitptr(size),nil,false)
	//stringStructOf(&s).str=p
	//stringStructOf(&s).len=size
	//*(*slice)(unsafe.Pointer(&b))=slice(p,size,size)
	//return
	return s, b
}

//[]byte转string 伪代码
//buf固定大小
//func slicebytetostring(buf *tmpBuf, b []byte) (str string) {
//	var p unsafe.Pointer
//	if buf!=nil && len(b)<=len(buf){
//如果预留buf够用 则用预留buf
//	p=unsafe.Pointer(buf)
//}else {
//否则重新申请内存
//	p=mallocgc(unitptr(len(b),nil,false))
//}
//stringStructOf(&str).str=p
//stringStructOf(&str).len=len(b)
//将切片底层数组中的数据拷贝到字符串
//	memmove(*(*slice)(unsafe.Pointer(&b)).array,unitptr(len(b)))
//	return
//}

//string转[]byte伪代码
//func stringtoslicebyte(buf *tempBuf, s string) []byte {
//	var b []byte
//	if buf!=nil && len(s)<=len(buf){
//		*buf=tmpBuf{}
//从预留buf中切出新的切片
//	b=buf[:len(s)]
//}else {
//生成新的切片
//		b=rawbyteslice(len(s))
//	}
//	copy(b,s)
//	return b
//}
//临时需要字符串的情况下,byte切片专场string不会拷贝内存
//而是直接返回一个string，这个string的指针指向切片的内存
//包含如下场景
//1.使用m[string(b)]来查找map
//2.字符串拼接: ...+string(b)+...
//3.字符串比较: string(a)=="foo"
