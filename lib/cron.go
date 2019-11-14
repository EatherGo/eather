package lib

import (
	"github.com/robfig/cron"
)

// StartCrons will start crons in background
func StartCrons() {
	c := cron.New()
	// c.AddFunc("* * * * *", func() { fmt.Println("Every minute cron") })
	c.Start()
}
