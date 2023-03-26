package currency_service

import (
	"github.com/impactify-api/internal/src/models"
)

// CurrencyMock is a mock that implements ICurrencyService. Used for testing
type CurrencyMock struct {
	CurrencyProviderList map[string]IProvider
}

func NewCurrencyMock() *CurrencyMock {
	return &CurrencyMock{
		CurrencyProviderList: make(map[string]IProvider),
	}
}

func (c *CurrencyMock) GetProvider() string {
	return "test"
}

func (c *CurrencyMock) AddToCurrencyProviders(provider IProvider) {
	c.CurrencyProviderList[provider.GetProvider()] = provider
}

func (c *CurrencyMock) ReturnRate(providerName string, symbol string) *models.Currency {
	return &models.Currency{
		Symbol: symbol,
		Rate:   0.90,
	}
}
