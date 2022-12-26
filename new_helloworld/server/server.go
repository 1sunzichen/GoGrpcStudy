package main

import (
	"PackageFive/new_helloworld/handler"
	"PackageFive/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

type helloServer struct{

}

func main(){
	//
	listen,_:=net.Listen("tcp","localhost:1234")
	//
	_=server_proxy.ResigisterHanle(&handler.HelloService{})
	conn,_:=listen.Accept()
	rpc.ServeConn(conn)
}