package main

import (
	"log"
	"net/http"
	"project/app"
)

func main() {
	app.Launch()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
