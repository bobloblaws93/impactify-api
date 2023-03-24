package service

import (
	"database/sql"

	"time"

	"github.com/impactify-api/internal/src/models"
	"github.com/impactify-api/internal/src/repository/currency"
	"github.com/impactify-api/internal/src/repository/publishers"
)

type Service struct {
	PublisherRepo *publishers.Repository
	CurrencyRepo  *currency.Repository
}

type IService interface {
	RetrievePublisherRevenue(id, symbol string, startDate time.Time, endDate time.Time) *models.PublisherInformation
	RetrievePublisher(id string) *models.Publisher
	RetrieveAllPublishers() []models.Publisher
	RetrieveCurrency(symbol string) *models.Currency
}

func NewService(client *sql.DB) *Service {
	pubRepo := publishers.NewRepository(client)
	currencyRepo := currency.NewRepository(client)
	return &Service{
		PublisherRepo: pubRepo,
		CurrencyRepo:  currencyRepo,
	}
}

func (s *Service) RetrievePublisher(id string) *models.Publisher {
	return s.PublisherRepo.GetPublisher(id)
}

func (s *Service) RetrievePublisherRevenue(id, symbol string, startDate time.Time, endDate time.Time) *models.PublisherInformation {
	pubInfo := s.PublisherRepo.GetPublisherInformationByID(id, startDate, endDate)
	currencyInfo := s.CurrencyRepo.GetCurrency(symbol)
	pubInfo.Revenue = pubInfo.Revenue * currencyInfo.Rate
	return pubInfo
}

func (s *Service) RetrieveAllPublishers() []models.Publisher {
	return s.PublisherRepo.GetPublishers()
}

func (s *Service) RetrieveCurrency(symbol string) *models.Currency {
	return s.CurrencyRepo.GetCurrency(symbol)
}
