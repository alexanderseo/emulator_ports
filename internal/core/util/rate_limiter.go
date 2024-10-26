package util

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func RateLimiter(l *log.Logger, next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(5, 10)
	l.Println("rate limiter")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			message := Message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later.",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			err := json.NewEncoder(w).Encode(&message)
			if err != nil {
				return
			}
			return
		} else {
			next(w, r)
		}
	})
}
