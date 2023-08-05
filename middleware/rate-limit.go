package middleware

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"tuanm.dev/golb/config"
)

var ctx = context.Background()

func GetClientIP(r *http.Request) string {
	var ip string
	ip = r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwared-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func RateLimit(next http.Handler, conf config.Config) http.Handler {
	rc := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPass,
		DB:       conf.RedisDB,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var total = 0
		ip := GetClientIP(r)
		t, err := rc.Get(ctx, ip).Result()
		if err == nil {
			total, _ = strconv.Atoi(t)
		}
		rc.Set(ctx, ip, total+1, time.Duration(time.Minute)).Err()
		if total+1 > conf.RateLimit {
			w.WriteHeader(403)
			return
		}
		next.ServeHTTP(w, r)
	})
}
