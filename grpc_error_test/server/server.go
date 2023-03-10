package main

import (
	helloproto "PackageFive/grpc_proto_test/protobuf/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context,request *helloproto.HelloRequest)(*helloproto.HelloReply,error){
	time.Sleep(5*time.Second)
	return nil,status.Errorf(codes.NotFound,"记录未找到%v",request.Name)
}
func (s *Server) SayBase(ctx context.Context,request *helloproto.Empty)(*helloproto.HelloReply,error){
	return &helloproto.HelloReply{
		Message: "京津冀"+string(request.G),
	},nil
}

func main(){
	interceptor:=func(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(resp interface{},err error){
		md,ok:=metadata.FromIncomingContext(ctx)
		if ok{
			fmt.Println("get metadata error")
		}
		for key,val:= range md{
			fmt.Println(key,val)
		}
		if nameSlice,ok:=md["name"];ok{
			fmt.Println(nameSlice)
			for i,e:=range nameSlice{
				fmt.Println(i,e)
			}
		}
		return handler(ctx,req)
	}
	opt:=grpc.UnaryInterceptor(interceptor)
	g:=grpc.NewServer(opt)
	helloproto.RegisterGreeterServer(g,&Server{})
	lis,err:=net.Listen("tcp","0.0.0.0:8080")
	if err!=nil{
		panic("failed to listen:"+err.Error())
	}
	err=g.Serve(lis)
	if err!=nil{
		panic("failed to start grpc"+err.Error())
	}
}
