//文件名：hello.go
package main //声明本文件的package名

//1）在import中，你可以使用相对路径，如 ./或 ../ 来引用你的package
//
//2）如果没有使用相对路径，那么，go会去找$GOPATH/src/目录。
import "fmt" //import语言的fmt库——用于输出

func main() {
	fmt.Println("hello world")
}

//你可以有两种运行方式
//
//#解释执行（实际是编译成a.out再执行）
//$go run hello.go
//hello world
//#编译执行
//$go build hello.go
//$ls
//hello hello.go
//$./hello
//hello world
