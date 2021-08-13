package controllers

import (
	"fmt"
	"net/http"

	"github.com/arctur1/api_server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateCoinInput struct {
	Coin    string `json:"coin" binding:"required"`
	Balance int    `json:"balance" binding:"required"`
}

type UpdateCoinInput struct {
	Coin    string `json:"coin"`
	Balance int    `json:"balance"`
}

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) CreateCoin(c *gin.Context) {
	coin := models.Coin{}
	err := c.BindJSON(&coin)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.Create(&coin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, coin)
}

func (a *APIEnv) GetAllCoins(context *gin.Context) {
	coins, err := GetCoins(a.DB)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{"coins": coins})
}

func (a *APIEnv) GetCoin(context *gin.Context) {
	id := context.Params.ByName("id")
	coin, exists, err := GetCoinByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		context.JSON(http.StatusNotFound, "there is no coin in db")
		return
	}

	context.JSON(http.StatusOK, coin)
}

func (a *APIEnv) DeleteCoin(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := GetCoinByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	err = DeleteCoin(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
}

func (a *APIEnv) UpdateCoin(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := GetCoinByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	updatedCoin := models.Coin{}
	err = c.BindJSON(&updatedCoin)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := UpdateCoin(a.DB, &updatedCoin); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	a.GetCoin(c)
}
