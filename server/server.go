package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"gRPC_todo/constant"
	"gRPC_todo/proto"
)

func RunServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", constant.Port))
	if err != nil {
		log.Printf("failed to run server: %v", err.Error())
		return
	}

	grpcServer := grpc.NewServer()
	todoHandler := InjectTodoAPI()
	proto.RegisterTodoCommandServer(grpcServer, todoHandler)
	proto.RegisterTodoQueryServer(grpcServer, todoHandler)

	log.Printf("Listening on %v:", constant.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to run server: %v", err.Error())
		return
	}
}
