package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type helloServer struct{

}
func (h *helloServer)Hellohandle(req string,reply *string)error{
	*reply="h"+req
	return nil
}
func main(){
	//
	listen,_:=net.Listen("tcp","localhost:1235")
	//
	_=rpc.RegisterName("HServer",&helloServer{})
	for{
		conn,_:=listen.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}