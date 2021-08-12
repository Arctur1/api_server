package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	host, port, dbName, user, pass, ssl := os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbName, pass, ssl)

	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DBURL)
	}

	db.DB().SetMaxIdleConns(10)
	DB = db
	db.Debug().AutoMigrate(&Coin{})
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
