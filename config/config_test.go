package config

import (
	"testing"
)

func TestEmptyConfig(t *testing.T) {
	cfg, _ := NewConfig()
	if cfg.K != DefaultCount {
		t.Fatal("Incorrect default K")
	}
	if cfg.Type != DefaultType {
		t.Fatal("Incorrect default TYPE")
	}
}

func TestValidateConfig(t *testing.T) {
	cfg := Config{
		K:    10,
		Type: "url",
	}
	err := ValidateConfig(&cfg)
	if err != nil {
		t.Fatal("Incorrect validate success config")
	}
	cfg.K = -1
	err = ValidateConfig(&cfg)
	if err != ValidateKError {
		t.Fatal("Incorrect validate config")
	}
	cfg.K = 10
	cfg.Type = "trololo"
	err = ValidateConfig(&cfg)
	if err != ValidateTypeError {
		t.Fatal("Incorrect validate config")
	}
}
