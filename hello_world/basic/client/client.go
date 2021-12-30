package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic("connect fail")
	}

	reply := ""
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("connect2 fail")
	}

	fmt.Println(string(reply))

}
