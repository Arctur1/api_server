package controllers

import (
	"net/http"

	"github.com/arctur1/api_server/models"
	"github.com/gin-gonic/gin"
)

type CreateCoinInput struct {
	Coin    string `json:"coin" binding:"required"`
	Balance int    `json:"balance" binding:"required"`
}

type UpdateCoinInput struct {
	Coin    string `json:"coin"`
	Balance int    `json:"balance"`
}

func GetAllCoins(context *gin.Context) {
	var coins []models.Coin

	models.DB.Find(&coins)

	context.JSON(http.StatusOK, gin.H{"coins": coins})
}

func GetCoin(context *gin.Context) {
	var coin models.Coin
	if err := models.DB.Where("id = ?", context.Param("id")).First(&coin).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"coins": coin})
}

func CreateCoin(context *gin.Context) {
	var input CreateCoinInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coin := models.Coin{Coin: input.Coin, Balance: input.Balance}
	models.DB.Create(&coin)

	context.JSON(http.StatusOK, gin.H{"coins": coin})
}

func UpdateCoin(context *gin.Context) {

	var coin models.Coin
	if err := models.DB.Where("id = ?", context.Param("id")).First(&coin).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
		return
	}

	var input UpdateCoinInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&coin).Update(input)

	context.JSON(http.StatusOK, gin.H{"coins": coin})
}

func DeleteCoin(context *gin.Context) {

	var coin models.Coin
	if err := models.DB.Where("id = ?", context.Param("id")).First(&coin).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
		return
	}

	models.DB.Delete(&coin)

	context.JSON(http.StatusOK, gin.H{"coins": true})
}