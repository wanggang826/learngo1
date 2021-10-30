package main

import (
	"fmt"
	"sync"
	"time"
)
//读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，
//如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待
// 读的场景远远大于写的场景   用读写锁
var(
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwLock sync.RWMutex
)

func read(){
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
	wg.Done()
}

func write(){
	//lock.Lock()
	rwLock.Lock()
	x = x+1
	time.Sleep(time.Millisecond*10)
	rwLock.Unlock()
	//lock.Unlock()
	wg.Done()
}
func main(){
	start := time.Now()
	for i:=0;i<1000;i++{
		wg.Add(1)
		go read()
	}

	for i:=0;i<10;i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
