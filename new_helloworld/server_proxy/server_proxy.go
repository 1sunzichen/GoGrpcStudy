package server_proxy

import (
	"PackageFive/new_helloworld/handler"
	"net/rpc"
)
type HelloServicer interface {
	Hello (req string,reply *string)error
}
func ResigisterHanle(srv HelloServicer)error{
	return rpc.RegisterName(handler.HandleName,srv)

}
