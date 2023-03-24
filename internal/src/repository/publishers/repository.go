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

func (r *Repository) GetPublisher(id string) *models.Publisher {
	// sql query to get a publisher by id
	// SELECT * FROM publisher WHERE id = ?
	var publisher models.Publisher
	err := r.dbClient.QueryRow("SELECT id, name FROM publishers WHERE id = ?", id).Scan(&publisher.ID, &publisher.Name)
	if err != nil {
		return &models.Publisher{}
	}

	return &publisher
}

func (r *Repository) GetPublisherByID(id string) *models.Publisher {
	// sql query to get a publisher by id
	// SELECT * FROM publisher WHERE id = ?
	var publisher models.Publisher
	err := r.dbClient.QueryRow("SELECT id, name FROM publishers WHERE id = ?", id).Scan(&publisher.ID, &publisher.Name)
	if err != nil {
		return &models.Publisher{}
	}

	return &publisher
}

func (r *Repository) GetPublisherInformationByID(id string, startDate, endDate time.Time) *models.PublisherInformation {
	var publisherInfo models.PublisherInformation
	query := fmt.Sprintf(`SELECT publisher_id, SUM(impressions), SUM(requests), SUM(clicks), SUM(revenue) FROM publishers_info 
	WHERE publisher_id = %s AND date_created BETWEEN '%s' AND '%s' GROUP BY publisher_id`, id, startDate.String(), endDate.String())
	fmt.Println("qq", query)
	err := r.dbClient.QueryRow(query).
		Scan(&publisherInfo.Publisher.ID,
			&publisherInfo.Requests,
			&publisherInfo.Impressions,
			&publisherInfo.Clicks,
			&publisherInfo.Revenue)

	if err != nil {
		return &models.PublisherInformation{}
	}

	return &publisherInfo
}

func (r *Repository) GetAllPublisherInformation(id string, startDate, endDate time.Time) []models.PublisherInformation {
	// sql query to sum up the revenue for a publisher between a start and end date
	// SELECT SUM(revenue) FROM publisher WHERE id = ? AND date BETWEEN ? AND ?
	query := fmt.Sprintf(`SELECT publisher_id, SUM(impressions), SUM(requests), SUM(clicks), SUM(revenue) FROM publishers_info 
	WHERE publisher_id = %s AND date_created BETWEEN '%s' AND '%s' GROUP BY publisher_id`, id, startDate.String(), endDate.String())

	rows, err := r.dbClient.Query(query)
	if err != nil {
		return []models.PublisherInformation{}
	}
	defer rows.Close()

	publisherInfoList := []models.PublisherInformation{}
	for rows.Next() {
		var publisherInfo models.PublisherInformation
		if err := rows.Scan(&publisherInfo.Publisher.ID,
			&publisherInfo.Requests,
			&publisherInfo.Impressions,
			&publisherInfo.Clicks,
			&publisherInfo.Revenue); err != nil {
			return publisherInfoList
		}
		publisherInfoList = append(publisherInfoList, publisherInfo)
	}

	if err = rows.Err(); err != nil {
		return publisherInfoList
	}

	return publisherInfoList

}

func (r *Repository) GetPublishers() []models.Publisher {
	// sql query to sum up the revenue for a publisher between a start and end date
	// SELECT SUM(revenue) FROM publisher WHERE id = ? AND date BETWEEN ? AND ?
	rows, err := r.dbClient.Query("SELECT id, name FROM publishers")
	if err != nil {
		return []models.Publisher{}
	}
	defer rows.Close()

	publisherList := []models.Publisher{}
	for rows.Next() {
		var publisher models.Publisher
		if err := rows.Scan(&publisher.ID, &publisher.Name); err != nil {
			return []models.Publisher{}
		}
		publisherList = append(publisherList, publisher)
	}

	if err = rows.Err(); err != nil {
		return publisherList
	}

	return publisherList
}
