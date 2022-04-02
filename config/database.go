package config

import (
	"app/helpers"
	"os"
)

type Database struct {
	Name      string `env:"MYSQL_DB"`
	Username  string `env:"MYSQL_USER"`
	Password  string `env:"MYSQL_PASS"`
	Host      string `env:"MYSQL_HOST"`
	JWTSecret string `env:"JWT_SECRET"`
}

func (d *Database) setValues() {
	decoder := helpers.Decoder{
		Interface: d,
		GetValue:  os.Getenv,
	}
	decoder.Decode("env")
}
