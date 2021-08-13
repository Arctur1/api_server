package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arctur1/api_server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_GetCoins(t *testing.T) {

	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	models.Init()
	db := models.GetDB()
	req, w := setGetCoinsRouter(db)
	defer db.Close()

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Coin{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.Coin{}
	a.Equal(expected, actual)
	models.ClearTable()
}

func setGetCoinsRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	api := &APIEnv{DB: db}
	r.GET("/", api.GetAllCoins)
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
}
