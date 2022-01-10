package main

import (
	"context"
	"go_rpc/hello_world/grpcs/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Reply: "hello" + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	hello.RegisterHelloServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("fail")
	}
	g.Serve(lis)
}
