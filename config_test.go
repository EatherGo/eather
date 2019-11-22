package eather

import (
	"net/http"
	"os"
	"testing"

	"github.com/rs/cors"
	"github.com/stretchr/testify/assert"
)

var c = GetConfig()

func TestGetConfigShouldReturnEmptyConfig(t *testing.T) {
	c1 := GetConfig()

	assert.Equal(t, &Config{}, c1, "Config should be empty")
}

func TestCronsCanBeAddedToConfig(t *testing.T) {
	assert.Equal(t, 0, len(c.GetCrons()), "Crons collection should be empty")

	c.AddCron("* * * * *", func() {})

	assert.Equal(t, 1, len(c.GetCrons()), "Crons collection should have one cron")
}

func TestModulesDirsCanBeAddedToConfig(t *testing.T) {
	c1 := GetConfig()

	c1.AddModuleDirs("testing/dir")
	assert.Equal(t, 1, len(c1.GetModuleDirs()), "It should contains only 1 added dir")
}

func TestCorsCanBeSetForConfig(t *testing.T) {
	assert.Equal(t, defaultCorsOpts(), c.GetCorsOpts(), "If Cors are not set it should be set to default")

	c1 := GetConfig()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONTEND_URL")},
	})
	c1.SetCorsOpts(cors)

	assert.Equal(t, cors, c1.GetCorsOpts(), "Cors should be the same as cors variable")
}

func TestServerConfigCanBeSetForConfig(t *testing.T) {
	assert.Equal(t, defaultServerConf(), c.GetServerConf(), "If server conf is not set it should be default")

	c1 := GetConfig()

	server := &http.Server{
		Handler: nil,
	}

	c1.SetServerConfig(server)
	assert.Equal(t, server, c1.GetServerConf(), "It should be equal to server variable")
}
