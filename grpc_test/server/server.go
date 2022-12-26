package main

import (
	helloworldproto "PackageFive/grpc_test/proto/api"
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context,request *helloworldproto.HelloRequest)(*helloworldproto.HelloReply,error){
   return &helloworldproto.HelloReply{
   	Message: "京津冀"+request.Name,
   },nil
}

func main(){
	g:=grpc.NewServer()
	helloworldproto.RegisterGreeterServer(g,&Server{})
	lis,err:=net.Listen("tcp","0.0.0.0:8080")
	if err!=nil{
		panic("failed to listen:"+err.Error())
	}
	err=g.Serve(lis)
	if err!=nil{
		panic("failed to start grpc"+err.Error())
	}
}
