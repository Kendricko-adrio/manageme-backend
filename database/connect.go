package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

func connect() (*gorm.DB, error) {
	// dsn :=

	dbUrl := "postgres://ywlthwdairldka:d6df6526e42e93a9603a8ee4c66cb255f19808019207a4cf1aef69125e830d42@ec2-44-193-111-218.compute-1.amazonaws.com:5432/d3ja58okmu16jf"
	// dbUrl := ""
	if dbUrl != "" {
		return gorm.Open(postgres.Open(dbUrl), &gorm.Config{
			PrepareStmt: true,
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					Colorful:      true,
					LogLevel:      logger.Info,
				},
			),
		})
	} else {
		dsn := "host=127.0.0.1 user=postgres password=asdfasw22 dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	// return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func GetInstance() *gorm.DB {

	if database == nil {
		database, _ = connect()

	}
	fmt.Println("masuk")
	return database
}
