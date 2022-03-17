package exampletest

import "fmt"

func SayHello() {
	fmt.Println("Hello World")
}

func SayGoodbye() {
	fmt.Println("Hello,")
	fmt.Println("goodbye")
}

func PrintNames() {
	students := make(map[int]string, 4)
	students[0] = "Jim"
	students[1] = "Bob"
	students[2] = "Tom"
	students[3] = "Sue"
	for _, value := range students {
		fmt.Println(value)
	}
}
