package main

import (
	"github.com/arctur1/api_server/controllers"
	"github.com/arctur1/api_server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Coin{})
}

func main() {
	db := models.Init()
	Migrate(db)
	defer db.Close()
	setupServer().Run(":8080")
}

func setupServer() *gin.Engine {

	route := gin.Default()

	models.Init()

	route.GET("/coins", controllers.GetAllCoins)
	route.POST("/coins", controllers.CreateCoin)
	route.GET("/coins/:id", controllers.GetCoin)
	route.PATCH("/coins/:id", controllers.UpdateCoin)
	route.DELETE("/coins/:id", controllers.DeleteCoin)

	return route
}
