package main

import (
	"eather/lib/interfaces"
	"fmt"
	"time"
)

var eventFuncs = map[string]interfaces.EventFunc{"added": added}

var added = func(data ...interface{}) {
	fmt.Println(data)
	time.Sleep(2 * time.Second)
	fmt.Println("Running event after added product")
}
