package main

import (
	"flag"
	"fmt"
)

var (
	generate = flag.Bool("admin", false, "Create admin user")
)

func main() {
	flag.Parse()

	if *generate {
		createAdmin()
		return
	}

	fmt.Println("Nothing to do")
}

func createAdmin() {

}
