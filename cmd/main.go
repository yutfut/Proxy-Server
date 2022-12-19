package main

import (
	"example.com/m/internal/app/proxy"
	"log"
	"time"
)

func main() {
	server := server.Server{
		Addr:         "8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	for  {
		go func() {
			log.Fatal(server.ListenAndServe())
		}()
	}
}
