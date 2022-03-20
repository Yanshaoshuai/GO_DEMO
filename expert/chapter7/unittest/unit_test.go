package unittest

import (
	"flag"
	"fmt"
	"testing"
)

//单元测试方法
//文件名以xxx_test结尾
//TestXXX开头
//有且仅有*testing.T参数
func TestAdd(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	//-args 指定参数
	for i, arg := range flag.Args() {
		fmt.Printf("arg %d is %s\n", i, arg)
	}
	var a = 1
	var b = 2
	var expected = 3
	actual := Add(a, b)

	if actual != expected {
		//标记测试失败
		t.Errorf("Add(%d,%d) = %d; expected: %d", a, b, actual, expected)
	}
}
