package eather

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Start eather application initialize http server and load all modules
func Start(conf ConfigInterface) {
	// Initialize application and build modules
	launch(conf)

	// Start http server
	serve(conf)
}

func launch(conf ConfigInterface) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if os.Getenv("USE_CACHE") == "1" {
		fmt.Println("Cache was enabled.")
		GetCache()
	}

	GetDb()
	InitVersion()
	LoadModules(conf.GetModuleDirs())
	GetRouter()
	RegisterRoutes(conf.GetCorsOpts())

	if os.Getenv("USE_CRONS") == "1" {
		fmt.Println("Crons were enabled.")
		StartCrons(conf.GetCrons())
	}
}

func serve(conf ConfigInterface) {
	log.Fatal(conf.GetServerConf().ListenAndServe())
}
