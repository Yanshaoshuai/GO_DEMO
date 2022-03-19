package principle

import (
	"runtime"
	"testing"
	"time"
)

type benchContext struct {
}

type B struct {
	common
	importPath       string // import path of the package containing the benchmark
	context          *benchContext
	N                int           //目标代码执行次数 不需要用户了解具体值 会自动调整
	previousN        int           // number of iterations in the previous run
	previousDuration time.Duration // total duration of the previous run
	benchFunc        func(b *B)    //性能测试函数
	benchTime        time.Duration //性能测试最少执行时间 默认为1秒 -benchtime指定
	bytes            int64         //每次迭代处理的字节数
	missingBytes     bool          // one of the subbenchmarks does not have bytes set.
	timerOn          bool          //是否已开始计时
	showAllocResult  bool
	result           testing.BenchmarkResult //测试结果
	parallelism      int                     // RunParallel creates parallelism*GOMAXPROCS goroutines
	// The initial states of memStats.Mallocs and memStats.TotalAlloc.
	startAllocs uint64 //计时开始时堆中分配的对象总数
	startBytes  uint64 //计时开始时堆中分配的字节总数
	// The net total of this test after being run.
	netAllocs uint64 //计时结束时堆中增加的对象总数
	netBytes  uint64 //计时结束时堆中增加的字节总数
	// Extra metrics collected by ReportMetric.
	extra map[string]float64
}

var memStatus runtime.MemStats

//StartTimer 启动计时
func (b *B) StartTimer() {
	if !b.timerOn {
		runtime.ReadMemStats(&memStatus)    //读取当前堆内存的分配信息
		b.startAllocs = memStatus.Mallocs   //记录当前堆内存分配的对象数
		b.startBytes = memStatus.TotalAlloc //记录当前堆内存分配的字节数
		b.start = time.Now()
		b.timerOn = true
	}
}

//StopTimer 停止计时
// 并不一定是测试结束 一个测试中可能有多个统计阶段 其统计值是累加的
func (b *B) StopTimer() {
	if b.timerOn {
		b.duration += time.Since(b.start) //累加测试耗时
		runtime.ReadMemStats(&memStatus)
		b.netAllocs += memStatus.Mallocs - b.startAllocs  //累加堆内存分配对象数
		b.netBytes += memStatus.TotalAlloc - b.startBytes //累加堆内存分配字节数
		b.timerOn = false
	}
}

//ResetTimer 重置计时
func (b *B) ResetTimer() {
	if b.timerOn {
		runtime.ReadMemStats(&memStatus)
		b.startAllocs = memStatus.Mallocs
		b.startBytes = memStatus.TotalAlloc
		b.start = time.Now()
	}
	b.duration = 0
	b.netAllocs = 0
	b.netBytes = 0
}

//SetByte 由用户设置 设置之后输出会打印 xxx MB/s的信息
func (b *B) SetByte(n int64) {
	b.bytes = n
}

//ReportAllocs 用于设置是否打印内存统计信息
// 等同于 -benchmem
func (b *B) ReportAllocs() {
	b.showAllocResult = true
}
