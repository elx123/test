package envconfig

import (
	"log"

	evconfig "github.com/kelseyhightower/envconfig"
)

var build = "develop"

type Config struct {
	Db *DB
}

type DB struct {
	User         string `default:"postgres"`
	Password     string `default:"postgres,mask"`
	Host         string `default:"localhost"`
	Name         string `default:"postgres"`
	MaxIdleConns int    `default:"0"`
	MaxOpenConns int    `default:"0"`
	DisableTLS   bool   `default:"true"`
}

// NewConfig constructs a Config struct which represents server settings, and populates it with default values.
func NewConfig() (*Config, error) {
	cfg := Config{
		Db: &DB{},
	}
	err := evconfig.Process(build, &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &cfg, nil
}
