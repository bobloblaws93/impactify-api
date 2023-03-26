package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/impactify-api/internal/src/config"
	"github.com/impactify-api/internal/src/models"
)

type Fixer struct {
	ProviderName string
	config       config.Config
}

func NewFixerProvider(config config.Config) *Fixer {
	return &Fixer{
		ProviderName: "fixer",
		config:       config,
	}
}

func (f *Fixer) GetProvider() string {
	return f.ProviderName
}

func (f *Fixer) GetRate(symbol string) *models.Currency {
	// call fixer api to get the rate
	// return the rate
	client := &http.Client{}
	requestURL := fmt.Sprintf("https://api.apilayer.com/fixer/convert?to=%s&from=USD&amount=1", symbol)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return &models.Currency{}
	}
	var data map[string]interface{}

	req.Header.Set("apikey", f.config.API.FixerKey)
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
	if data["success"] == true {
		return &models.Currency{
			Symbol: symbol,
			Rate:   data["result"].(float64),
		}
	}

	return &models.Currency{}
}
