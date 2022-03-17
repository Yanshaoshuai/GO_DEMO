package exampletest

//Example测试
//文件名以_test结尾
//函数名以Example开头
//单行输出 // Output: <期望字符串>
//多行输出 // Output: \n <期望字符串> \n <期望字符串>...
//无序输出 // Unordered output: \n <期望字符串> \n <期望字符串>...
//测试字符串时会自动忽略字符串前后的空白字符
// 没有Output或者Unordered output标识 则该函数不会被执行
func ExampleSayHello() {
	SayHello()
	// Output: Hello World
}

func ExampleSayGoodbye() {
	SayGoodbye()
	// Output:
	// Hello,
	// goodbye
}

func ExamplePrintNames() {
	PrintNames()
	// Unordered output:
	// Jim
	// Bob
	// Tom
	// Sue
}
