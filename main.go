package eather

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Start eather application initialize http server and load all modules
func Start(conf ConfigInterface) {
	// Initialize application and build modules
	launch(conf)

	// Start http server
	serve()
}

func launch(conf ConfigInterface) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if os.Getenv("USE_CRONS") == "1" {
		fmt.Println("Crons were enabled.")
		StartCrons(conf.GetCrons())
	}

	if os.Getenv("USE_CACHE") == "1" {
		fmt.Println("Cache was enabled.")
		GetCache()
	}

	GetDb()
	InitVersion()
	LoadModules(conf.GetModuleDirs())
	GetRouter()
	RegisterRoutes()
}

func serve() {
	srv := &http.Server{
		Handler:      nil,
		Addr:         os.Getenv("APP_URL"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
