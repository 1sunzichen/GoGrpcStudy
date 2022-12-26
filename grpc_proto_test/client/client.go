package main

import (
	helloproto "PackageFive/grpc_proto_test/protobuf/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main(){
	conn,err:=grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	defer conn.Close()
	u:=helloproto.Empty_Result{
		Url: "www.baidu.com",

	}
	c:=helloproto.NewGreeterClient(conn)
	r,err:=c.SayHello(context.Background(),&helloproto.HelloRequest{Name:"boddy"})
	if err!=nil{
		panic(err)
	}
	r,err=c.SayBase(context.Background(),&helloproto.Empty{
		Data:[]*helloproto.Empty_Result{
			{Url: "111"},
			{Url:"www.baidu.com"},
		},
		G:helloproto.Gender_MALE,
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

