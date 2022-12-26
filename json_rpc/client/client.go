package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	clientconn,err:=(net.Dial("tcp","localhost:1235"))
	if err!=nil{
		panic("调用失败1")
	}

	var reply string
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(clientconn))
	err=client.Call("HServer.Hellohandle","Ekko1235",&reply)
	if err!=nil{
		panic("调用失败2")
	}
	fmt.Println(reply)
}