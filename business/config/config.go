package config

import (
	"errors"
	"fmt"

	"github.com/ardanlabs/conf/v2"
)

var build = "develop"

type Config struct {
	conf.Version
	Socket  *SocketConfig  `yaml:"socket" json:"socket" usage:"Socket configuration."`
	Session *SessionConfig `yaml:"session" json:"session" usage:"Session authentication settings."`
	Db      *DB
}

// SocketConfig is configuration relevant to the transport socket and protocol.
type SocketConfig struct {
}

// SessionConfig is configuration relevant to the session.
type SessionConfig struct {
	SingleSocket bool `yaml:"single_socket" json:"single_socket" usage:"Only allow one socket per user. Older sessions are disconnected. Default false."`
}

type DB struct {
	User         string `conf:"default:postgres"`
	Password     string `conf:"default:postgres,mask"`
	Host         string `conf:"default:localhost"`
	Name         string `conf:"default:postgres"`
	MaxIdleConns int    `conf:"default:0"`
	MaxOpenConns int    `conf:"default:0"`
	DisableTLS   bool   `conf:"default:true"`
}

func NewSocketConfig() *SocketConfig {
	return &SocketConfig{}
}

func NewSessionConfig() *SessionConfig {
	return &SessionConfig{}
}

func NewDBConfig() (*DB, error) {
	db := DB{}
	const prefix = "DbConfig"
	help, err := conf.Parse(prefix, &db)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil, err
		}
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	return &db, nil
}

// NewConfig constructs a Config struct which represents server settings, and populates it with default values.
func NewConfig() (*Config, error) {
	dbconfig, err := NewDBConfig()
	if err != nil {
		return nil, err
	}
	config := Config{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
		Db: dbconfig,
	}
	return &config, nil
}

func (c *Config) GetSocket() *SocketConfig {
	return c.Socket
}

func (c *Config) GetSession() *SessionConfig {
	return c.Session
}

func (c *Config) GetDb() *DB {
	return c.Db
}
