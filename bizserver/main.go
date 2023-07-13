package main

import (
	"fmt"
	// "main/service"
	gen "main/gen/go"
	"net"

	"google.golang.org/grpc"
)

func main() {
	Create_database()
	listener, err := net.Listen("tcp", "0.0.0.0:5062")
	if err != nil {
		panic(err)
	}
	fmt.Println("gRPC biz server running on " + listener.Addr().String())

	s := grpc.NewServer()
	gen.RegisterBizServer(s, &server{})
	fmt.Println(s.GetServiceInfo())
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
