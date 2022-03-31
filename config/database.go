package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Name      string `env:"MYSQL_DB"`
	Username  string `env:"MYSQL_USER"`
	Password  string `env:"MYSQL_PASS"`
	Host      string `env:"MYSQL_HOST"`
	JWTSecret string `env:"JWT_SECRET"`
}

func (d *Database) setValues() {
	envEncode(d)
}

func GORM() *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		G_DATABASE.Username,
		G_DATABASE.Password,
		G_DATABASE.Host,
		G_DATABASE.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
