package database

import (
	_ "github.com/go-sql-driver/mysql"
)

// func ConnectDB() *gorm.DB {

// 	var dbc config.Database

// 	config.GetC(&dbc)

// 	// database := config.Get().Database.Name
// 	// dbUser := config.Get().Database.Username
// 	// dbPass := config.Get().Database.Password
// 	// dbHost := config.Get().Database.Host

// 	database := dbc.Name
// 	dbUser := dbc.Username
// 	dbPass := dbc.Password
// 	dbHost := dbc.Host

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, database)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return db
// }
