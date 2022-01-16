package concurrency

import (
	"fmt"
	"time"
)

// 测试一个简单地通道发送与接收,无人消费的无缓冲通道写入数据会报错
func TestSendAndReceive() {
	ch := make(chan int, 1)
	ch <- 10
	result, _ := <-ch
	close(ch)
	fmt.Println(result)
}

// 10

// channel 关闭就不能写了但是可以读
func TestTwoChannel() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			chan1 <- i
		}
		close(chan1)
	}()
	go func() {
		for {
			i, ok := <-chan1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			chan2 <- i * i
		}
		close(chan2)
	}()

	// 我们通常使用的是for range的方式
	for i := range chan2 {
		fmt.Println(i)
	}
}

// 单向通道 可以才参数列表里面约束
// out是一个只能写入的int的通道
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

// out是一个只能写入的int的通道 in是一个只能读出的通道
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// in是一个只能读出的通道
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func TestSingleDirectionChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

// 测试通道数据是否可以被重复读取
func TestConsumerChannelRepeat() {
	ch1 := make(chan int, 128)
	for i := 0; i < 128; i++ {
		ch1 <- i
	}
	for data := range ch1 {
		fmt.Println("consumer once ", data)
	}
	for data := range ch1 {
		fmt.Println("consumer twice ", data)
	}
}
// 通过下面结果，说明不行
/*
...
consumer once  127

main.main()
        G:/go/basic/main/main.go:9 +0x27
*/


func selectChannel(channel <-chan int) {
	select {
	case <-channel:
		fmt.Println("channel data received")
	}
}

// 测试通道数据是否可以被重复select
func TestSelectChannelRepeat() {
	ch1 := make(chan int, 1)
	go selectChannel(ch1)
	go selectChannel(ch1)
	ch1 <- 1
	<-time.After(3 * time.Second)
}
// 通过下面结果，说明不行
/*
channel data received

Process finished with exit code 0

*/
