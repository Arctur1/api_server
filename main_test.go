package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinsRoute(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/coins", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"coins\":[]}", w.Body.String())
}
