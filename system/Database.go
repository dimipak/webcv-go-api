package system

import (
	"app/config"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var sqlDB *sql.DB

func SetupDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.G_DATABASE.Username,
		config.G_DATABASE.Password,
		config.G_DATABASE.Host,
		config.G_DATABASE.Name,
	)

	var err error
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
}

func GORM() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		panic(err.Error())
	}

	return db
}
