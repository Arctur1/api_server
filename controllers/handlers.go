package controllers

import (
	"github.com/arctur1/api_server/models"
	"github.com/jinzhu/gorm"
)

func GetCoins(db *gorm.DB) ([]models.Coin, error) {
	coins := []models.Coin{}
	query := db.Select("coin.*").
		Group("coin.id")
	if err := query.Find(&coins).Error; err != nil {
		return coins, err
	}

	return coins, nil
}

func GetBookByID(id string, db *gorm.DB) (models.Coin, bool, error) {
	coin := models.Coin{}

	query := db.Select("coin.*")
	query = query.Group("coin.id")
	err := query.Where("coin.id = ?", id).First(&coin).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return coin, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return coin, false, nil
	}
	return coin, true, nil
}

func DeleteBook(id string, db *gorm.DB) error {
	var b models.Coin
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, b *models.Coin) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
