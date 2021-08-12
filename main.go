package main

import (
	"github.com/arctur1/api_server/controllers"
	"github.com/arctur1/api_server/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.Init()
	setupServer().Run(":8080")
}

func setupServer() *gin.Engine {

	route := gin.Default()

	api := &controllers.APIEnv{
		DB: models.GetDB(),
	}

	route.GET("/coins", api.GetAllCoins)
	route.POST("/coins", controllers.CreateCoin)
	route.GET("/coins/:id", controllers.GetCoin)
	route.PATCH("/coins/:id", controllers.UpdateCoin)
	route.DELETE("/coins/:id", controllers.DeleteCoin)

	return route
}
