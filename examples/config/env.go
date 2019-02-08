package config

import (
	"errors"
	"time"
	"webservice-library/config"
)

// Env is a collection of environment variable
type Env struct {
	config.Config
	DbEngine      string         `json:"db_engine"`
	MaxConnection int            `json:"max_connection"`
	Production    bool           `json:"production"`
	TimeZone      *time.Location `json:"time_zone"`
}

// GetPrefix returns the current config prefix value
func (env Env) GetPrefix() string {
	return env.Prefix
}

// SetPrefix sets the config prefix value
func (env *Env) SetPrefix(prefix string) error {
	if env == nil {
		return errors.New("The Config object is not initialized")
	}
	env.Prefix = prefix
	return nil
}

// GetPath returns the path of the .ini file
func (env Env) GetPath() string {
	return env.Path
}

// SetPath sets the path of the .ini file
func (env *Env) SetPath(path string) error {
	if env == nil {
		return errors.New("The Config object is not initialized")
	}
	if path == "" {
		env.Path = "env.ini"
		return nil
	}
	env.Path = path
	return nil
}
