package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project/lib"

	"time"

	"github.com/joho/godotenv"
)

func main() {
	launch()

	srv := &http.Server{
		Handler:      nil,
		Addr:         os.Getenv("APP_URL"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func launch() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("USE_CRONS") == "1" {
		fmt.Println("Crons were enabled.")
		lib.StartCrons()
	}

	if os.Getenv("USE_CACHE") == "1" {
		fmt.Println("Cache was enabled.")
		lib.GetCache()
	}

	lib.GetDb()
	lib.InitVersion()
	lib.LoadModules()
	lib.RegisterRoutes()
}
