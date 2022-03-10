package main

import (
	"fmt"
	"time"
)

//k8s apiserver 使用select实现永久阻塞
//func block() {
//	server:=webhooktesting.NewTestServer(nil)
//	server.StartTLS
//	fmt.Println("serving on",server.URL)
//	select {}
//}

//k8s调度器使用select快速检错
//func fastCheckError() {
//	errCh:=make(chan  error,active)
//	jm.deleteJobPods(&job,activePods,errCh)//传入管道用于记录错误
//	select {
//	case manageJobErr=<-errCh://检查是否有错
//		if manageJobErr!=nil{
//			break
//		}
//	default:
//		//没有错误快速结束
//	}
//}

//k8s控制器使用select实现限时等待
func waitForStopOrTimeout(stopCh <-chan struct{}, timout time.Duration) <-chan struct{} {
	stopChWithTimeout := make(chan struct{})
	go func() {
		select {
		case <-stopCh: //自然结束
		case <-time.After(timout): //最长等待timout退出
		}
		stopChWithTimeout <- struct{}{}
		close(stopChWithTimeout)
	}()
	return stopChWithTimeout
}
func main() {
	stopCh := make(chan struct{})
	finishCh := waitForStopOrTimeout(stopCh, time.Second*3)
	if s, ok := <-finishCh; ok {
		fmt.Printf("finish %v\n", s)
	}
	time.Sleep(time.Second * 10)
}
