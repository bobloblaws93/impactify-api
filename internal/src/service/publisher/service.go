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
	RetrievePublisherInformation(id string, startDate time.Time, endDate time.Time) (*models.PublisherInformation, error)
	RetrieveAllPublisherInformation(startDate time.Time, endDate time.Time) ([]*models.PublisherInformation, error)
	RetrievePublisher(id string) (*models.Publisher, error)
	RetrieveAllPublishers() ([]models.Publisher, error)
	RetrieveAllPublisherRows(id string, startDate time.Time, endDate time.Time) ([]*models.PublisherInformation, error)
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

// RetrievePublisherInformation is a service function that retrieves an aggregate of a publisher's information  by id
func (s *PublisherService) RetrievePublisherInformation(id string, startDate time.Time, endDate time.Time) (*models.PublisherInformation, error) {
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

// RetrieveAllPublisherRows is a service function that retrieves all rows by a publisher id
func (s *PublisherService) RetrieveAllPublisherRows(id string, startDate time.Time, endDate time.Time) ([]*models.PublisherInformation, error) {
	pubInfoRows, err := s.PublisherRepo.GetAllPublisherInformationRows(id, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return pubInfoRows, nil
}

// RetrieveAllPublishers is a service function that retrieves all publishers
func (s *PublisherService) RetrieveAllPublishers() ([]models.Publisher, error) {
	allPubs, err := s.PublisherRepo.GetPublishers()
	if err != nil {
		return nil, err
	}
	return allPubs, nil
}
