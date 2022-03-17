package subtest

import (
	"testing"
	"time"
)

func parallelTest1(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
	//do sth
}

func parallelTest2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	//do sth
}

func parallelTest3(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
	//do sth
}

// go test .\subparallel_test.go -v -run SubParallel
func TestSubParallel(t *testing.T) {
	t.Logf("Setup")

	t.Run("group", func(t *testing.T) {
		//多个自测试并发执行
		t.Run("Test1", parallelTest1)
		t.Run("Test2", parallelTest2)
		t.Run("Test3", parallelTest3)
	})

	//在所有子测试结束后才执行
	t.Logf("teardown")
}
