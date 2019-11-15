package eather

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config structure of config for Eather
type Config struct {
	cronlist   CronList
	moduleDirs []string
}

// GetConfig will return default config settings
func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		cronlist:   CronList{},
		moduleDirs: []string{os.Getenv("CORE_MODULES_DIR"), os.Getenv("CUSTOM_MODULES_DIR")},
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
	return c.moduleDirs
}
