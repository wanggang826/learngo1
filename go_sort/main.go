package main

import "fmt"

func main(){
	var arr = []int{3,44,5,7,20,11,400,29}
	//sortMp(arr)
	//sortXz(arr)
	//sortCr(arr)
	sortKs(arr,0,7)
	for k,v := range arr {
		fmt.Println("k",k,"v",v)
	}
}

// sortMp 冒泡 一次比较两个 顺序错误就互换位置  直到无需交换
func sortMp(arr []int){
	size := len(arr)
	for i:= 0; i< size -1;i++{
		for j:= 0; j < size - 1 -i;j++{
			fmt.Println("i",i,"j",j)
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				fmt.Println("temp",temp)
				arr[j+1] = arr[j]
				fmt.Println("arr[j+1]",arr[j])
				arr[j] = temp
				fmt.Println("arr[j]",arr[j])
				println("j",j)
			}
		}
	}
}

// sortXz 选择排序 先找最大的放首位 剩下元素中接着找最大的  以此类推
func sortXz(arr []int){
	size := len(arr)
	for i:=0; i < size - 1;i++{
		maxIndex := i
		for j := i+1;j<size;j++{
			if arr[j]  < arr[maxIndex]{
				maxIndex = j
			}
		}
		temp := arr[i]
		arr[i] = arr [maxIndex]
		arr[maxIndex] = temp
	}
}

// sortCr 插入排序  构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
func sortCr(arr []int)  {
	size := len(arr)
	for i:=1; i <= size - 1;i++{
		for j := i;j > 0;j--{
			if arr[j-1]  < arr[j]{
				arr[j-1],arr[j] = arr[j],arr[j-1]
			}
		}
	}
}
// sortKs 快速排序
// 通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，
// 则可分别对这两部分记录继续进行排序，以达到整个序列有序

// 从数列中挑出一个元素，称为 “基准”（pivot）；
// 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面
//（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
func sortKs(arr []int,l,r int){
	if l < r {
		pivot := arr[r]
		i := l -1
		for j := l;j < r ;j++{
			if arr[j] < pivot {
				i ++
				arr[j],arr[i] = arr[i],arr[j]
			}
		}
		i++
		arr[r],arr[i] = arr[i],arr[r]
		sortKs(arr,l,i-1)
		sortKs(arr,i+1,r)
	}
}

