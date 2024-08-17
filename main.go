package main

import (
	"flag"

	"github.com/Rafrucin/krypto/handlers"
	"github.com/Rafrucin/krypto/helpers"
	"github.com/gin-gonic/gin"
)

func main() {

	apikeyflag := flag.String("apikey", "", "open exchange api key")
	release := flag.Bool("release", false, "release mode")
	flag.Parse()

	helpers.SetApiKey(*apikeyflag)

	if *release {
		println("running in release mode")
		gin.SetMode(gin.ReleaseMode)
	}

	helpers.SetApiKey("")

	router := gin.Default()
	router.GET("/", handlers.Welcome)
	router.GET("/rates", handlers.GetRates)
	router.GET("/exchange", handlers.ExchangeCrypto)
	router.Run(":8080")
}

