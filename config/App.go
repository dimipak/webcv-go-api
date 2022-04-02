package config

import (
	"app/helpers"
	"os"
)

type App struct {
	Port string `env:"APP_PORT"`
}

func (a *App) setValues() {
	decoder := helpers.Decoder{
		Interface: a,
		GetValue:  os.Getenv,
	}
	decoder.Decode("env")
}
