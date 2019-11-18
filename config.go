package eather

import "os"

// Config structure of config for Eather
type Config struct {
	cronlist   CronList
	moduleDirs []string
}

// ConfigInterface interface of config
type ConfigInterface interface {
	AddCron(spec Spec, cmd Cmd)
	AddModuleDirs(dir ...string)
	GetCrons() CronList
	GetModuleDirs() []string
}

// GetConfig will return default config settings
func GetConfig() ConfigInterface {
	return &Config{
		cronlist:   CronList{},
		moduleDirs: []string{},
	}
}

// AddCron will append new cron to the config
func (c *Config) AddCron(spec Spec, cmd Cmd) {
	cron := Cron{
		Spec: spec,
		Cmd:  cmd,
	}

	c.cronlist = append(c.cronlist, cron)
}

// AddModuleDirs will replace moduleDirs by new
func (c *Config) AddModuleDirs(dir ...string) {
	for _, d := range dir {
		c.moduleDirs = append(c.moduleDirs, d)
	}
}

// GetCrons returns slice of Cron
func (c *Config) GetCrons() CronList {
	return c.cronlist
}

// GetModuleDirs returns directories of modules
func (c *Config) GetModuleDirs() []string {
	if len(c.moduleDirs) == 0 {
		return defaultModuleDirs()
	}
	return c.moduleDirs
}

func defaultModuleDirs() []string {
	return []string{os.Getenv("CORE_MODULES_DIR"), os.Getenv("CUSTOM_MODULES_DIR")}
}
