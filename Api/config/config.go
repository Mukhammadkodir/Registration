package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Enivorentment   string
	RegisterServiceHost string
	RegisterServicePort int

	CtxTimeout      int

	LogLevel string
	HTTPPort string

	
	SigninKey string
}

func Load() Config {
	c := Config{}

	c.Enivorentment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	
	c.RegisterServiceHost = cast.ToString(getOrReturnDefault("TASK_SERVICE_HOST", "localhost"))
	c.RegisterServicePort = cast.ToInt(getOrReturnDefault("TASK_SERVICE_PORT", 9191))

	c.SigninKey = cast.ToString(getOrReturnDefault("SIGNING_KEY", "najottalimsecretkey"))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
