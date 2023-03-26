package providers

import "github.com/impactify-api/internal/src/models"

type ProviderMock struct{}

func NewProviderMock() *ProviderMock {
	return &ProviderMock{}
}

func (p *ProviderMock) GetProvider() string {
	return "test"
}

func (p *ProviderMock) GetRate(symbol string) *models.Currency {
	return &models.Currency{
		Symbol: symbol,
		Rate:   0.90,
	}
}
