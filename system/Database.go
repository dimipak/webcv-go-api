package system

import (
	"app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GORM() *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.G_DATABASE.Username,
		config.G_DATABASE.Password,
		config.G_DATABASE.Host,
		config.G_DATABASE.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
