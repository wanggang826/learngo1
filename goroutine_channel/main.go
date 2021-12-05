package main

import (
	"fmt"
	"sync"
)
/*

*/

/*
两个goroutine
1 生成0-100的数字发送到ch1
2 从ch1取值 计算平方把结果发送到ch2中
*/
var wg sync.WaitGroup

func p(ch1 chan int){
	for i:=0;i<100;i++{
		ch1 <- i
	}
	close(ch1)
	wg.Done()
}

func c(ch1 chan int,ch2 chan int){
	for{
		tmp,ok := <- ch1
		if !ok{
			break
		}
		ch2 <- tmp*tmp
	}
	close(ch2)
	wg.Done()
}

// channel 限制最大并发数 利用channel 阻塞特性 加上channel缓冲
func channel(){
	count := 10
	sum := 100
	c := make(chan struct{},count)
	sc := make(chan struct{},sum)
	defer close(c)
	defer close(sc)
	for i:= 0;i<sum;i++{
		c <- struct{}{}
		go func(j int) {
			fmt.Println(j)
			<- c
			sc <- struct{}{}
		}(i)
	}
	for i:= sum;i>0;i-- {
		<-sc
	}

}
func main()  {
	ch1 := make(chan int,100)
	ch2 := make(chan int,20)
	wg.Add(2)
	go p(ch1)
	go c(ch1,ch2)
	for ret := range ch2{
		fmt.Println(ret)
	}
	wg.Wait()
}