package config_test

import (
	"fmt"
	"test/business/config"
	"testing"

	"github.com/ardanlabs/conf/v2"
)

func TestConfig(t *testing.T) {
	config, _ := config.NewConfig()
	dbconfig := config.GetDb()
	t.Log(dbconfig.Host)

	out, err := conf.String(config)
	if err != nil {
		t.Fatal(fmt.Errorf("generating config for output: %w", err))
		return
	}
	t.Log("startup", "config", out)
}
