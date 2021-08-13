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
	route.POST("/coins", api.CreateCoin)
	route.GET("/coins/:id", api.GetCoin)
	route.PATCH("/coins/:id", api.UpdateCoin)
	route.DELETE("/coins/:id", api.DeleteCoin)

	return route
}
