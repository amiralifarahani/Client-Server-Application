package main

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Req struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"serverNonce"`
	Message_Id  int64  `json:"message_id"`
	A           int64  `json:"a"`
}

type Res struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"serverNonce"`
	Message_Id  int64  `json:"message_id"`
	B           int64  `json:"B"`
}

func main() {

	http.HandleFunc("/", grpc2)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func grpc2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	b := rand.Int63n(100)

	BValue := math.Mod(11, math.Pow(float64(req.A), float64(b)))

	response := Res{Nonce: req.Nonce, ServerNonce: req.ServerNonce, Message_Id: req.Message_Id + 1, B: int64(BValue)}

	redis_key := string(sha1.New().Sum([]byte(req.Nonce + req.ServerNonce)))
	redis_value := `nonce: "` + req.Nonce + `", serverNonce: "` + req.ServerNonce + `", message_id: "` + strconv.FormatInt(req.Message_Id+1, 10)  + `", B: "`+ strconv.FormatInt(int64(BValue),10) +`"`

	CacheInRedis(redis_key, redis_value)

	json.NewEncoder(w).Encode(response)

}

func CacheInRedis(key string, value string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})
	rdb.Set(context.Background(), key, value, 20*time.Minute)
}