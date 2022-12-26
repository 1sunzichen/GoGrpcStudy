package main

import (
	helloworldproto "PackageFive/grpc_test/proto/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main(){
	conn,err:=grpc.Dial("127.0.0.1:8088",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	defer conn.Close()
	c:=helloworldproto.NewGreeterClient(conn)
	r,err:=c.SayHello(context.Background(),&helloworldproto.HelloRequest{Name:"boddy"})
	if err!=nil{
		panic(err)
	}
	fmt.Println(r.Message)

}
