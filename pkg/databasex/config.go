package databasex

import "time"

type Configs map[string]*Config

type Config struct {
	DriverName string        `mapstructure:"driver_name" json:"driver_name" yaml:"driver_name"`
	DSN        string        `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	Timeout    time.Duration `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
