package main

import (
	streamproto "PackageFive/stream_grpc_test/proto/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main(){
	conn,err:=grpc.Dial("localhost:50052",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	defer conn.Close()
	c:=streamproto.NewGreeterClient(conn)
	res,_:=c.GetStream(context.Background(),&streamproto.StreamReqData{Data:"2022世界杯"})
	for{
		a,err:=res.Recv()
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(a)
	}
	//客户端流模式
	putS,_:=c.PutStream(context.Background())
	i:=0
	for{
		i ++
		_=putS.Send(&streamproto.StreamReqData{Data:fmt.Sprintf("2022世界杯%d",i)})
		time.Sleep(time.Second)
		if i>10{
			break
		}
	}
	//双向流模式
	allStr,_:=c.AllStream(context.Background())
	wg:=sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		for{
			data,_:=allStr.Recv()
			fmt.Println("收到客户端信息:",data.Data)
		}
	}()
	go func(){
		defer wg.Done()
		for{
			_=allStr.Send(&streamproto.StreamReqData{Data:"服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}