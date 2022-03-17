package maintest

import (
	"learn_go/expert/chapter7/benchmarktest"
	"learn_go/expert/chapter7/exampletest"
	"learn_go/expert/chapter7/unittest"
	"os"
	"testing"
)

// go test .\main_test.go -v -bench="."
func TestMain(m *testing.M) {
	println("TestMain setup")
	retCode := m.Run() //执行测试 包括 单元测试 性能测试 和 示例测试
	println("TestMain tear-down.")
	os.Exit(retCode)
}
func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarktest.MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarktest.MakeSliceWithPreAlloc()
	}
}
func ExampleSayHello() {
	exampletest.SayHello()
	// Output: Hello World
}
func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := unittest.Add(a, b)

	if actual != expected {
		//标记测试失败
		t.Errorf("Add(%d,%d) = %d; expected: %d", a, b, actual, expected)
	}
}
