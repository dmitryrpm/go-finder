package config

import (
	"errors"
	"flag"
	"fmt"
)

const (
	Kmax         = 1000
	DefaultCount = 5
	DefaultType  = "url"
)

var (
	ValidateTypeError = errors.New("incorrect Type param (url or file)")
	ValidateKError    = fmt.Errorf("incorrect K param (1..%d)", Kmax)
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
		DefaultCount,
		"limit K goroutines",
	)

	flag.StringVar(
		&cfg.Type,
		"type",
		DefaultType,
		"set type",
	)

	flag.Parse()

	return cfg
}

func ValidateConfig(cfg *Config) error {
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
