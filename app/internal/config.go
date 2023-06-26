package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	isDebug bool `env:"IS_DEBUG"	env-default:"false"`
	isDev   bool `env:"IS_DEV"		env-default:"false"`
	Listen  struct {
		Type   string `env:"LISTEN_TYPE"	env-default:"port"`
		BindIp string `env:"BIND_IP"		env-default:"0.0.0.0"`
		Port   string `env:"PORT"			env-default:"10000"`
	}
	AppConfig struct {
		LogLevel struct {
			Email    string `env:"ADMIN_EMAIL"		env-required:"true"`
			Password string `env:"ADMIN_PASSWORD"	env-required:"true"`
		}
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func () {
		log.Print("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Golang tasks"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}
