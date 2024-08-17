package helpers

import (
	"log"
	"os"

	"github.com/Rafrucin/krypto/services"
)

func SetApiKey(apikey string) {
	if services.ApiKey != "" {
		return
	}
	envAPIKEY := os.Getenv("APIKEY")
	if apikey != "" {
		services.ApiKey = apikey
	} else if envAPIKEY != "" {
		services.ApiKey = envAPIKEY
	} else if !readApiKeyFromFile() {
		log.Fatal("open exchange api key required to run the app!\n",
			"please provide apikey as the app parameter \"apikey\"\n",
			"or as environment variable APIKEY")
	}
}

func readApiKeyFromFile() bool {
	data, err := os.ReadFile("apikey.txt")
	if err != nil || len(data) == 0 {
		data, err = os.ReadFile("../apikey.txt")
		if err != nil || len(data) == 0 {
			return false
		}
	}
	services.ApiKey = string(data)
	return true
}
