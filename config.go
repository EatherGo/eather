package eather

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

// Config structure of config for Eather
type Config struct {
	cronlist   CronList
	moduleDirs []string
	corsOpts   *cors.Cors
	serverConf *http.Server
}

// ConfigInterface interface of config
type ConfigInterface interface {
	AddCron(spec Spec, cmd Cmd)
	AddModuleDirs(dir ...string)
	SetCorsOpts(cors *cors.Cors)
	SetServerConfig(*http.Server)

	GetCrons() CronList
	GetModuleDirs() []string
	GetCorsOpts() *cors.Cors
	GetServerConf() *http.Server
}

// GetConfig will return default config settings
func GetConfig() ConfigInterface {
	return &Config{}
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

// SetCorsOpts will set cors for application
func (c *Config) SetCorsOpts(cors *cors.Cors) {
	c.corsOpts = cors
}

// SetServerConfig will set server configuration
func (c *Config) SetServerConfig(s *http.Server) {
	c.serverConf = s
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

// GetCorsOpts returns cor config
func (c *Config) GetCorsOpts() *cors.Cors {
	if c.corsOpts == nil {
		c.corsOpts = defaultCorsOpts()
	}
	return c.corsOpts
}

// GetServerConf returns server configuration
func (c *Config) GetServerConf() *http.Server {
	if c.serverConf == nil {
		c.serverConf = defaultServerConf()
	}

	return c.serverConf
}

func defaultModuleDirs() []string {
	return []string{os.Getenv("CORE_MODULES_DIR"), os.Getenv("CUSTOM_MODULES_DIR")}
}

func defaultCorsOpts() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONTEND_URL")}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})
}

func defaultServerConf() *http.Server {
	return &http.Server{
		Handler:      nil,
		Addr:         os.Getenv("APP_URL"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
