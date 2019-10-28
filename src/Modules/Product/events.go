package main

import (
	"eather/lib/types"
	"fmt"
	"time"
)

var eventFuncs = map[string]types.EventFunc{"added": added}

var added = func(data ...interface{}) {
	fmt.Println(data)
	time.Sleep(2 * time.Second)
	fmt.Println("Running event after added product")
}
