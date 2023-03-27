package currency_service

import "github.com/impactify-api/internal/src/models"

// CurrencyService is a serivce that retrieves rates for difference providers
type CurrencyService struct {
	CurrencyProviderList map[string]IProvider
}

// ICurrencyService is an interface that for currency service
type ICurrencyService interface {
	AddToCurrencyProviders(provider IProvider)
	ReturnRate(providerName string, symbol string) *models.Currency
	GetProviderMapping() map[string]IProvider
}

// IProvider is an interface for currency providers
type IProvider interface {
	GetProvider() string
	GetRate(symbol string) *models.Currency
}

// NewCurrencyService creates new currecy service
func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		CurrencyProviderList: make(map[string]IProvider),
	}
}

// AddToCurrencyProviders adds to struct containing all providers
func (c *CurrencyService) AddToCurrencyProviders(provider IProvider) {
	c.CurrencyProviderList[provider.GetProvider()] = provider
}

func (c *CurrencyService) GetProviderMapping() map[string]IProvider {
	return c.CurrencyProviderList
}

// ReturnRate returns the rate for a given provider
func (c *CurrencyService) ReturnRate(providerName string, symbol string) *models.Currency {
	return c.CurrencyProviderList[providerName].GetRate(symbol)
}
