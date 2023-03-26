package publishers

import (
	"database/sql"
	"fmt"

	"time"

	"github.com/impactify-api/internal/src/models"
)

type Repository struct {
	dbClient *sql.DB
}

func NewRepository(client *sql.DB) *Repository {
	return &Repository{
		dbClient: client,
	}
}

func (r *Repository) GetDBClient() *sql.DB {
	return r.dbClient
}

// GetPublisher is a repository function that retrieves a publisher by id
func (r *Repository) GetPublisher(id string) (*models.Publisher, error) {
	// sql query to get a publisher by id
	var publisher models.Publisher
	err := r.dbClient.QueryRow("SELECT id, name FROM publishers WHERE id = ?", id).Scan(&publisher.ID, &publisher.Name)
	if err != nil {
		return &models.Publisher{}, err
	}

	return &publisher, nil
}

// GetPublisherInformationByID is a repository function that retrieves a publisher's information by id
func (r *Repository) GetPublisherInformationByID(id string, startDate, endDate time.Time) (*models.PublisherInformation, error) {
	var publisherInfo models.PublisherInformation
	query := fmt.Sprintf(`SELECT publisher_id, SUM(impressions), SUM(requests), SUM(clicks), SUM(revenue) FROM publishers_info WHERE publisher_id = %s AND date_created BETWEEN '%s' AND '%s' GROUP BY publisher_id`, id, startDate.String(), endDate.String())
	err := r.dbClient.QueryRow(query).
		Scan(&publisherInfo.Publisher.ID,
			&publisherInfo.Impressions,
			&publisherInfo.Requests,
			&publisherInfo.Clicks,
			&publisherInfo.Revenue)

	if err != nil {

		return &models.PublisherInformation{}, err
	}

	return &publisherInfo, nil
}

// GetAllPublisherInformation is a repository function that retrieves all publishers' information
func (r *Repository) GetAllPublisherInformation(startDate, endDate time.Time) ([]*models.PublisherInformation, error) {
	// sql query to sum up the fields of interest for a publisher between a start and end date
	query := fmt.Sprintf(`SELECT publisher_id, SUM(impressions), SUM(requests), SUM(clicks), SUM(revenue) FROM publishers_info 
	WHERE date_created BETWEEN '%s' AND '%s' GROUP BY publisher_id`, startDate.String(), endDate.String())
	fmt.Println("qq", query)

	rows, err := r.dbClient.Query(query)
	if err != nil {
		return []*models.PublisherInformation{}, nil
	}
	defer rows.Close()

	publisherInfoList := []*models.PublisherInformation{}
	for rows.Next() {
		var publisherInfo models.PublisherInformation
		if err := rows.Scan(&publisherInfo.Publisher.ID,
			&publisherInfo.Impressions,
			&publisherInfo.Requests,
			&publisherInfo.Clicks,
			&publisherInfo.Revenue); err != nil {
			return publisherInfoList, err
		}
		publisherInfoList = append(publisherInfoList, &publisherInfo)
	}

	if err = rows.Err(); err != nil {
		return publisherInfoList, err
	}

	return publisherInfoList, nil

}

// GetPublishers is a repository function that retrieves all publishers
func (r *Repository) GetPublishers() ([]models.Publisher, error) {
	// sql query to retrieve all publishers
	rows, err := r.dbClient.Query("SELECT id, name FROM publishers")
	if err != nil {
		return []models.Publisher{}, nil
	}
	defer rows.Close()

	publisherList := []models.Publisher{}
	for rows.Next() {
		var publisher models.Publisher
		if err := rows.Scan(&publisher.ID, &publisher.Name); err != nil {
			return []models.Publisher{}, err
		}
		publisherList = append(publisherList, publisher)
	}

	if err = rows.Err(); err != nil {
		return publisherList, err
	}

	return publisherList, nil
}
