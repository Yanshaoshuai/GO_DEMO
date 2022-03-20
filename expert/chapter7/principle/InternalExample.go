package principle

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type InternalExample struct {
	Name      string //测试名称
	F         func() //测试函数
	Output    string //期望字符串
	Unordered bool   //输出是否是无序的

}

var chatty *bool //-v

func runExample(eg InternalExample) (ok bool) {
	if *chatty {
		fmt.Printf("===RUN   %s\n", eg.Name)
	}
	//Capture stdout
	stdout := os.Stdout    //备份标准输出
	r, w, err := os.Pipe() //创建一个管道
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Stdout = w //把管道入口赋值给标准输出 即所有输出都会进入管道
	outC := make(chan string)
	go func() {
		var buf strings.Builder
		_, err := io.Copy(&buf, r) //从管道中读出数据
		r.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "testing: copying pipe: %v\n", err)
			os.Exit(1)
		}
		outC <- buf.String() //将管道中读出的数据写入channel
	}()
	start := time.Now()
	ok = true

	defer func() {
		dstr := fmtDuration(time.Since(start))
		w.Close()          //关闭管道
		os.Stdout = stdout //恢复原标准输出
		out := <-outC      //从channel中取出数据
		var fail string
		err := recover()
		got := strings.TrimSpace(out)        //实际得到的打印字符串
		want := strings.TrimSpace(eg.Output) //期望字符串
		if eg.Unordered {
			if sortLines(got) != sortLines(want) && err == nil {
				fail = fmt.Sprintf("got:\n%s\nwant (unordered):\n%s\n", out, eg.Output)
			}
		} else {
			if got != want && err == nil {
				fail = fmt.Sprintf("got:\n%s\nwant:\n%s\n", got, want)
			}
		}
		if fail != "" || err != nil {
			fmt.Printf("--- FAIL:%s (%s)\n%s", eg.Name, dstr, fail)
			ok = false
		} else if *chatty {
			fmt.Printf("=== PASS:%s (%s)\n", eg.Name, dstr)
		}
		if err != nil {
			panic(err)
		}
	}()
	eg.F() //Run example
	return
}

func sortLines(want string) string {
	//TODO implement me
	panic("implement me")
}

func fmtDuration(since time.Duration) string {
	//TODO implement me
	panic("implement me")
}
