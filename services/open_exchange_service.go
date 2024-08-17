package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ApiKey = ""

const apiUrl = "https://openexchangerates.org/api/latest.json"

type RateResponse struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

type APIResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchRates(currencies string) ([]RateResponse, error) {
	url := fmt.Sprintf("%s?app_id=%s", apiUrl, ApiKey)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch rates")
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, errors.New("failed to parse response")
	}

	currencyList := strings.Split(currencies, ",")
	if len(currencyList) < 2 {
		return nil, errors.New("invalid currencies parameter")
	}

	var rates []RateResponse
	for i := 0; i < len(currencyList); i++ {
		if apiResp.Rates[currencyList[i]] == 0 {
			continue
		}
		for j := i + 1; j < len(currencyList); j++ {

			if apiResp.Rates[currencyList[j]] == 0 {
				continue
			}
			rates = append(rates, RateResponse{
				From: currencyList[i],
				To:   currencyList[j],
				Rate: apiResp.Rates[currencyList[j]] / apiResp.Rates[currencyList[i]],
			})
			rates = append(rates, RateResponse{
				From: currencyList[j],
				To:   currencyList[i],
				Rate: apiResp.Rates[currencyList[i]] / apiResp.Rates[currencyList[j]],
			})
		}
	}
	return rates, nil
}



