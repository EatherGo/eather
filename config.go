package eather

// import "github.com/EatherGo/eather/

// Config structure of config for Eather
type Config struct {
	cronlist CronList
}

// AddCron will append new cron to the config
func (c *Config) AddCron(spec Spec, cmd Cmd) {
	cron := Cron{
		Spec: spec,
		Cmd:  cmd,
	}

	c.cronlist = append(c.cronlist, cron)
}

func (c *Config) getCrons() CronList {
	return c.cronlist
}
