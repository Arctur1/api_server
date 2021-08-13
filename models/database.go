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

	host, port, dbName, user, pass := os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, pass)

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

func ClearTable() {
	DB.Exec("DELETE FROM coins")
	DB.Exec("ALTER SEQUENCE coins_id_seq RESTART WITH 1")
}
