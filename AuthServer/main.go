package main

import (
	"context"
	"crypto/sha1"
	"fmt"
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
	listener, err := net.Listen("tcp", "0.0.0.0:5052")
	if err != nil {
		fmt.Print(err)
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
	CacheBInRedis(redis_key, BValue)

	return &pb.Req_DHParams_Response{
		Nonce:       req.Nonce,
		ServerNonce: req.ServerNonce,
		MessageId:   req.MessageId + 1,
		B:           int64(BValue),
	}, nil
}

func (s *server) ReqPq(ctx context.Context, req *pb.ReqPq_Request) (*pb.ReqPq_Response, error) {
	sNonce := ServerNonceGenerator()
	redis_key := string(sha1.New().Sum([]byte(req.Nonce + sNonce)))
	CacheAllInRedis(redis_key, req.Nonce, sNonce, req.MessageId+1, 5, 23)
	return &pb.ReqPq_Response{
		Nonce:       req.Nonce,
		ServerNonce: sNonce,
		MessageId:   req.MessageId + 1,
		P:           23,
		G:           5,
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

func CacheBInRedis(key string, value float64) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	rdb.Set(ctx, key, value, 20*time.Minute)
}

func CacheAllInRedis(key string, nonce string, serverNonce string, message_id int64, g int64, p int64) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	rdb.HSet(ctx, key, "nonce", nonce)
	rdb.HSet(ctx, key, "serverNonce", serverNonce)
	rdb.HSet(ctx, key, "message_id", message_id)
	rdb.HSet(ctx, key, "g", 5)
	rdb.HSet(ctx, key, "p", 23)
}
