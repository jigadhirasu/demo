::golang
protoc -I=protofile ^
 --go_out=C:\_workspace ^
 protofile/product/*.proto 

protoc -I=protofile ^
 --go-grpc_out=C:\_workspace ^
 protofile/service/*.proto