package main

func main() {
	//for range会被编译器转成 for循环
	//对于不同数据类型 会有一些差异

	//作用于数组 伪代码
	//len_temp:=len(range)
	//range_temp:=range
	//for index_temp := 0; index_temp < len_temp; index_temp++ {
	//	value_temp=range_temp[index_temp]
	//	index=index_temp
	//	value=value_temp
	//	original body
	//}

	//作用于切片伪代码
	//for_temp:=range
	//len_temp:=len(range)
	//for index_temp := 0; index_temp < len_temp; index_temp++ {
	//	value_temp=for_temp[index_temp]
	//	index=index_temp
	//	value=value_temp
	//	original body
	//}

	//作用于string伪代码
	//len_temp:=len(range)
	//var next_index_temp int
	//for index_temp := 0; index_temp < len_temp; index_temp=next_index_temp {
	//	value_temp=rune[index_temp]
	//	if value_temp<utf8.RuneSelf{//用于判断Unicode编码是否仍为自身
	//		next_index_temp=index_temp+1
	//	}else {
	//		value_temp,next_index_temp=decoderune(range,index_temp)
	//	}
	//	index=index_temp
	//	value=value_temp
	//	original body
	//}

	//作用于map
	//var hiter map_iteration_struct
	//for mapiterinit(type,range,&hiter);hiter.key!=nil;mapiternext(&hiter){
	//	index_temp=*hiter.key
	//	value_temp=*hiter.val
	//	index=index_temp
	//	value=value_temp
	//	original body
	//}

	//作用于channel
	//for  {
	//	index_temp,ok_temp=<-range
	//	if !ok_temp{//如果ok_temp说明channel中没有数据且已关闭
	//		break
	//	}
	//	index=index_temp
	//	original body
	//}
}
