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
	interceptor:=func(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(resp interface{},err error){
		md,ok:=metadata.FromIncomingContext(ctx)
		if !ok{
			return resp,status.Error(codes.Unauthenticated,"无token认证信息")
			fmt.Println("get metadata error")
		}

		for key,val:= range md{
			fmt.Println(key,val)
		}
		if nameSlice,ok:=md["appid"];ok{
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
