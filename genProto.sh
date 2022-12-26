function genProto {
   DOMAIN=$1
   DO=$2
    PROTO_PATH=./${DOMAIN}/protobuf
    GO_OUT_PATH=./${DOMAIN}/protobuf/api
    protoc -I=$PROTO_PATH --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH ${DO}.proto

}


genProto metadata_test base
genProto metadata_test hello
