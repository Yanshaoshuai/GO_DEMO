package main

import (
	"bytes"
	"fmt"
)

func testScale01() {
	foo := make([]int, 5)       //cap=len=5
	fmt.Printf("foo:%v\n", foo) //[0 0 0 0 0]
	foo[3] = 42                 //[0 0 0 42 0]
	foo[4] = 100                //[0 0 0 42 100]
	fmt.Printf("foo:%v\n", foo)
	bar := foo[1:4] //bar:[0 0 42]
	bar[1] = 99     //bar:[0 99 42] foo:[0 0 99 42 100]
	fmt.Printf("foo:%v\n", foo)
	fmt.Printf("bar:%v\n", bar)
	bar = append(bar, 1000) //bar:[0 99 42,1000] foo:[0 0 99 42 1000]
	fmt.Printf("foo:%v\n", foo)
	fmt.Printf("bar:%v\n", bar)
	bar = append(bar, 1000) //bar:[0 99 42,1000,1000] foo:[0 0 99 42 1000]
	fmt.Printf("foo:%v\n", foo)
	fmt.Printf("bar:%v\n", bar)
	println()
}
func testScale02() {
	foo := make([]int, 5) //cap=len=5
	bar := foo[1:4]       //bar:[0,0,0] foo:[0,0,0,0,0]
	fmt.Printf("foo:%v\n", foo)
	fmt.Printf("bar:%v\n", bar)
	foo = append(foo, 1) //bar:[0,0,0] foo:[0,0,0,0,0,1]
	fmt.Printf("foo:%v\n", foo)
	fmt.Printf("bar:%v\n", bar)
}
func testScale03() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}
func testScale04() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex:sepIndex] //dir1 cap=sepIndex
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...) //scale

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
}
func main() {
	//append()这个函数在 cap 不够用的时候就会重新分配内存以扩大容量，而如果够用的时候不不会重新分享内存
	println("testScale01")
	testScale01()
	println("testScale02")
	testScale02()
	println("testScale03")
	testScale03()
	println("testScale04")
	testScale04()
}
