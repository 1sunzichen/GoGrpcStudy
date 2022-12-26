package client_proxy

import (
	"PackageFive/new_helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct{
	*rpc.Client
}
func NewProxyClient(network string,address string)HelloServiceStub{
	client,_:=rpc.Dial(network,address)
	return HelloServiceStub{client}
}
func (H *HelloServiceStub)Hello(req string,reply *string)error{
	H.Call(handler.HandleName+".Hellohandle",req,&reply)
	return nil
}
