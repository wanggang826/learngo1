package main
// server
import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpc_gateway/pb/person"
	"net"
	"net/http"
	"sync"
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

func main(){
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateway(&wg)
	go registerGRPC(&wg)
	wg.Wait()

}

func registerGateway(wg *sync.WaitGroup){
	conn,_ := grpc.DialContext( // 访问grpc服务的中间层
		context.Background(),
		"127.0.0.1:8888",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	mux := runtime.NewServeMux()
	gwServer := &http.Server{
		Handler: mux,
		Addr: ":8090",
	}

	person.RegisterSearchServiceHandler(context.Background(),mux,conn)
	gwServer.ListenAndServe()
	wg.Done()
}


func registerGRPC(wg *sync.WaitGroup){
	// 注册服务
	l,_ := net.Listen("tcp",":8888") //创建Listen
	s := grpc.NewServer() // 创建grpc.service
	person.RegisterSearchServiceServer(s,&personServe{}) //注册服务
	s.Serve(l) //监听
	wg.Done()
}