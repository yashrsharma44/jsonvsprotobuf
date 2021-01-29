echo "Compiling Protobuf....\n"
protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. test_pb.proto