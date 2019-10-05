protoc -I $GOPATH/src --go_out=$GOPATH/src/github.com/ydsxiong/ $GOPATH/src/github.com/ydsxiong/summingservice/gRPC/proto-files/domain/numbers.proto

protoc -I $GOPATH/src --go_out=plugins=grpc:$GOPATH/src/github.com/ydsxiong/ $GOPATH/src/github.com/ydsxiong/summingservice/gRPC/proto-files/service/numbers-service.proto
