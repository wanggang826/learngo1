package main

import (
	"fmt"
	"time"
)

func worker(workerId int,jobs<-chan int,result chan <- int){
	for job := range jobs{
		fmt.Printf("worker:%d start job:%d\n",workerId,job)
		result <- job * 2
		time.Sleep(time.Millisecond*500)
		fmt.Printf("worker:%d stop job:%d\n",workerId,job)
	}
}
func main()  {
	jobs := make(chan int,100)
	result := make(chan int,100)

	// 开启三个goroutine
	for j:=0;j<3;j++ {
		go worker(j,jobs,result)
	}

	// 发送5任务
	for i:=0;i<5;i++ {
		jobs <- i
	}
	close(jobs)

	for i :=0;i<5;i++{
		ret := <- result
		fmt.Println(ret)
	}
}
