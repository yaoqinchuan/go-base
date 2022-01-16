package concurrency

import (
	"fmt"
	"time"
)

func fillChannel(channel chan<- int, value int) {
	<-time.After(time.Duration(value) * time.Second)
	channel <- value
}

// select在两个条件同时到达的时候会随机选择一个，没有default的情况下会等待某一个case完成
func TestSelect() {
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)
	go fillChannel(chan1, 1)
	go fillChannel(chan2, 2)
	for i := 0; i < 2; i++ {
		select {
		case value := <-chan1:
			fmt.Println("channel 1 receive message: ", value)
		case value := <-chan2:
			fmt.Println("channel 2 receive message: ", value)
		}
	}
}

func produceDataToData(channel chan<- int) {
	i := 0
	for {
		i++
		channel <- i
		fmt.Println("produce send message ", i)
		<-time.After(time.Second)
	}
}
// default用来检测通道是否满了
func TestSelectToJudgeChannelFilled() {
	chan1 := make(chan int, 10)
	go produceDataToData(chan1)
	i := 0
	for {
		<-time.After(2 * time.Second)
		select {
		case chan1 <- i:
			fmt.Println("main send message ", i)
		default:
			fmt.Println("channel filled, can not write in.")
		}
		i++
	}
}
/*
produce send message  5
produce send message  6
main send message  2
produce send message  7
channel filled, can not write in.
channel filled, can not write in.
channel filled, can not write in.

 */