package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"simple-rate-limiting/redisratelimit"
)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the IP address of the current user
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("SplitHostPort::Error::%s", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		key := fmt.Sprintf("IP_%s", ip)

		var limit redisratelimit.Limit
		var isRateLimiter bool

		// when rate limiter name unknown, no requests limits
		switch env.RateLimiterName {
		case "FixedWindow":
			limit, err = rdb.FixedWindowRateLimiter(key, env.MaxQuotas, env.WindowSeconds)
			isRateLimiter = true
		case "SlidingWindowLogs":
			limit, err = rdb.SlidingWindowLogsRateLimiter(key, env.MaxQuotas, env.WindowSeconds)
			isRateLimiter = true
		}

		if err != nil {
			log.Printf("FixedWindowRateLimiter::Error::%s", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if isRateLimiter {
			w.Header().Set("X-Rate-Limit-Limit", fmt.Sprint(env.MaxQuotas))
			w.Header().Set("X-Rate-Limit-Remaining", fmt.Sprint(limit.Remaining))
			w.Header().Set("X-Rate-Limit-Reset", fmt.Sprint(limit.RetryAfterSeconds))

			if limit.Remaining < 0 {
				msg := fmt.Sprintf("Rate limit exceeded. Try again in %d seconds.", limit.RetryAfterSeconds)
				http.Error(w, msg, http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
