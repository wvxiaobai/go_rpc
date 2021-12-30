package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go_rpc/hello_world/proto/hello"
)

func main() {
	req := hello.HelloRequest{
		Name:    "bobby",
		Age:     28,
		Courses: []string{"rpc", "go"},
	}
	rsp, _ := proto.Marshal(&req)
	fmt.Println(string(rsp))
	newPeq := hello.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newPeq)
	fmt.Println(newPeq.Age)
}
