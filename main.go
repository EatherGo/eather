package eather

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EatherGo/eather/lib"

	"time"

	"github.com/joho/godotenv"
)

var cronList []lib.Cron

// Start eather application initialize http server and load all modules
func Start(crons []lib.Cron) {
	cronList = crons

	lib.Cron{spec: "", cmd: func() { fmt.Println("t") }}
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
		log.Println("Error loading .env file")
	}

	if os.Getenv("USE_CRONS") == "1" {
		fmt.Println("Crons were enabled.")
		lib.StartCrons(cronList)
	}

	if os.Getenv("USE_CACHE") == "1" {
		fmt.Println("Cache was enabled.")
		lib.GetCache()
	}

	lib.GetDb()
	lib.InitVersion()
	lib.LoadModules()
	lib.GetRouter()
	lib.RegisterRoutes()
}
