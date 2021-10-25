package main
// server
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_service/pb/person"
	"net"
	"time"
)

// 取出未注册的server
type  personServe struct {
	person.UnimplementedSearchServiceServer
}
// grpc服务挂载方法  外部调用我们的时候的处理逻辑
func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name:= req.GetName()
	res := &person.PersonRes{Name:"我收到了"+name+"的信息"}
	return res, nil
}
func (*personServe) SearchIn(server person.SearchService_SearchInServer) error {
	for {
		req,err := server.Recv() //一直去读
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name:"完成了"})
			break
		}
	}
	return nil
}
func (*personServe) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.Name
	i := 0
	for {
		if i > 10 {
			break
		}
		time.Sleep(1*time.Second)
		i++
		server.Send(&person.PersonRes{Name: "我拿到了"+name})
	}
	return nil
}
func (*personServe) SearchIo(server person.SearchService_SearchIoServer) error {
	i := 0
	str := make(chan string)
	go func() {
		for  {
			i++
			req,_ := server.Recv()
			if i > 10{
				str <- "结束"
				break
			}
			str <- req.Name
		}
	}()
	for  {
		s := <- str
		if s == "结束" {
			break
		}
		server.Send(&person.PersonRes{Name:s })
	}

	return nil
}

func main(){
	// 注册服务
	l,_ := net.Listen("tcp",":8888") //创建Listen
	s := grpc.NewServer() // 创建grpc.service
	person.RegisterSearchServiceServer(s,&personServe{}) //注册服务
	s.Serve(l) //监听
}
