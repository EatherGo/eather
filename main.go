package main

import (
	"log"
	"net/http"
	"os"
	"project/app"
	"time"
)

func main() {
	app.Launch()

	srv := &http.Server{
		Handler:      nil,
		Addr:         os.Getenv("APP_URL"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
