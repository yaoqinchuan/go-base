package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var x int64

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
}

//测试并发问题
func TestConcurrencyProblem() {
	go add()
	go add()
	go add()
	<-time.After(5 * time.Second)
	fmt.Println("x value is", x)
}

/*
x value is 10214
*/

// 互斥锁 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁
// WaitGroup类似于countdownLatch
var wg sync.WaitGroup
var lock sync.Mutex

func addMutex() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

//测试并发问题加锁
func TestConcurrencyProblemResolved() {
	wg.Add(3)
	go addMutex()
	go addMutex()
	go addMutex()
	wg.Wait()
	fmt.Println("x value is", x)
}

/*
x value is 15000
*/

// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，
//这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。
//读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
//当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
var (
	rwLock sync.RWMutex
)

func write() {
	rwLock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwLock.Unlock()
	wg.Done()
}
func read() {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	wg.Done()
}
func TestReadWriteLock() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。
//Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
//sync.Once只有一个Do方法，其签名如下
func testSyncOnce() {}
