package main

import (
	"fmt"
	"go_rpc/hello_world/grpc/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "127.0.0.1:1235")
	reply := ""
	err := client.Hello("bobby1", &reply)
	if err != nil {
		panic("connect2 fail")
	}
	fmt.Println(string(reply))

}
