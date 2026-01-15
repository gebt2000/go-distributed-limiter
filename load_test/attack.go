package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	totalRequests := 20
	url := "http://localhost:8080"

	fmt.Println("🚀 Starting Load Test...")
	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Request %d failed: %v\n", id, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				fmt.Printf("✅ Request %d: ALLOWED\n", id)
			} else if resp.StatusCode == 429 {
				fmt.Printf("⛔ Request %d: BLOCKED\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nTest finished in %v\n", time.Since(start))
}
