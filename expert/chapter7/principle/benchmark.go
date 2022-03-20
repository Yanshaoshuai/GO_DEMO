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

func (b *B) runN(n int) {
	b.N = n        //指定n值
	b.ResetTimer() //重置时间
	b.StartTimer()
	b.benchFunc(b) //执行测试
	b.StopTimer()
}

func (b *B) Run(name string, f func(b *B)) bool {
	sub := &B{
		common: common{signal: make(chan bool),
			name:   name,
			parent: &b.common,
		},
		benchFunc: f,
	}
	if sub.run1() { //先执行一次 如果子测试不出错且子测试没有子测试则继续执行sub.run()
		sub.run() //run里决定要执行多少次runN()
	}
	b.add(sub.result) //累加统计结果到父测试中
	return !sub.failed
}

func (b *B) run1() bool {
	//TODO implement me
	panic("implement me")
}

func (b *B) run() {

}

func (b *B) add(result testing.BenchmarkResult) {

}

func (b *B) launch() {
	d := b.benchTime
	for n := 1; !b.failed && b.duration < d && n < 1e9; { //最少执行b.benchTime(默认1s)时间
		//最多执行1e9次
		//last:=n
		//n=int(d.Nanoseconds())//预测接下来要执行多少次 b.benchTime/每个操作耗时
		//if nsop:=b.nsPerOP();nsop!=0{
		//	n/=int(nsop)
		//}
		//n=max(min(m+n/5,100*last),last+1)//避免增长较快 先增长20%,至少增长1次
		//n=roundUp(n)//下次迭代次数向上取整到10的指数 方便阅读
		b.runN(n)
	}
}
