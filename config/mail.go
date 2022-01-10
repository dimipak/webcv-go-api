package config

type mail struct {
	From     string `env:"MAIL_FROM"`
	Host     string `env:"MAIL_HOST"`
	Port     string `env:"MAIL_PORT"`
	Username string `env:"MAIL_USERNAME"`
	Password string `env:"MAIL_PASSWORD"`
}

func Mail() mail {

	var mail mail

	EnvEncode(&mail)

	return mail
}
