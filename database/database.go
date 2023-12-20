package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database = func() (db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		database = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("connection error")
		panic(err)
	} else {
		log.Println("connected to mysql")
		return db
	}
}()
