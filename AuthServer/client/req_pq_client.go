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
	request := &pb.ReqPq_Request{Nonce: "bardia rezaei kalantari", MessageId: 1823756}

	resp, err := client.ReqPq(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receive response => ", resp.Nonce)
	fmt.Println("Receive response => ", resp.ServerNonce)
	fmt.Println("Receive response => ", resp.MessageId)
	fmt.Println("Receive response => ", resp.P)
	fmt.Println("Receive response => ", resp.G)
}
