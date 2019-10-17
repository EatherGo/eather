package app

import (
	"fmt"
	"log"
	"os"
	"project/lib"

	"github.com/joho/godotenv"
)

// Launch the application
func Launch() {
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
	lib.GetRouter()
	lib.LoadModules()
}
