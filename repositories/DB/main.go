package db

import (
	"app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db() *gorm.DB {

	database := config.Database().Name
	dbUser := config.Database().Username
	dbPass := config.Database().Password
	dbHost := config.Database().Host

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func New() *gorm.DB {
	return db()
}
