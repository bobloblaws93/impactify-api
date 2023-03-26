package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/gin-gonic/gin"
	"github.com/impactify-api/internal/src/models"
	currency_service "github.com/impactify-api/internal/src/service/currency"
	"github.com/impactify-api/internal/src/service/currency/providers"
	publisher_service "github.com/impactify-api/internal/src/service/publisher"
)

// unit tests for handlers
func Test_GetPublisherByID(t *testing.T) {

	pubService := publisher_service.NewPublisherMock()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	GetPublisherByID(pubService)(c)
	assert.Equal(t, 200, w.Code)

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	responseID := int64(response["id"].(float64))
	assert.Equal(t, int64(123), responseID)
	assert.Equal(t, "test", response["name"])
}

func Test_GetPublisherInformation(t *testing.T) {
	pubService := publisher_service.NewPublisherMock()
	currencyService := currency_service.NewCurrencyMock()
	currencyService.AddToCurrencyProviders(providers.NewProviderMock())
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	logger := zaptest.NewLogger(t)

	assert.Equal(t, 200, w.Code)
	publisherRequest := models.PublisherTimeRequest{
		StartDate: "2019-01-01",
		EndDate:   "2019-01-31",
	}

	body, _ := json.Marshal(publisherRequest)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
	GetPublisherInformation(pubService, currencyService, logger)(c)

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	publisherModel := response["publisher"].(map[string]interface{})
	revenue := response["revenue"].(float64)
	requests := int64(response["requests"].(float64))
	clicks := int64(response["clicks"].(float64))

	assert.Equal(t, publisherModel["name"], "test")
	assert.Equal(t, int64(100), requests)
	assert.Equal(t, float64(90), revenue)
	assert.Equal(t, int64(100), clicks)

}

func Test_GetAllPublisherInformation(t *testing.T) {
	pubService := publisher_service.NewPublisherMock()
	currencyService := currency_service.NewCurrencyMock()
	currencyService.AddToCurrencyProviders(providers.NewProviderMock())
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	logger := zaptest.NewLogger(t)
	assert.Equal(t, 200, w.Code)
	publisherRequest := models.PublisherTimeRequest{
		StartDate: "2019-01-01",
		EndDate:   "2019-01-31",
	}

	body, _ := json.Marshal(publisherRequest)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
	GetAllPublisherInformation(pubService, currencyService, logger)(c)

	var response []gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	publisherModel := response[0]["publisher"].(map[string]interface{})
	revenue := response[0]["revenue"].(float64)
	requests := int64(response[0]["requests"].(float64))
	clicks := int64(response[0]["clicks"].(float64))

	assert.Equal(t, publisherModel["name"], "test")
	assert.Equal(t, int64(100), requests)
	assert.Equal(t, float64(90), revenue)
	assert.Equal(t, int64(100), clicks)

}

func Test_GetPublishers(t *testing.T) {
	pubService := publisher_service.NewPublisherMock()
	currencyService := currency_service.NewCurrencyMock()
	currencyService.AddToCurrencyProviders(providers.NewProviderMock())
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	assert.Equal(t, 200, w.Code)
	publisherRequest := models.PublisherTimeRequest{
		StartDate: "2019-01-01",
		EndDate:   "2019-01-31",
	}

	logger := zaptest.NewLogger(t)

	body, _ := json.Marshal(publisherRequest)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
	GetAllPublisherInformation(pubService, currencyService, logger)(c)

	var response []gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	publisherModel := response[0]["publisher"].(map[string]interface{})

	assert.Equal(t, publisherModel["name"], "test")

}
