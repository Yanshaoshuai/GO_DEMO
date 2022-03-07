package main

import "fmt"

// MapInitByLiteral 字面值初始化
func MapInitByLiteral() {
	m := map[string]int{
		"apple":  2,
		"banana": 3,
	}
	for k, v := range m {
		fmt.Printf("%s-%d\n", k, v)
	}
}

//MapInitByMake make 初始化
func MapInitByMake() {
	m := make(map[string]int, 10) //指定容量cap=10
	m["apple"] = 2
	m["banana"] = 3
	for k, v := range m {
		fmt.Printf("%s-%d\n", k, v)
	}
}
func MapCRUD() {
	m := make(map[string]string, 10)
	m["apple"] = "red"     //add
	m["apple"] = "green"   //update
	delete(m, "apple")     //delete
	v, exist := m["apple"] //查询 v value exist 键是否存在 键不存在返回值的零值
	if exist {
		fmt.Printf("apple-%s\n", v)
	}
	//删除键不存在或者nil map都不会报错 相当于空操作
	delete(m, "banana")
	var m1 map[string]int
	delete(m1, "banana")
}
func main() {
	MapInitByLiteral()
	MapInitByMake()
	MapCRUD()
	var m map[string]int
	fmt.Printf("len(m)=%d\n", len(m))
	m = make(map[string]int, 10)
	m["apple"] = 1
	fmt.Printf("len(m)=%d\n", len(m))
}
