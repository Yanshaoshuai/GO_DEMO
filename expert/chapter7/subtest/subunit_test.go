package subtest

import (
	"learn_go/expert/chapter7/unittest"
	"testing"
)

func sub1(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := unittest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d,%d)=%d;expected:%d", a, b, actual, expected)
	}
}

func sub2(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := unittest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d,%d)=%d;expected:%d", a, b, actual, expected)
	}
}

func sub3(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := unittest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d,%d)=%d;expected:%d", a, b, actual, expected)
	}
}

//筛选  go test subunit_test.go -v -run Sub/A= (包含匹配)
func TestSub(t *testing.T) {
	//自测试命名:<父测试名字>/<传递给Run的名字>
	//TestSub/A=1
	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
	t.Run("B=1", sub3)
}
