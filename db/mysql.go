package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_USER = "root"
	DB_PASS = "root"
	DB_NAME = "nest_academy_final_project"
)

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
