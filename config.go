package eather

import "github.com/EatherGo/eather/lib"

// Config structure of config for Eather
type Config struct {
	cronlist lib.CronList
}

// AddCron will append new cron to the config
func (c *Config) AddCron(spec lib.Spec, cmd lib.Cmd) {
	cron := lib.Cron{
		Spec: spec,
		Cmd:  cmd,
	}

	c.cronlist = append(c.cronlist, cron)
}

func (c *Config) getCrons() lib.CronList {
	return c.cronlist
}
