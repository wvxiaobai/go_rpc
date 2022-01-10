package main

import (
	"context"
	"fmt"
	"go_rpc/hello_world/grpcs/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := hello.NewHelloClient(conn)
	r, err := c.Hello(context.Background(), &hello.HelloRequest{Name: "wv"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Reply)

}
