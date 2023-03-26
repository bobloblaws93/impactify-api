package publisher_service

import (
	"time"

	"github.com/impactify-api/internal/src/models"
)

// PublisherMock is a mock that implements IPublisherService. Used for testing
type PublisherMock struct{}

func NewPublisherMock() *PublisherMock {
	return &PublisherMock{}
}

func (p *PublisherMock) RetrievePublisher(id string) (*models.Publisher, error) {
	return &models.Publisher{
		ID:   123,
		Name: "test",
	}, nil
}

func (p *PublisherMock) RetrievePublisherRevenue(id string, startDate time.Time, endDate time.Time) (*models.PublisherInformation, error) {
	return &models.PublisherInformation{
		Publisher: models.Publisher{
			ID:   123,
			Name: "test",
		},
		Impressions: 100,
		Requests:    100,
		Clicks:      100,
		Revenue:     100,
	}, nil
}

func (p *PublisherMock) RetrieveAllPublisherInformation(startDate time.Time, endDate time.Time) ([]*models.PublisherInformation, error) {
	return []*models.PublisherInformation{
		{
			Publisher: models.Publisher{
				ID:   123,
				Name: "test",
			},
			Impressions: 100,
			Requests:    100,
			Clicks:      100,
			Revenue:     100,
		},
	}, nil
}
func (p *PublisherMock) RetrieveAllPublishers() ([]models.Publisher, error) {
	return []models.Publisher{
		{
			ID:   123,
			Name: "test",
		},
	}, nil
}
