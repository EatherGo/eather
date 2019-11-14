package lib

import (
	"github.com/robfig/cron"
)

// Cron structure
type Cron struct {
	spec string
	cmd  func()
}

// StartCrons will start crons in background
func StartCrons(cronList []Cron) {
	c := cron.New()

	for _, cr := range cronList {
		c.AddFunc(cr.spec, cr.cmd)
	}

	c.Start()
}
