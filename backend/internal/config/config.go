// backend/internal/config/config.go
package config

import "os"

type Config struct {
	Port        string
	AllowOrigin string
	OpenAIKey   string
}

func Load() Config {
	get := func(k, def string) string {
		if v := os.Getenv(k); v != "" {
			return v
		}
		return def
	}
	return Config{
		Port:        get("PORT", "8080"),
		AllowOrigin: get("ALLOW_ORIGIN", "*"),
		OpenAIKey:   get("OPENAI_API_KEY", ""),
	}
}
