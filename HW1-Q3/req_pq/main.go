package main

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Req struct {
	Nonce      string `json:"nonce"`
	Message_Id int64  `json:"message_id"`
}

type Res struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"serverNonce"`
	Message_Id  int64  `json:"message_id"`
	P           int64  `json:"p"`
	G           int64  `json:"g"`
}

func main() {
	http.HandleFunc("/", grpc1)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func grpc1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	serverNonce := ServerNonceGenerator()

	redis_key := string(sha1.New().Sum([]byte(req.Nonce + serverNonce)))

	CacheInRedis(redis_key, req.Nonce, serverNonce, req.Message_Id+1, 3, 11)

	response := Res{Nonce: req.Nonce, ServerNonce: serverNonce, Message_Id: req.Message_Id + 1, P: 11, G: 3}
	json.NewEncoder(w).Encode(response)

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
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})
	rdb.HSet(ctx, key, "nonce", nonce)
	rdb.HSet(ctx, key, "serverNonce", serverNonce)
	rdb.HSet(ctx, key, "message_id", message_id)
	rdb.HSet(ctx, key, "g", g)
	rdb.HSet(ctx, key, "p", p)
}
