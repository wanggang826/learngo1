package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch = make(chan int, 20)
	var chPro = make(chan bool, 1)

	//并发处理管道内的数据
	var wg = sync.WaitGroup{}
	delta := 3
	wg.Add(delta)
	for i := 0; i < delta; i++ {
		go consumer(ch)
	}

	go producer(ch, chPro)

	//监控生产消费是否完毕
	bol := false
	for {
		a, ok := <-chPro
		if ok {
			bol = a
		}
		fmt.Println("+++++chPro+++++", bol, len(ch))
		if bol && len(ch) == 0 {
			fmt.Println("+++++Done start+++++")
			for i := 0; i < delta; i++ {
				wg.Done()
			}
			fmt.Println("+++++Done end+++++")
		} else {
			fmt.Println("+++++Wait+++++")
			wg.Wait()
		}
	}
}

func producer(ch chan int, chPro chan bool) {
	defer close(ch)
	//defer close(chPro)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	//chPro <- true
	fmt.Println("+++++close+++++")
	return
}

func consumer(ch chan int) {
	m := 0
	for n := range ch {
		m += n
	}
	fmt.Println("+++++m+++++", m)
}