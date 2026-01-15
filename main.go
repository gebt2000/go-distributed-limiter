package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gebt2000/go-distributed-limiter/limiter"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	rateLimiter := limiter.NewRateLimiter(rdb)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		
		// 1. Get IP, stripping the port number (e.g., 127.0.0.1:54321 -> 127.0.0.1)
		ip := r.RemoteAddr
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			ip = host
		}
		
		// 2. Normalize Localhost IPv6 (::1) to IPv4
		if ip == "::1" {
			ip = "127.0.0.1"
		}

		// 3. Check Limit (5 requests per 10 seconds)
		allowed, err := rateLimiter.Allow(ctx, ip, 5, 10*time.Second)
		if err != nil {
			fmt.Printf("Redis Error: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			http.Error(w, "429 - Too Many Requests", http.StatusTooManyRequests)
			return
		}

		fmt.Fprintf(w, "Request Allowed! Welcome.")
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
