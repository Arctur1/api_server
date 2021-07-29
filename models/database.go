package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=rozhn dbname=coins password=rozhn sslmode=disable")
	if err != nil {
		panic("Can't connect to database")
	}

	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=rozhn dbname=coins_test password=rozhn sslmode=disable")
	if err != nil {
		panic("Can't connect to database")
	}

	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../gorm_test.db")
	return err
}

func GetDB() *gorm.DB {
	return DB
}
