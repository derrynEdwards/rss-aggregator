package main

import (
	"log"
	"net/http"
	"time"
)

func middleWareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		location, _ := time.LoadLocation("EST")
		timestamp := time.Now().In(location).Format(time.RFC3339)
		log.Printf("%s %s %s %s", timestamp, r.Method, r.URL.Path, r.Referer())
		next.ServeHTTP(w, r)
	})
}
