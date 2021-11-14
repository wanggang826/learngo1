package main

import "fmt"

// array vs slice
// array
// 长度固定

// 数组是值类型,复制和传参复制整个数组。因此改变副本得值不会改变自身的值

// 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。

// [n]*T表示指针数组，*[n]T表示数组指针



//slice
// 切片是一个引用类型，它的内部结构包含地址(底层数组的指针)、长度和容量。切片一般用于快速地操作一块数据集合

// 切片是引用类型，不支持直接比较，只能和nil比较

// 切片之间是不能比较的,判断切片是否为空 len(s) == 0  不用 s == nil

// 当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
// “扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。



func main()  {
	// array
	var testArray [3]int
	var numArray = [...]int{1, 2}
	cityArray := [...]string{"北京", "上海", "深圳"}  // 多维数组只有第一层可以使用...来让编译器推导数组长度
	arr := [...]int{1,2,3,4,5,6,7}
	fmt.Println(testArray,numArray,cityArray,arr)
	//slice
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)

	// 引用类型 赋值拷贝演示

	s1 := make([]int,3)
	s2 := s1
	s2[0] = 100
	s2 = append(s2,999)
	s3 := s1[1:2]
	s3 = append(s3, 888)
	fmt.Println(s1) //[100 0 0]


	sa := arr[0:5]
	sb := sa[1:3]
	//sc := arr[1:5]
	//sd := arr[0:6]
	sb = append(sb, 100,200,300,400,500)
	//sc = append(sc, 100)
	//sd = append(sd, 100)
	fmt.Println(sa,sb,arr) // [1 2 3 100 5] [2 3 100] [1 2 3 100 5 6 7]
	//make([]T, size, cap) //创建切片
	c := make([]int, 2, 10)
	fmt.Println(c)


}
