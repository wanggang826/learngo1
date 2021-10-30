package main

import (
	"fmt"
	"time"
)


// ===========模拟任务Task===========
type Task struct {
	f func() error
}

// NewTask 创建Task
func NewTask(arg_f func() error) *Task {
	task := Task{
		f: arg_f,
	}
	return &task
}

// Tash 业务方法
func (t *Task) Do() {
	t.f()
}

// =================协程池Poll=================

// 定义Pool
type Pool struct {
	EntryChannel chan *Task
	JobChannel   chan *Task
	WorkerNum    int
}

//创建Pool
func NewPool(num int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobChannel:   make(chan *Task),
		WorkerNum:    num,
	}
	return &p
}

//Pool 绑定worker
func (p *Pool) worker(workerId int) {
	for task := range p.JobChannel {
		task.Do()
		fmt.Println("workerId", workerId, "完成一个任务")
	}
}

//Pool 运行
func (p *Pool) Run() {
	for i := 0; i < p.WorkerNum; i++ {
		go p.worker(i)
	}

	for task := range p.EntryChannel {
		p.JobChannel <- task
	}
}

func main() {
	// 测一下协程池
	t := NewTask(func() error {
		fmt.Println("打印一下当前时间", time.Now())
		return nil
	})

	p := NewPool(4)

	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	p.Run()
}
