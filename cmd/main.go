package main

import (
	"example.com/m/internal/app/proxy"
	"log"
	"time"
)

func main() {
	s := server.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
