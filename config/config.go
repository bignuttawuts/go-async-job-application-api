package config

import "time"

type Config struct {
	Port         string
	Prefork      bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
