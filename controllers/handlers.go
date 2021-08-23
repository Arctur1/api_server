package controllers

import (
	"github.com/arctur1/api_server/models"
	"github.com/jinzhu/gorm"
)

func GetCoins(db *gorm.DB) ([]models.Coin, error) {
	coins := []models.Coin{}
	query := db.Select("coins.*").
		Group("coins.id")
	if err := query.Find(&coins).Error; err != nil {
		return coins, err
	}

	return coins, nil
}

func GetCoinByID(id string, db *gorm.DB) (models.Coin, bool, error) {
	coin := models.Coin{}

	query := db.Select("coins.*")
	query = query.Group("coins.id")
	err := query.Where("coins.id = ?", id).First(&coin).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return coin, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return coin, false, nil
	}
	return coin, true, nil
}

func DeleteCoin(id string, db *gorm.DB) error {
	var c models.Coin
	if err := db.Where("id = ? ", id).Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCoin(db *gorm.DB, c *models.Coin) error {
	if err := db.Save(&c).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByName(username string, db *gorm.DB) (models.User, bool, error) {
	user := models.User{}

	query := db.Select("users.username")
	query = query.Group("users.username")
	//err := db.Where("username = ?", username).First(&user).Error
	err := db.Find(&user, "username = ?", username).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return user, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return user, false, nil
	}
	return user, true, nil
}
