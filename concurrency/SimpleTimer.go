package concurrency

import (
	"fmt"
	"time"
)

// 测试一下定时器基本能力
func TestTimer() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C // 会阻塞在这2秒
	fmt.Println("2 秒到了")

	// 定时器会阻塞，但是time本身的方法不会
	timer1 := time.NewTimer(2 * time.Second)
	t1 := <-timer1.C
	fmt.Println("2秒又到了，现在时间 %v", t1)
	t2 := time.Now()
	fmt.Println("现在时间 %v", t2)

	// 只有使用了<-才能够卡主
	<-time.After(2 * time.Second)
	fmt.Println("2 秒又到了")

	// 停止定时器
	time4 := time.NewTimer(2 * time.Second)
	go func() {
		<-time4.C
		fmt.Println("定时器4执行完成")
	}()
	time4.Stop()
	fmt.Println("关闭定时器4")

	// 定时器只能跑一次,可以重置
	time5 := time.NewTimer(time.Second)
	for i := 1; i < 3; i++ {
		<-time5.C
		fmt.Println("定时器5计时到达", i)
		time5.Reset(time.Second)
	}

}

/*
2 秒到了
2秒又到了，现在时间 %v 2022-01-11 00:02:40.9030364 +0800 CST m=+4.017546301
现在时间 %v 2022-01-11 00:02:40.9509353 +0800 CST m=+4.065445201
2 秒又到了
关闭定时器4


定时器5计时到达 1
定时器5计时到达 2
*/

// 定时器是可以重复触发的
func TestTicker() {
	ticker := time.NewTicker(time.Second)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for  {}
}
/*
2022-01-11 00:11:30.3775867 +0800 CST m=+1.015215801
2022-01-11 00:11:31.3786253 +0800 CST m=+2.016254401
2022-01-11 00:11:32.3786951 +0800 CST m=+3.016324201
2022-01-11 00:11:33.3776615 +0800 CST m=+4.015290601
2022-01-11 00:11:34.3788892 +0800 CST m=+5.016518301
 */
