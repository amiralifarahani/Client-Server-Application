package main

import (
	"context"
	"example.com/pb"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:5052", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewAuthenticationServiceClient(cc)
	request := &pb.Req_DHParams_Request{Nonce: "bardia rezaei kalantari", MessageId: 1823756,ServerNonce: "123123",A: 1223}

	resp, err := client.Req_DHParams(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receive response => ", resp.B)
}
