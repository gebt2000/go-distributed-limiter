# Distributed Rate Limiter (Go + Redis)

A high-performance, distributed rate limiter built in Go. It uses **Redis Lua scripts** to ensure atomicity and prevent race conditions in concurrent environments.

## Key Features
* **Distributed State:** Uses Redis as the backend, allowing multiple API instances to share the same rate limits.
* **Atomic Operations:** Implements the "Fixed Window" algorithm via custom Lua scripts to guarantee thread safety.
* **Concurrency Safe:** Proven to handle concurrent requests without "leaking" excess traffic.
* **Dockerized:** Ready to run with a single command.

## 🛠 Tech Stack
* **Language:** Go (Golang)
* **Database:** Redis (Alpine)
* **Containerization:** Docker & Docker Compose
* **Testing:** Custom concurrent load generator

## How to Run
1. **Start Redis:**
   ```bash
   docker-compose up -d
   ```

2. **Run the Server:**
   ```bash
   go run main.go
   ```
   Server will start on `:8080`.

3. **Run the Load Test:**
   ```bash
   go run load_test/attack.go
   ```

## Performance Proof
Load test results simulating 20 concurrent requests (Limit: 5/10s):

```text
🚀 Starting Load Test...
✅ Request 10: ALLOWED
✅ Request 9: ALLOWED
✅ Request 1: ALLOWED
✅ Request 6: ALLOWED
✅ Request 7: ALLOWED
⛔ Request 14: BLOCKED
... (All subsequent requests blocked)
Test finished in 19.69ms
```
