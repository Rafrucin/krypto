package handlers

import (
	"net/http"
	"github.com/Rafrucin/krypto/services"
	"github.com/gin-gonic/gin"
)

func GetRates(c *gin.Context) {
	currencies := c.Query("currencies")

	if len(currencies) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	rates, err := services.FetchRates(currencies)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, rates)
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome!")
}