package test

import (
	"log"
	"testing"

	"github.com/arctur1/api_server/models"
	"github.com/stretchr/testify/assert"
)

func TestConnectingDatabase(t *testing.T) {
	asserts := assert.New(t)
	db := models.Init()
	// Test create & close DB
	asserts.NoError(db.DB().Ping(), "Db should be able to ping")

	// Test get a connecting from connection pools
	connection := models.GetDB()
	asserts.NoError(connection.DB().Ping(), "Db should be able to ping")
	db.Close()

}

func TestConnectingTestDatabase(t *testing.T) {
	asserts := assert.New(t)
	db := models.TestDBInit()

	// Test create & close DB
	asserts.NoError(db.DB().Ping(), "Db should be able to ping")

	// Test get a connecting from connection pools
	connection := models.GetDB()
	asserts.NoError(connection.DB().Ping(), "Db should be able to ping")
	db.Close()

}

func refreshCoinTable() error {
	err := models.DB.DropTableIfExists(&models.Coin{}).Error
	if err != nil {
		return err
	}
	err = models.DB.AutoMigrate(&models.Coin{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneCoin() error {

	refreshCoinTable()

	coin := models.Coin{
		Coin:    "Bitcoin",
		Balance: 1,
	}

	err := models.DB.Model(&models.Coin{}).Create(&coin).Error
	if err != nil {
		log.Fatalf("cannot seed coins table: %v", err)
	}
	return nil
}

func TestFindAllCoins(t *testing.T) {
	db := models.TestDBInit()
	var coins []models.Coin

	err := refreshCoinTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedOneCoin()
	if err != nil {
		log.Fatal(err)
	}

	db.Find(&coins)

	assert.Equal(t, len(coins), 1)
}
