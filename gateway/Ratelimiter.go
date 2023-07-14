package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"log"
	"strconv"
	"time"
)

const (
	RateLimitPerSecond = 100
	BlockDuration      = 24 * time.Hour
)

var limit = ratelimit.New(RateLimitPerSecond)

func CheckBlocked() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIp := ctx.ClientIP()
		// Check if the user has exceeded the rate limit
		if blocked, err := isBlocked(ctx, clientIp); err != nil || blocked {
			ctx.AbortWithStatus(429)
			return
		}
		ctx.Next()
	}
}

func ApplyLeakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		sub := now.Sub(prev)
		log.Print(color.CyanString("%v", sub))
		if isRateLimitExceeded(sub) {
			ip := ctx.ClientIP()
			log.Print(color.RedString("Rate limit exceeded for %s", ip))
			blockClient(ctx, ip)
			ctx.AbortWithStatus(429)
			return
		}
		prev = now
	}
}

func isRateLimitExceeded(sub time.Duration) bool {
	return sub < 1/RateLimitPerSecond*time.Second
}

func isBlocked(ctx context.Context, clientIp string) (bool, error) {
	key := fmt.Sprintf("blocked_%s", clientIp)
	value := GetValue(ctx, key)
	if value == "" {
		return false, nil
	}
	blocked, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return blocked, nil
}

func blockClient(ctx context.Context, clientIp string) {
	key := fmt.Sprintf("blocked_%s", clientIp)
	value := strconv.FormatBool(true)
	CacheData(ctx, key, value, BlockDuration)
}
