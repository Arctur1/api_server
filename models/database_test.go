package models

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	asserts := assert.New(t)
	db := Init()
	asserts.NoError(db.DB().Ping(), "Db should be able to ping")

	// Test get a connecting from connection pools
	connection := GetDB()
	asserts.NoError(connection.DB().Ping(), "Db should be able to ping")
	db.Close()

}

func refreshCoinTable() error {
	connection := GetDB()
	err := connection.DropTableIfExists(&Coin{}).Error
	if err != nil {
		return err
	}
	err = connection.AutoMigrate(&Coin{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneCoin() error {

	refreshCoinTable()

	coin := Coin{
		Coin:    "Bitcoin",
		Balance: 1,
	}

	err := DB.Model(&Coin{}).Create(&coin).Error
	if err != nil {
		log.Fatalf("cannot seed coins table: %v", err)
	}
	return nil
}

func TestFindAllCoins(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	db := Init()
	var coins []Coin

	err := seedOneCoin()
	if err != nil {
		log.Fatal(err)
	}
	db.Find(&coins)
	assert.Equal(t, len(coins), 1)

	err = refreshCoinTable()
	if err != nil {
		log.Fatal(err)
	}

}
