package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDBConnection() (*gorm.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_SERVER_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
