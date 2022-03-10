package main

import "unsafe"

// map结构
type hmap struct {
	count      int            //当前保存的元素个数
	B          uint8          //bucket数组大小
	buckets    unsafe.Pointer //buckets数组,数组的长度为2^B
	oldbuckets unsafe.Pointer //老旧bucket数组,用于扩容
	//...
}

//buckets结构
//data 和 overflow并没有显式地在结构体中声明 在访问bucket时直接通过指针的偏移来访问这些虚拟成员
type bmap struct {
	tophash  [8]uint8 //存储Hash值的高8位
	data     []byte   //key value 数据:key/key/key/.../value/value/value/...
	overflow *bmap    //溢出bucket的地址
}
