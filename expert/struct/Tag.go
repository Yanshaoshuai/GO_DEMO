package main

import (
	"fmt"
	"reflect"
)

type TypeMeta struct {
	//字段后面``之间的就是字段的Tag
	// key:"value"组成 key非空字符串 不能包含控制字符 空格 引号 冒号
	//value 以双引号标记的字符串
	Kind       string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}

// struct 结构
//type StructField struct {
//	Name string
//	Type Type
//	Tag reflect.StructTag //StructTag是string的别名
//	//...
//}

//PrintTag 获取并打印Tag
func PrintTag() {
	t := TypeMeta{}
	ty := reflect.TypeOf(t)
	for i := 0; i < ty.NumField(); i++ {
		fmt.Printf("Field: %s,Tag:%s\n", ty.Field(i).Name, ty.Field(i).Tag.Get("json"))
	}
}
func main() {
	PrintTag()
}
