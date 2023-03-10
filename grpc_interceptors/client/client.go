package main

import (
	interproto "PackageFive/grpc_interceptors/protobuf/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main(){
	interceptor:=func(ctx context.Context,method string,req,reply interface{},cc *grpc.ClientConn,invoker grpc.UnaryInvoker,opts ...grpc.CallOption)error{
		start:=time.Now();
		err:=invoker(ctx,method,req,reply,cc,opts...)
		fmt.Println(time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts=append(opts,grpc.WithInsecure())
	opts=append(opts,grpc.WithUnaryInterceptor(interceptor))
	//opt:=grpc.WithUnaryInterceptor(interceptor)

	conn,err:=grpc.Dial("127.0.0.1:8080",opts...)
	if err!=nil{
		panic(err)
	}
	defer conn.Close()
	u:=interproto.Empty_Result{
		Url: "www.baidu.com",

	}
	c:=interproto.NewGreeterClient(conn)
	r,err:=c.SayHello(context.Background(),&interproto.HelloRequest{Name:"boddy"})
	if err!=nil{
		panic(err)
	}
	r,err=c.SayBase(context.Background(),&interproto.Empty{
		Data:[]*interproto.Empty_Result{
			{Url: "111"},
			{Url:"www.baidu.com"},
		},
		G:interproto.Gender_MALE,
		Mp: map[string]string{
			"姓名":"小明",
		},
		Tt: timestamppb.New(time.Now()),
	})

	if err!=nil{
		panic(err)
	}
	fmt.Println(r.Message,u)

}

