package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rafrucin/krypto/services"
	"github.com/gin-gonic/gin"
)

func ExchangeCrypto(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	amountStr := c.Query("amount")

	if from == "" || to == "" || amountStr == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	result, err := services.ExchangeCrypto(from, to, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, result)
}
