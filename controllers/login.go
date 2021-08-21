package controllers

import (
	"fmt"
	"net/http"

	"github.com/arctur1/api_server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func (a *APIEnv) Signup(c *gin.Context) {
	creds := &Credentials{}
	err := c.BindJSON(&creds)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	hashedString := string(hashedPassword)
	user := models.User{Username: creds.Username, Password: hashedString}
	if err := a.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Signed up")

}
