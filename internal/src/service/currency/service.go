package currency_service

import "github.com/impactify-api/internal/src/models"

type CurrencyService struct {
	CurrencyProviderList map[string]IProvider
}

type ICurrencyService interface {
	AddToCurrencyProviders(provider IProvider)
	ReturnRate(providerName string, symbol string) *models.Currency
}

type IProvider interface {
	GetProvider() string
	GetRate(symbol string) *models.Currency
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		CurrencyProviderList: make(map[string]IProvider),
	}
}

func (c *CurrencyService) AddToCurrencyProviders(provider IProvider) {
	c.CurrencyProviderList[provider.GetProvider()] = provider
}

func (c *CurrencyService) ReturnRate(providerName string, symbol string) *models.Currency {
	return c.CurrencyProviderList[providerName].GetRate(symbol)
}
