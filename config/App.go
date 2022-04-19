package config

import (
	"app/helpers"
	"os"
)

type App struct {
	URL  string `env:"APP_URL"`
	Port string `env:"APP_PORT"`
}

func (a *App) setValues() {
	decoder := helpers.Decoder{
		Interface: a,
		GetValue:  os.Getenv,
	}
	decoder.Decode("env")
}
