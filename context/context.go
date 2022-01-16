package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，
//比如数据库和RPC服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、
//请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。
// 类似于c的信号量

var wg sync.WaitGroup
// 官方的基本做法，ctx可以被多个协程接收到
func work(ctx context.Context)  {
	go worker2(ctx)
LOOP:
	for  {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("worker stop")
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			fmt.Println("worker2 stop")
			break LOOP
		default:
		}
	}
}
// 基本能力
/*
Go内置两个函数：Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。
我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。
Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。
TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。
background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。
 */
func TestBasicContext() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go work(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
	time.Sleep(3 * time.Second)
	fmt.Println("end")
}
/*
worker
worker
worker
end
 */

/* context.Context是一个接口，该接口定义了四个需要实现的方法。具体签名如下：
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
其中：

Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
	如果当前Context被取消就会返回Canceled错误；
	如果当前Context超时就会返回DeadlineExceeded错误；
Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；
此外，context包中还定义了四个With系列函数。
 */

/*WithDeadline 返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，
则WithDeadline(parent, d)在语义上等同于父上下文。当截止日过期时，当调用返回的cancel函数时，
或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
 */