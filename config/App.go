package config

import "os"

type App struct {
	Port string `env:"APP_PORT"`
}

func (a *App) setValues() {
	enc := Encode{
		Struct:   a,
		GetValue: os.Getenv,
	}
	enc.EnvEncode2("env")
}
