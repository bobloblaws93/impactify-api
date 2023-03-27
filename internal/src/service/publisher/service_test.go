package publisher_service

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_CreatePublisherService(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pubService := NewService(db)
	assert.Equal(t, pubService.PublisherRepo.GetDBClient(), db)

}
func Test_RetrievePublisher(t *testing.T) {
	// mock the db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pubService := NewService(db)

	// mock the query
	mock.ExpectQuery("SELECT id, name FROM publishers WHERE id = ?").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(int64(1), "test-pub"))

	publisher, err := pubService.RetrievePublisher("1")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, publisher.ID, int64(1))
}

func Test_RetrievePublisherRevenue(t *testing.T) {
	// mock the db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pubService := NewService(db)
	beginningDate := "2017-08-31"
	endingDate := "2017-09-01"
	layout := "2006-01-02"

	startDate, _ := time.Parse(layout, beginningDate)
	endDate, _ := time.Parse(layout, endingDate)
	query := fmt.Sprintf(`SELECT publisher_id, SUM(impressions), SUM(requests), SUM(clicks), SUM(revenue) FROM publishers_info WHERE publisher_id = %s AND date_created BETWEEN '%s' AND '%s' GROUP BY publisher_id`, "1", startDate.String(), endDate.String())

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"publisher_id", "SUM(impressions)", " SUM(requests)", "SUM(clicks)", "SUM(revenue)"}).
			AddRow(int64(1), "4000", "4000", "4000", "4000"))

	publisherInfo, err := pubService.RetrievePublisherInformation("1", startDate, endDate)
	if err != nil {
		t.Errorf("Error was not expected while getting publisher: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, publisherInfo.Publisher.ID, int64(1))
	assert.Equal(t, publisherInfo.Clicks, int64(4000))
	assert.Equal(t, publisherInfo.Impressions, int64(4000))
	assert.Equal(t, publisherInfo.Requests, int64(4000))
	assert.Equal(t, publisherInfo.Revenue, float64(4000))
}

func Test_GetAllPublisherInformation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	beginningDate := "2017-08-31"
	endingDate := "2017-09-01"
	layout := "2006-01-02"

	startDate, _ := time.Parse(layout, beginningDate)
	endDate, _ := time.Parse(layout, endingDate)

	defer db.Close()
	pubService := NewService(db)
	query := fmt.Sprintf(`SELECT publisher_id, SUM(impressions), SUM(requests), SUM(clicks), SUM(revenue) FROM publishers_info 
	WHERE date_created BETWEEN '%s' AND '%s' GROUP BY publisher_id`, startDate.String(), endDate.String())

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"publisher_id", "SUM(impressions)", " SUM(requests)", "SUM(clicks)", "SUM(revenue)"}).
			AddRow(int64(1), "4000", "4000", "4000", "4000").
			AddRow(int64(2), "6000", "6000", "6000", "6000"))

	publisherInfoList, err := pubService.RetrieveAllPublisherInformation(startDate, endDate)
	if err != nil {
		t.Errorf("Error was not expected while getting publisher: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, publisherInfoList[0].Publisher.ID, int64(1))
	assert.Equal(t, publisherInfoList[0].Clicks, int64(4000))
	assert.Equal(t, publisherInfoList[0].Impressions, int64(4000))
	assert.Equal(t, publisherInfoList[0].Requests, int64(4000))
	assert.Equal(t, publisherInfoList[0].Revenue, float64(4000))

	assert.Equal(t, publisherInfoList[1].Publisher.ID, int64(2))
	assert.Equal(t, publisherInfoList[1].Clicks, int64(6000))
	assert.Equal(t, publisherInfoList[1].Impressions, int64(6000))
	assert.Equal(t, publisherInfoList[1].Requests, int64(6000))
	assert.Equal(t, publisherInfoList[1].Revenue, float64(6000))
}
