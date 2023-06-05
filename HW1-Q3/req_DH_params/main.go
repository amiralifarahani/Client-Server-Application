package main

import (
	"context"
	"crypto/sha1"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"time"

	"example.com/pb"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type server struct {
	pb.AuthenticationServiceServer
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Req_DHParams(ctx context.Context, req *pb.Req_DHParams_Request) (*pb.Req_DHParams_Response, error) {
	
	BValue := math.Mod(23, math.Pow(float64(req.A), float64(rand.Int63n(100))))
	redis_key := string(sha1.New().Sum([]byte(req.Nonce + req.ServerNonce)))
	CacheInRedis(redis_key, BValue)

	return &pb.Req_DHParams_Response{
		Nonce:       req.Nonce,
		ServerNonce: req.ServerNonce,
		MessageId:   req.MessageId + 1,
		B:           int64(BValue),
	}, nil
}

func CacheInRedis(key string, value float64) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	rdb.Set(ctx, key, value, 20*time.Minute)
}
