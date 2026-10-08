package config

import (
	"flag"
	"fmt"
	"os"
)

// Config holds all configuration for the service.
type Config struct {
	Addr   string
	NoAuth bool
}

// Default returns a Config with sensible defaults.
func Default() *Config {
	return &Config{
		Addr:   ":8470",
		NoAuth: false,
	}
}

// Load parses CLI flags and env vars, returning a Config.
// Precedence: defaults < env vars < flags.
func Load() *Config {
	cfg := Default()

	// Env vars
	if v := os.Getenv("LISTKIT_ADDR"); v != "" {
		cfg.Addr = v
	}
	if v := os.Getenv("LISTKIT_NO_AUTH"); v == "1" || v == "true" {
		cfg.NoAuth = true
	}

	// Flags
	flag.StringVar(&cfg.Addr, "addr", cfg.Addr, "listen address (env: LISTKIT_ADDR)")
	flag.BoolVar(&cfg.NoAuth, "no-auth", cfg.NoAuth, "disable auth for development (env: LISTKIT_NO_AUTH)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "listkit - agentic-first list and array manipulation service\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n  listkit [flags]\n\nFlags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	return cfg
}
