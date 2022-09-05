package main

import (
	"context"
	"fmt"
	coreProto "golang/simple-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// HelloService ...
type HelloService struct {
	coreProto.UnimplementedHelloServiceServer
}

// SayHello this will be implementation method hello world
func (s *HelloService) SayHello(ctx context.Context, in *coreProto.HelloRequest) (*coreProto.HelloResponse, error) {
	return &coreProto.HelloResponse{
		Message: "Hello From the Server !",
		Name: in.Name,
	}, nil
}

func main() {

	/**
	go get google.golang.org/protobuf/runtime/protoimpl@v1.26.0
	go get google.golang.org/protobuf/reflect/protoreflect@v1.26.0
	Generate proto enter folder go-simple-grpc then execute this command
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto
	*/

	serverAddress := "9090"
	srv, err := net.Listen("tcp", fmt.Sprintf(`:%s`, serverAddress))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	coreProto.RegisterHelloServiceServer(grpcServer, new(HelloService))

	fmt.Println(fmt.Sprintf(`Server running in address : %s`, serverAddress))

	if err = grpcServer.Serve(srv); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}