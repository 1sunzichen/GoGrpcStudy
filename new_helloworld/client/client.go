package main

import (
	"PackageFive/new_helloworld/client_proxy"
	"fmt"
)

func main()  {
	client:=client_proxy.NewProxyClient("tcp","localhost:1234")


	var reply string

	err:=client.Hello("Ekko",&reply)
	if err!=nil{
		panic("调用失败2")
	}
	fmt.Println(reply)
}