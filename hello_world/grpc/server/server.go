package main

import (
	"go_rpc/hello_world/grpc/handler"
	"go_rpc/hello_world/grpc/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":1235")

	_ = server_proxy.RegisterHelloService(&handler.HelloService{})

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}

}
