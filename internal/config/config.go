package config

import "time"

type Config struct {
	Token        string
	CronInterval time.Duration
}

func New() *Config {
	return &Config{
		CronInterval: 24 * time.Hour,
	}
}
