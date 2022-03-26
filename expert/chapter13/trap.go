package main

import (
	"errors"
	"fmt"
	"testing"
)

//errs 初始容量为0,每次添加新元素后会扩容生成新切片
//不接收append返回值则不能获取最新的切片

func Validation() []error {
	var errs []error
	_ = append(errs, errors.New("error 1"))
	_ = append(errs, errors.New("error 2"))
	_ = append(errs, errors.New("error 3"))
	return errs
}

func ValidateName(name string) error {
	if name != "" {
		return nil
	}
	return errors.New("empty name")
}

func Validations(name string) []error { //如果后续根据errs的长度判断是否有错误 会出现逻辑错误
	var errs []error
	errs = append(errs, ValidateName(name)) //向切片中添加nil元素不会报错而且添加后slice的长度会为1 cap也会为1
	return errs
}

//append 每次返回一个新切片 即使底层数组不变 不会修改原切片
//append 从原切片的len位置插入到底层数组中

func AppendDemo() {
	x := make([]int, 0, 10) //底层数组 arr 0 0 0 0 0 0 0 0 0 0
	//x len=0 cap=10 data=*arr
	x = append(x, 1, 2, 3) //arr 1 2 3 0 0 0 0 0 0 0
	//x len=3 cap=10
	y := append(x, 4) // arr 1 2 3 4 0 0 0 0 0 0
	//x len=3 cap=10
	//y len=4 cap=10 data=*arr
	z := append(x, 5) // arr 1 2 3 5 0 0 0 0 0
	// x len=3 cap=10
	//y len=4 cap=10
	//z len=4 cap=10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func foo() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
}
func suggestFoo() {
	var out []*int
	for i := 0; i < 3; i++ {
		iCopy := i
		out = append(out, &iCopy)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
}

func Process1(tasks []string) {
	for _, task := range tasks {
		go func() { //task不确定 执行时才确定
			fmt.Printf("Worker start process task:%s\n", task)
		}()
	}
}
func Process2(tasks []string) {
	for _, task := range tasks {
		go func(t string) {
			fmt.Printf("Worker start process task:%s\n", t)
		}(task) //绑定了task task确定
	}
}
func Double(a int) int {
	return a * 2
}
func TestDouble(t *testing.T) {
	var tests = []struct {
		name         string
		input        int
		expectOutput int
	}{
		{
			name:         "double 1 should got 2",
			input:        1,
			expectOutput: 2,
		},
		{
			name:         "double 2 should got 4",
			input:        2,
			expectOutput: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) { //t.Run并不会启动新协程 不会有问题
			if test.expectOutput != Double(test.input) {
				t.Fatalf("except: %d,but got: %d", test.input, test.expectOutput)
			}
		})
	}
}
func main() {
	errs := Validations("1")
	println(cap(errs))
	AppendDemo()
	foo()
	suggestFoo()
}
