package config

type database struct {
	Name      string `env:"MYSQL_DB"`
	Username  string `env:"MYSQL_USER"`
	Password  string `env:"MYSQL_PASS"`
	Host      string `env:"MYSQL_HOST"`
	JWTSecret string `env:"JWT_SECRET"`
}

func Database() database {

	var db database

	EnvEncode(&db)

	return db
}
