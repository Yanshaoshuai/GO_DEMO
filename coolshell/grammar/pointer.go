package main

import "fmt"

func main() {
	var i int = 1
	var pInt *int = &i
	//输出：i=1     pInt=0xf8400371b0       *pInt=1
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

	*pInt = 2
	//输出：i=2     pInt=0xf8400371b0       *pInt=2
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

	i = 3
	//输出：i=3     pInt=0xf8400371b0       *pInt=3
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)
}
