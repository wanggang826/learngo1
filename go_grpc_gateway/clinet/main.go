package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_gateway/pb/person"
	"sync"
	"time"
)

// client

func main(){
	l,_ := grpc.Dial("127.0.0.1:8888",grpc.WithInsecure()) //创建一个链接，WithInsecure不做安全措施
	client := person.NewSearchServiceClient(l) // 创建一个客户端
	c,_ := client.SearchIo(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for  {
			time.Sleep(1*time.Second)
			err := c.Send(&person.PersonReq{Name: "刚刚"})
			if err != nil{
				wg.Done()
				break
			}
		}
	}()
	wg.Add(1)
	go func() {
		for  {
			res,err := c.Recv()
			if err != nil{
				fmt.Println(err)
				wg.Done()
				break
			}
			fmt.Println(res)

		}
	}()
	wg.Wait()
	//c,_ := client.SearchOut(context.Background(),&person.PersonReq{Name: "刚刚"})
	//for  {
	//	req,err := c.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(req)
	//}

	//c,_ := client.SearchIn(context.Background())
	//i:=0
	//for  {
	//	if i > 10 {
	//		res,_ := c.CloseAndRecv()
	//		fmt.Println(res)
	//		break
	//	}
	//	time.Sleep(1*time.Second)
	//	c.Send(&person.PersonReq{Name:"流入:我是进来的信息"})
	//	i++
	//}

	//res,err := client.Search(context.Background(),&person.PersonReq{Name:"刚刚"}) // 远程调用方法
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
}
