package config

import (
	"time"

	"gabe565.com/webos-dev-mode/pkg/webosdev"
)

type Config struct {
	BaseURL        string
	Insecure       bool
	Token          string
	RequestTimeout time.Duration
	CronInterval   time.Duration
	JSON           bool
}

func New() *Config {
	return &Config{
		BaseURL:        webosdev.DefaultBaseURL,
		RequestTimeout: time.Minute,
		CronInterval:   24 * time.Hour,
	}
}
