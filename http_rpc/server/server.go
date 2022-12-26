package main

import (
	"io"
	"net/http"
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
	_=rpc.RegisterName("HServer",&helloServer{})
	//
	http.HandleFunc("/jsonprc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser=struct{
			io.Writer
			io.ReadCloser
		}{
			ReadCloser:request.Body,
			Writer:writer,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234",nil)
}