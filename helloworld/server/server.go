package main

import (
	"net"
	"net/rpc"
)

type helloServer struct{

}
func (h *helloServer)Hellohandle(req string,reply *string)error{
	*reply="h"+req
	return nil
}
func main(){
	//
	listen,_:=net.Listen("tcp","localhost:1234")
	//
	_=rpc.RegisterName("HServer",&helloServer{})
	conn,_:=listen.Accept()
	rpc.ServeConn(conn)
}