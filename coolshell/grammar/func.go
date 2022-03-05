package main

import "fmt"

func max(a int, b int) int { //注意参数和返回值是怎么声明的

	if a > b {
		return a
	}
	return b
}

//多个返回值
func multi_ret(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	var err bool
	var val int
	val, err = m[key]
	return val, err
}

//函数不定参数
func sum(nums ...int) {
	fmt.Print(nums, " ") //输出如 [1, 2, 3] 之类的数组
	total := 0
	for _, num := range nums { //要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}

//返回闭包
func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}
func main() {
	fmt.Println(max(4, 5))

	//Go中很多Package 都会返回两个值，一个是正常值，一个是错误
	v, e := multi_ret("one")
	fmt.Println(v, e) //输出 1 true
	v, e = multi_ret("four")
	fmt.Println(v, e) //输出 0 false
	//通常的用法(注意分号后有e)
	if v, e = multi_ret("four"); e {
		// 正常返回
	} else {
		// 出错返回
	}

	//不定参数
	sum(1, 2)
	sum(1, 2, 3)
	//传数组
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//函数闭包
	nextNumFunc := nextNum()
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}
}
