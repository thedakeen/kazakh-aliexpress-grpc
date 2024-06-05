package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string        `env:"ENV" env-default:"development"`
	MongoURI       string        `env:"MONGO_URI" env-required:"true"`
	TokenTTL       time.Duration `env:"TOKEN_TTL" env-required:"true"`
	ServicePort    int           `env:"AUTH_SERVICE_PORT"`
	ServiceTimeout time.Duration `env:"AUTH_SERVICE_TIMEOUT"`
}

func MustLoad() *Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic("failed to read config " + err.Error())
	}

	return &cfg
}
