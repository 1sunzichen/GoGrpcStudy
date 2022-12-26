package main

import (
	streamproto "PackageFive/stream_grpc_test/proto/api"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT=":50052"
type server struct{

}
func (s *server) GetStream(req *streamproto.StreamReqData,res streamproto.Greeter_GetStreamServer)error{
	i:=0
	for{
		i++
		res.Send(&streamproto.StreamResData{
			Data:fmt.Sprintf("%v",time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i>10{
			break
		}
	}
	return nil
}
func (s *server) PutStream(cliStr streamproto.Greeter_PutStreamServer)error{
	for {
		if a,err:=cliStr.Recv();err!=nil{
			fmt.Println(err)
		}else{
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (s *server) AllStream(allStr streamproto.Greeter_AllStreamServer)error{
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
			_=allStr.Send(&streamproto.StreamResData{Data:"服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}
func main(){

	lis,err:=net.Listen("tcp",PORT)
	if err!=nil{
		panic(err)
	}
	s:=grpc.NewServer()
	streamproto.RegisterGreeterServer(s,&server{})
	err=s.Serve(lis)
	if err!=nil{
		panic(err)
	}
}

