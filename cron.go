package eather

import (
	"github.com/robfig/cron"
)

// Cron structure
type Cron struct {
	Spec Spec
	Cmd  Cmd
}

// Spec for specification when cron should be running
type Spec string

// Cmd function of cron
type Cmd func()

// CronList structure of list of crons
type CronList []Cron

// StartCrons will start crons in background
func StartCrons(cronList []Cron) {
	c := cron.New()

	addCrons(c, cronList)

	r := GetRegistry()

	for _, r := range r.GetCollection() {
		if cronable := r.GetCronable(); cronable != nil {
			addCrons(c, cronable.Crons())
		}
	}

	c.Start()
}

func addCrons(c *cron.Cron, cronList CronList) {
	for _, cr := range cronList {
		c.AddFunc(string(cr.Spec), cr.Cmd)
	}
}
