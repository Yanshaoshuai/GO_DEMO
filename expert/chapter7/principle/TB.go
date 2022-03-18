package principle

// TB common实现了此接口
type TB interface {
	Cleanup(func())
	//标记失败+记录日志
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	//标记失败
	Fail()
	//标记失败并退出
	FailNow()

	Failed() bool
	//标记失败+记录日志+结束测试
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Helper()
	//记录日志
	Log(args ...interface{})
	Logf(format string, args ...interface{})

	Name() string
	Setenv(key, value string)
	//跳过测试并退出
	SkipNow()
	//跳过测试+记录日志并退出
	Skip(args ...interface{})
	Skipf(format string, args ...interface{})

	Skipped() bool
	TempDir() string

	// A private method to prevent users implementing the
	// interface and so future additions to it will not
	// violate Go 1 compatibility.
	private() //保证了即使用户实现了类似的接口也不会跟testing.TB接口冲突
	//因为用户不能够实现此方法
}
