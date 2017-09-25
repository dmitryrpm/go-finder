package config

import (
	"flag"
	"errors"
)

var (
	ValidateTypeError = errors.New("incorrect Type param (url or file)")
	ValidateKError = errors.New("incorrect K param (1..1000)")
)

const (
	DEFAULT_K = 5
	DEFAULT_TYPE = "url"
)

func NewConfig() (*Config, error) {
	cfg := ParseConfig()
	return cfg, ValidateConfig(cfg)
}


func ParseConfig() *Config {
	cfg := &Config{}

	flag.IntVar(
		&cfg.K,
		"k",
		DEFAULT_K,
		"limit K goroutines",
	)

	flag.StringVar(
		&cfg.Type,
		"type",
		DEFAULT_TYPE,
		"set type",
	)

	flag.Parse()

	return cfg
}

func ValidateConfig (cfg *Config) error {
	if cfg.Type != "url" && cfg.Type != "file" {
		return ValidateTypeError
	}

	if cfg.K <= 0 || cfg.K > 1000 {
		return ValidateKError
	}

	return nil
}


type Config struct {
	K    int
	Type string
}
