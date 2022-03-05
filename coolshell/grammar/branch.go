package main

import "fmt"

//Go的正式的语法使用分号来终止语句。和C不同的是，这些分号由词法分析器在扫描源代码过程中使用简单的规则自动插入分号，因此输入源代码多数时候就不需要分号了。
//Go程序仅在for循环语句中使用分号，以此来分开初始化器、条件和增量单元。如果你在一行中写多个语句，也需要用分号分开
//无论任何时候，你都不应该将一个控制结构（(if、for、switch或select）的左大括号放在下一行。如果这样做，将会在大括号的前方插入一个分号，这可能导致出现不想要的结果。
func main() {
	//if 语句
	x := 100
	var i = 100
	num := 100

	if x%2 == 0 {
		//...
	}
	//if - else
	if x%2 == 0 {
		//偶数...
	} else {
		//奇数...
	}
	//多分支
	if num < 0 {
		//负数
	} else if num == 0 {
		//零
	} else {
		//正数
	}
	//switch 语句 没有break，还可以使用逗号case多个值
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6:
		fmt.Println("four, five, six")
	default:
		fmt.Println("invalid value!")
	}
	//经典的for语句 init; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//精简的for语句 condition
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//死循环的for语句 相当于for(;;)
	for {
		if i > 10 {
			break
		}
		i++
	}

}
