package config

import (
	"errors"
	"fmt"

	"github.com/ardanlabs/conf/v2"
)

var build = "develop"

type Config struct {
	conf.Version
	Db db
}

type db struct {
	User         string `conf:"default:postgres"`
	Password     string `conf:"default:postgres,mask"`
	Host         string `conf:"default:localhost"`
	Name         string `conf:"default:postgres"`
	MaxIdleConns int    `conf:"default:0"`
	MaxOpenConns int    `conf:"default:0"`
	DisableTLS   bool   `conf:"default:true"`
}

// NewConfig constructs a Config struct which represents server settings, and populates it with default values.
func NewConfig() (*Config, error) {
	cfg := Config{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
		Db: db{},
	}

	const prefix = "Test"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil, err
		}
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return &cfg, nil
}
