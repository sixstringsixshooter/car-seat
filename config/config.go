package config

type Config struct {
	Versions Versions `toml:"versions"`
}

type Versions map[string]bool
