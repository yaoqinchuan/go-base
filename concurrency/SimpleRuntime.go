package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Gosched()  让出CPU时间片，重新等待安排任务
func TestGosched() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
/*
hello
world
hello
 */

// runtime.Goexit()  退出当前协程
func TestGoexit()  {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {}
}
/*
B.defer
A.defer
 */

// runtime.GOMAXPROCS
/*
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。
Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
两个任务只有一个逻辑核心，此时是一个核心来回切。 将逻辑核心数设为2，此时两个任务并行执行，代码如下。
 */
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
		time.Sleep(time.Second)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
		time.Sleep(time.Second)
	}
}
func TestGoMaxProcs() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(60 * time.Second)
}