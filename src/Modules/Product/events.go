package main

import (
	"fmt"
	"time"
)

var eventFuncs = map[string]func(){"added": added}

var added = func() {
	time.Sleep(2 * time.Second)
	fmt.Println("Running event after added product")
}
