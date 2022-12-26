package main

import (
	helloproto "PackageFive/grpc_proto_test/protobuf/api"
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context,request *helloproto.HelloRequest)(*helloproto.HelloReply,error){
	return &helloproto.HelloReply{
		Message: "京津冀"+request.Name,
	},nil
}
func (s *Server) SayBase(ctx context.Context,request *helloproto.Empty)(*helloproto.HelloReply,error){
	return &helloproto.HelloReply{
		Message: "京津冀"+string(request.G),
	},nil
}

func main(){
	g:=grpc.NewServer()
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
