package principle

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

// 简略版
func tRunner(t *T, fn func(t *T)) {
	t.runner = callerName(0)
	defer func() {
		t.duration += time.Since(t.start)
		if len(t.sub) > 0 { //存在子测试
			t.context.release()         //减少运行技术
			close(t.barrier)            //启动子测试 关闭管道 阻塞在管道上的协程会被唤醒
			for _, sub := range t.sub { //等待所有子测试结束
				<-sub.signal
			}
			if !t.isParallel { //非并发模式 等待并发执行
				t.context.waitParallel()
			}
		} else if t.isParallel { //并发模式 释放信号以启动新的测试
			t.context.release()
		}
		signal := true
		t.report() //测试结束后向父测试报告日志
		t.done = true
		t.signal <- signal //向调度者发送结束信号
	}()
	t.start = time.Now()
	fn(t)
	t.finished = true
}

// 简略版
func (t *T) Run(name string, f func(t *T)) bool {
	t = &T{
		//	common:common{
		//	barrier: make(chan bool),
		//	signal: make(chan bool),
		//	name: testName,//由name及父测试名组合而成 此处忽略
		//	parent: &t.common,
		//	level: t.level+1,//子测试层数+1
		//	chatty: t.chatty,
		//},context:t.context,
	}
	go tRunner(t, f) //启动协程执行子测试
	if !<-t.signal { //阻塞等待子测试结束
		//子测试要么执行结束 要么时子测试设置了Parallel
		//如果信号为false,说明出现异常退出
		runtime.Goexit()
	}
	return !t.failed
}

func (t *T) Parallel() {
	t.isParallel = true
	t.duration += time.Since(t.start)
	t.parent.sub = append(t.parent.sub, t) //将当前测试加入父测试子列表中 由父测试调度
	t.signal <- true                       //当前测试即将进入并发模式 标记测试结束 以便父测试不必等待并退出Run()
	<-t.parent.barrier                     //等待父测试发送子测试启动信号
	t.context.waitParallel()               //阻塞等待并发调度
	t.start = time.Now()
}

type T struct {
	common
	isParallel bool         //是否并发
	context    *testContext //控制测试的并发调度
}

func (t T) report() {

}

type matcher struct {
}

type testContext struct {
	match         *matcher   //匹配器 用于管理测试名称的匹配 过滤等
	mu            sync.Mutex //用于控制textContext成员的互斥访问
	startParallel chan bool  //用于通知测试可以并发执行的控制管道,测试并发达到最大限制时,需要阻塞等待该管道的通知事件
	running       int        //当前并发执行的测试个数
	numWaiting    int        //等待并发执行的测试个数,所有等待执行的测试都阻塞在startParallel管道处
	maxParallel   int        //最大并发数,默认为系统CPU数 可以通过-parallel 指定
}

func (c *testContext) waitParallel() {
	c.mu.Lock()
	if c.running < c.maxParallel { //当前运行的测试数未达到最大值
		c.running++
		c.mu.Unlock()
		return
	}
	c.numWaiting++
	c.mu.Unlock()
	<-c.startParallel //阻塞等待
}

func (c *testContext) release() {
	c.mu.Lock()
	if c.numWaiting == 0 { //没有等待执行的测试
		c.running--
		c.mu.Unlock()
		return
	}
	c.numWaiting--
	c.mu.Unlock()
	c.startParallel <- true //通知等待执行的测试执行
}

type M struct {
	tests      []testing.InternalTest      //单元测试
	benchmarks []testing.InternalBenchmark //性能测试
	examples   []testing.InternalExample   //示例测试
	timer      *time.Timer                 //测试超时时间
}
