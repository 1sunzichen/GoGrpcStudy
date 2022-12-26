package main

import (
	"fmt"
	"net/rpc"
)

func main()  {
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		panic("调用失败1")
	}

	var reply string

	err=client.Call("HServer.Hellohandle","Ekko",&reply)
	if err!=nil{
		panic("调用失败2")
	}
	fmt.Println(reply)
}