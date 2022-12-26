package main

import (
	helloworldpb "PackageFive/helloworld/protobuf/api"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main(){
	//序列化
	req:=helloworldpb.HelloRequest{
		Name:"ekko",
		Age: 12,

	}
	rsp,_:=proto.Marshal(&req)//具体的编码是如何做到的。protobuf 的原理 varint
	fmt.Println(string(rsp))
}
