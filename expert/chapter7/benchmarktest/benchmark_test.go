package benchmarktest

import "testing"

//benchmark test
//文件名以_test.go结尾
//函数以BenchmarkXXX开始
//使用go test -bench=.开始性能测试(win下-bench=.改为-bench=".")
func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithPreAlloc()
	}
}
