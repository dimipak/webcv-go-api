package config

type server struct {
	Port string `env:"SERVER_PORT"`
}

func Server() server {

	var sv server

	EnvEncode(&sv)

	return sv
}
