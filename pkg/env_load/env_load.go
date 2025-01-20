package envload

import (
	"fmt"

	"github.com/joeshaw/envdecode"
)

var cfg Config

type Config struct {
	Environment string `env:"ENV,required"`
	DB          struct {
		Host     string `env:"DB_HOST,required"`
		Port     string `env:"DB_PORT,required"`
		User     string `env:"DB_USER,required"`
		Password string `env:"DB_PASSWORD,required"`
		DBName   string `env:"DB_NAME,required"`
	}
}

// Get returns a config structure.
func Get() Config {
	return cfg
}

func Init() {
	if err := envdecode.Decode(&cfg); err != nil {
		panic(fmt.Sprintf("error to decode config: %s", err))
	}
}
