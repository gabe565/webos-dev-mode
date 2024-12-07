package config

import "time"

type Config struct {
	Token          string
	RequestTimeout time.Duration
	CronInterval   time.Duration
}

func New() *Config {
	return &Config{
		RequestTimeout: time.Minute,
		CronInterval:   24 * time.Hour,
	}
}
