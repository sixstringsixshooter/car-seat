package config

type Config struct {
	Versions Versions `json:"versions"`
}

type Versions map[string]bool
