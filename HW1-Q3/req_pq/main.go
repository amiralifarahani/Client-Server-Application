package main

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"example.com/pb"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type server struct {
	pb.ReqPqAuthenticationServiceServer
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Print(err)
	}
	s := grpc.NewServer()
	pb.RegisterReqPqAuthenticationServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) ReqPq(ctx context.Context, req *pb.ReqPq_Request) (*pb.ReqPq_Response, error) {
	sNonce := ServerNonceGenerator()
	redis_key := string(sha1.New().Sum([]byte(req.Nonce + sNonce)))
	CacheInRedis(redis_key, req.Nonce, sNonce, req.MessageId+1, 5, 23)
	return &pb.ReqPq_Response{
		Nonce:       req.Nonce,
		ServerNonce: sNonce,
		MessageId:   req.MessageId + 1,
		P:           11,
		G:           3,
	}, nil
}

func ServerNonceGenerator() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CacheInRedis(key string, nonce string, serverNonce string, message_id int64, g int64, p int64) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	rdb.HSet(ctx, key, "nonce", nonce)
	rdb.HSet(ctx, key, "serverNonce", serverNonce)
	rdb.HSet(ctx, key, "message_id", message_id)
	rdb.HSet(ctx, key, "g", 5)
	rdb.HSet(ctx, key, "p", 23)
}
