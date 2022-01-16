package concurrency

import (
	"fmt"
	"math/rand"
)

type Job struct {
	id      int
	randNum int
}

type Result struct {
	job *Job
	sum int
}
// 可以有效控制goroutine数量，防止暴涨
func createPool(num int, jobChan chan *Job, result chan *Result) {
	for i := 1; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				r_num := job.randNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}

		}(jobChan, result)
	}
}

func TestRoutinePool() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(10, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		for channel := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", channel.job.id,
				channel.job.randNum, channel.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			id:      id,
			randNum: r_num,
		}
		jobChan <- job
	}
}
/*
job id:104027 randnum:2881529496136327593 result:93
job id:104028 randnum:6024352312151558634 result:66
job id:104029 randnum:4138689334198994689 result:112
job id:104030 randnum:6515244689386219695 result:99
job id:104031 randnum:7870596785565970853 result:110
job id:104032 randnum:5231450307762985951 result:82
job id:104033 randnum:8437812401890446516 result:81
job id:104034 randnum:8182538402223171686 result:77

 */