package services

import "errors"

func ExchangeCrypto(from string, to string, amount float64) (map[string]interface{}, error) {
	rates := map[string]float64{
		"BEER":  0.00002461,
		"FLOKI": 0.0001428,
		"GATE":  6.87,
		"USDT":  0.999,
		"WBTC":  57037.22,
	}

	fromRate, okFrom := rates[from]
	toRate, okTo := rates[to]

	

	if !okFrom || !okTo {
		return nil, errors.New("currency not supported")
	}

	result := amount * (fromRate / toRate)
	return map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": result,
	}, nil
}