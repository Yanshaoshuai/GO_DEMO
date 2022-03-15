package main

import (
	"context"
	"sync"
	"time"
)

//struct
type Context interface {
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key interface{}) interface{}
}

//空 context
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

//context.Background()方法
var background = new(emptyCtx)

func Background() Context {
	return background
}

//cancelCtx
type canceler interface {
	cancel(bool, error)
} //只是为了去除报错 具体类型见源码
type cancelCtx struct {
	Context
	mu       sync.Mutex
	done     chan struct{}
	children map[canceler]struct{}
	err      error
}

//只需返回一个channel即可
func (c *cancelCtx) Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}
func (c *cancelCtx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

//作用是关闭自己及后代
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	c.mu.Lock()
	c.err = err                     //设置一个error，说明关闭原因
	close(c.done)                   //将channel关闭，以此通知派生的context
	for child := range c.children { //遍历所有children 逐个调用cancel方法
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()
	if removeFromParent { //正常情况下 需要将自己从parent中删除
		removeChild(c.Context, c)
	}
}
func WithCancel(parent Context) (ctx Context, cancel context.CancelFunc) {
	//c:=newCancelCtx(parent)
	//propagateCancel(parent,&c)//将自身添加到父节点
	//return  &c, func() {
	//	c.cancel(true,Canceled)
	//}
	//TODO implement me
	panic("implement me")
}
func removeChild(context Context, c *cancelCtx) {
	//just to avoid bug
}

type timeCtx struct {
	cancelCtx
	timer    *time.Timer
	deadline time.Time //指定最后期限
}

type valueCtx struct {
	Context
	key, val interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key) //可以通过context查询到父节点的value值
}
func main() {
}
