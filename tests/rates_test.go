package tests

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rafrucin/krypto/handlers"
	"github.com/Rafrucin/krypto/helpers"
	"github.com/Rafrucin/krypto/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var apikeyflag = flag.String("apikey", "", "open exchange api key")

func TestGetRates_Success(t *testing.T) {
	helpers.SetApiKey(*apikeyflag)
	router := gin.Default()
	router.GET("/rates", handlers.GetRates)
	helpers.SetApiKey(*apikeyflag)
	println(services.ApiKey)
	req, _ := http.NewRequest("GET", "/rates?currencies=USD,GBP,EUR", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "rate")
}

func TestGetRates_NonExistingCurrency(t *testing.T) {
	router := gin.Default()
	router.GET("/rates", handlers.GetRates)
	req, _ := http.NewRequest("GET", "/rates?currencies=USD,GBP,ZZZ", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "rate")
	req, _ = http.NewRequest("GET", "/rates?currencies=USD,GBP,EUR", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "rate")
}

func TestGetRates_BadRequest(t *testing.T) {
	router := gin.Default()
	router.GET("/rates", handlers.GetRates)

	req, _ := http.NewRequest("GET", "/rates?currencies=USD", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}