package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_HOST = getEnvVariable("DB_HOST")
	DB_PORT = getEnvVariable("DB_PORT")
	DB_USER = getEnvVariable("DB_USER")
	DB_PASS = getEnvVariable("DB_PASS")
	DB_NAME = getEnvVariable("DB_NAME")
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME,
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return gormDB, nil
}
