package main

import (
	"flag"
	"fmt"
	"strings"
)

// init 方法 先于main执行
func init() {
	//获取启动参数
	//构造一个参数并返回值的string指针 name 参数名 value 默认值 usage发生错误时回调信息
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n pro:正式环境\n")
	mode := flag.String("mode", "local", "请输入启动模式.\nlocal 本地 \nserver 服务器\n")
	//解析
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		println("active dev")
	case "fat":
		println("active fat")
	case "uat":
		println("active uat")
	case "pro":
		println("active pro")
	default:
		println("active fat")
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'fat' will be used.")
	}
	switch *mode {
	case "local":
		println("local mode")
	case "server":
		println("server mode")
	}
}

// init()方法可以有多个
func init() {

}
func main() {

}
