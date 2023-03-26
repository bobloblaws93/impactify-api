package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/impactify-api/internal/src/config"
	"github.com/impactify-api/internal/src/models"
)

type ExchangeRateProvider struct {
	ProviderName string
	config       config.Config
}

func NewExchangeRateProvider(config config.Config) *ExchangeRateProvider {
	return &ExchangeRateProvider{
		ProviderName: "ExchangeRateProviderAPI",
		config:       config,
	}
}

func (erp *ExchangeRateProvider) GetProvider() string {
	return erp.ProviderName
}

func (erp *ExchangeRateProvider) GetRate(symbol string) *models.Currency {
	// call fixer api to get the rate
	// return the rate
	client := &http.Client{}
	requestURL := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/USD/%s/1", erp.config.API.ExchangeRateKey, symbol)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return &models.Currency{}
	}
	var data map[string]interface{}
	response, err := client.Do(req)
	if err != nil {
		return &models.Currency{}
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &models.Currency{}
	}
	// parse out response body
	err = json.Unmarshal([]byte(string(responseBody)), &data)
	if err != nil {
		return &models.Currency{}
	}
	if data["result"] == "success" {
		return &models.Currency{
			Symbol: symbol,
			Rate:   data["conversion_rate"].(float64),
		}
	}

	return &models.Currency{}
}
