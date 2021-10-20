package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int, 200)

	//并发处理管道内的数据

	delta := 3
	wg.Add(delta)
	for i := 0; i < delta; i++ {
		go consumer(ch)
	}
	wg.Add(1)
	go producer(ch)

	wg.Wait()
	fmt.Println("test==========")


}

func producer(ch chan int) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	fmt.Println("+++++close+++++")
}

func consumer(ch chan int) {
	defer wg.Done()
	m := 0
	for n := range ch {
		m += n
	}
	fmt.Println("+++++m+++++", m)
}