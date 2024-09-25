package database

import (
	"fmt"
	"jwt-auth-test/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

var dsn string = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var DB *gorm.DB

func DBconnect() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	db.AutoMigrate(&models.User{})
}
