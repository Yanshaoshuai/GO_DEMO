package main

import "fmt"

func main() {
	m := make(map[string]int) //使用make创建一个空的map

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)      //输出 map[three:3 two:2 one:1] (顺序在运行时可能不一样)
	fmt.Println(len(m)) //输出 3

	v := m["two"]  //从map里取值
	fmt.Println(v) // 输出 2

	delete(m, "two")
	fmt.Println(m) //输出 map[three:3 one:1]

	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1) //输出 map[two:2 three:3 one:1] (顺序在运行时可能不一样)

	for key, val := range m1 {
		fmt.Printf("%s => %d \n", key, val)
		/*输出：(顺序在运行时可能不一样)
		  three => 3
		  one => 1
		  two => 2*/
	}
}
