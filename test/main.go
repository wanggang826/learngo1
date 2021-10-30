package main

import (
	"fmt"
	"reflect"
)

func main()  {
	DeferCall()
}

func DeferCall() {
	//defer func() { fmt.Println("打印前") }()
	//defer func() { fmt.Println("打印中") }()
	//defer func() { fmt.Println("打印后") }()
	//fmt.Println("test")
	//panic("触发异常")

	paseStudent()

	//runtime.GOMAXPROCS(1)
	//wg := sync.WaitGroup{}
	//wg.Add(20)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println("J: ", i)
	//		wg.Done()
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("i: ", i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()


}

type student struct {
	Name string
	Age  int
}

func paseStudent() {
	//m := make(map[string]*student)
	//stus := []student{
	//	{Name: "zhou", Age: 24},
	//	{Name: "li", Age: 23},
	//	{Name: "wang", Age: 22},
	//}
	////for _, stu := range stus {
	////	m[stu.Name] = &stu
	////}
	//for i:=0;i<len(stus);i++{
	//	m[stus[i].Name] = &stus[i]
	//}
	//fmt.Println(m["li"].Name)
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	t := reflect.TypeOf(myMap)
	fmt.Println("type:", t)
	v := reflect.ValueOf(myMap)
	fmt.Println("value:", v)

	//a := []int{}
	//b := []int{1,2,3}
	//c := a
	//a = append(b, 1)
	//fmt.Println(a,b,c)

	//mySlice := []int{10, 20, 30, 40, 50}
	//for _, value := range mySlice {
	//	value *= 2
	//}
	//fmt.Printf("mySlice %+v\n", mySlice)
	//for index, _ := range mySlice {
	//	mySlice[index] *= 2
	//}
	//fmt.Printf("mySlice %+v\n", mySlice)

}


//type student struct {
//	Name string
//	Age  int
//}
//
//func paseStudent() {
//	m := make(map[string]*student)
//	stus := []student{
//		{Name: "zhou", Age: 24},
//		{Name: "li", Age: 23},
//		{Name: "wang", Age: 22},
//	}
//	for _, stu := range stus {
//		m[stu.Name] = &stu
//	}
//}
