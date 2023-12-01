package config_test

import (
	"test/business/config"
	"testing"
)

func TestConfig(t *testing.T) {
	config, _ := config.NewConfig()
	dbconfig := config.GetDb()
	t.Log(dbconfig.Host)
}
