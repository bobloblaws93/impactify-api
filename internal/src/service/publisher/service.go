package publisher_service

import (
	"database/sql"

	"time"

	"github.com/impactify-api/internal/src/models"
	"github.com/impactify-api/internal/src/repository/publishers"
)

type PublisherService struct {
	PublisherRepo *publishers.Repository
}

type IPublisherService interface {
	RetrievePublisherRevenue(id string, startDate time.Time, endDate time.Time) (*models.PublisherInformation, error)
	RetrieveAllPublisherInformation(startDate time.Time, endDate time.Time) ([]*models.PublisherInformation, error)
	RetrievePublisher(id string) (*models.Publisher, error)
	RetrieveAllPublishers() ([]models.Publisher, error)
}

// create new publisher service
func NewService(client *sql.DB) *PublisherService {
	pubRepo := publishers.NewRepository(client)
	return &PublisherService{
		PublisherRepo: pubRepo,
	}
}

// RetrievePublisher is a service function that retrieves a publisher by id
func (s *PublisherService) RetrievePublisher(id string) (*models.Publisher, error) {
	publisher, err := s.PublisherRepo.GetPublisher(id)
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

// RetrievePublisherRevenue is a service function that retrieves a publisher's revenue by id
func (s *PublisherService) RetrievePublisherRevenue(id string, startDate time.Time, endDate time.Time) (*models.PublisherInformation, error) {
	pubInfo, err := s.PublisherRepo.GetPublisherInformationByID(id, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return pubInfo, nil
}

// RetrieveAllPublisherInformation is a service function that retrieves all publishers' revenue
func (s *PublisherService) RetrieveAllPublisherInformation(startDate time.Time, endDate time.Time) ([]*models.PublisherInformation, error) {
	allPubsInfo, err := s.PublisherRepo.GetAllPublisherInformation(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return allPubsInfo, nil
}

// RetrieveAllPublishers is a service function that retrieves all publishers
func (s *PublisherService) RetrieveAllPublishers() ([]models.Publisher, error) {
	allPubs, err := s.PublisherRepo.GetPublishers()
	if err != nil {
		return nil, err
	}
	return allPubs, nil
}
