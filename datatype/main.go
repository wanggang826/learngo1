package main

import (
	"errors"
	"fmt"
	"sync"
)

// 值类型 ：变量直接存储值，内容通常在栈中分配( 数值类型，bool,string, 数组 结构体)

// 引用类型： 变量存储的是一个地址，这个地址存储最终的值，内容通常在堆上分配，通过GC回收(指针，slice,map,chan,interface

// GO参数传递：指针还是值
// 对不需要函数进行修改的、小的结构体中使用值传递（值传递可以避免由于别名而带来的值的修改错误，有时使用小的结构体的值进行参数传递可以避免缓存未命中或者重新分配堆空间）
// 对于大的结构体，或者不得不修改的结构体传指针,否则传值，因为由于传递指针导致结构体被修改的问题很难被排查

// interface
var wg sync.WaitGroup
func main() {
	//res := funcName(22)
	//
	//wg.Add(1)
	//go func() {
	//	fmt.Println("eeeeee")
	//	wg.Done()
	//}()
	//wg.Wait()
	//fmt.Println("res",res)
	test()
}

func funcName(a interface{}) string {

	defer func() {
		// 捕获 test 抛出的panic
		if err := recover(); err != nil {
			fmt.Println("funcName 发生错误", err)
		}
	}()
	value := a.(string)

	//if !ok {
	//	fmt.Println("It's not ok for type string")
	//	return "11"
	//}
	//fmt.Println("The value is ", value)
	return value
}

func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			//debug.PrintStack()
		}
	}()
	return funcB()
}

func funcB() error {
	// simulation
	//panic("foo")
	return errors.New("success")
}

func test() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}