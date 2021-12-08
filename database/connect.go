package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func connect() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=postgres password=asdfasw22 dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func GetInstance() *gorm.DB {

	if database == nil {
		database, _ = connect()

	}
	fmt.Println("masuk")
	return database
}
