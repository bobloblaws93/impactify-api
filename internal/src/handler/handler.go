package handler

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/impactify-api/internal/src/constants"
	"github.com/impactify-api/internal/src/models"
	currency_service "github.com/impactify-api/internal/src/service/currency"
	service "github.com/impactify-api/internal/src/service/publisher"
	"go.uber.org/zap"
)

// GetPublishers godoc
// @Summary      GetPublisherByID is a route meant to a specific publisher by id
// @Description  Get publishers
// @Tags         publishers
// @Produce      json
// @Param        id    path   string true "param for publisher id"
// @Success      200  {object}  models.Publisher
// @Failure      404  {string} string "publisher not found"
// @Router       /publisher/{id} [get]
func GetPublisherByID(service service.IPublisherService) gin.HandlerFunc {
	return func(c *gin.Context) {
		publisher, err := service.RetrievePublisher(c.Param("id"))
		if err != nil {
			c.JSON(404, "publisher not found")
			return
		}
		c.JSON(200, publisher)
	}
}

// GetPublishers godoc
// @Summary      GetPublisherInformation is a route meant to retrieve AGGREGATE information on a specific publisher by id
// @Description  Get publishers data information
// @Tags         publishers
// @Produce      json
// @Param        id    path   string true "param for publisher id. By default, 1 and 2 are ids that are seeded in the database"
// @Param        currency    path   string true "param for currency  [SGD, USD, EUR, ETC...]"
// @Param        publishertimerequest body models.PublisherTimeRequest true "time range for publisher data"
// @Success      200  {object}  models.PublisherInformation
// @Failure      404  {string} string "unable to retrieve publisher information"
// @Failure      400  {string} string "unable to parse enddate (format: 'yyyy-mm-dd')"
// @Failure      400  {string} string "unable to parse startdate (format: 'yyyy-mm-dd')"
// @Failure      400  {string} string "start date is after end date"
// @Router       /publisher/data/{id}/{currency} [post]
func GetPublisherInformation(service service.IPublisherService,
	currencyService currency_service.ICurrencyService, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var pubRequest models.PublisherTimeRequest
		c.Header("Content-Type", "application/json")

		// bind time request
		if err = c.BindJSON(&pubRequest); err != nil {
			logger.Sugar().Errorf("unable to unmarshal request)', %s", err)
			c.JSON(404, "bad request")
			return
		}

		// check if provider is valid, if not default to fixer
		var provider string
		providerMapping := currencyService.GetProviderMapping()
		if _, ok := providerMapping[pubRequest.CurrencyProvider]; ok {
			provider = pubRequest.CurrencyProvider
		} else {
			provider = constants.FIXER
		}

		// parse start date
		startDate, err := time.Parse("2006-01-02", pubRequest.StartDate)
		if err != nil {
			logger.Sugar().Errorf("unable to parse enddate (format: 'yyyy-mm-dd')', %s", err)
			c.JSON(400, "unable to parse startdate (format: 'yyyy-mm-dd')')")
			return
		}

		// parse end date
		endDate, err := time.Parse("2006-01-02", pubRequest.EndDate)
		if err != nil {
			logger.Sugar().Errorf("unable to parse enddate (format: 'yyyy-mm-dd')', %s", err)
			c.JSON(400, "unable to parse startdate (format: 'yyyy-mm-dd')')")
			return
		}

		// check if start date is before end date
		if startDate.After(endDate) {
			logger.Sugar().Errorf("start date is after end date")
			c.JSON(400, "start date is after end date")
			return
		}

		// retrieve publisher information
		publisher, err := service.RetrievePublisherInformation(c.Param("id"), startDate, endDate)

		// handle error if we cannot retrieve publisher information
		if err != nil {
			logger.Sugar().Errorf("unable to retrieve publisher information", err)
			c.JSON(404, "unable to retrieve publisher information")
			return
		}

		// if currency is USD, return the revnue as is
		if strings.ToUpper(c.Param("currency")) == constants.USD {
			publisher.Revenue = publisher.Revenue * 1
			c.JSON(200, publisher)
			return
		}

		// if not USD, convert the revenue to the currency given
		currencyModel := currencyService.ReturnRate(provider, c.Param("currency"))

		publisher.Revenue = publisher.Revenue * currencyModel.Rate

		c.JSON(200, publisher)
	}
}

// GetPublishers godoc
// @Summary      GetAllPublisherInformation is a route meant to retrieve all data rows for a specific publisher
// @Description  Get all publishers data information
// @Tags         publishers
// @Produce      json
// @Param        id    path   string true "param for publisher id. By default, 1 and 2 are ids that are seeded in the database"
// @Param        currency    path  string true "param for currency [SGD, USD, EUR, ETC...]"
// @Param        publishertimerequest body models.PublisherTimeRequest true "time range for publisher data"
// @Success      200  {array}  models.PublisherInformation
// @Failure      404  {string} string "unable to retrieve all publisher information"
// @Failure      400  {string} string "bad request"
// @Failure      400  {string} string "unable to parse enddate (format: 'yyyy-mm-dd')"
// @Failure      400  {string} string "unable to parse startdate (format: 'yyyy-mm-dd')"
// @Failure      400  {string} string "start date is after end date"
// @Router       /publisher/data/rows/{id}/{currency} [post]
func GetPublisherInformationRows(service service.IPublisherService, currencyService currency_service.ICurrencyService, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var pubRequest models.PublisherTimeRequest
		c.Header("Content-Type", "application/json")

		if err = c.BindJSON(&pubRequest); err != nil {
			logger.Sugar().Errorf("unable to unmarshal request)', %s", err)
			c.JSON(400, "bad request")
			return
		}

		// check if provider is valid, if not default to fixer
		var provider string
		providerMapping := currencyService.GetProviderMapping()
		if _, ok := providerMapping[pubRequest.CurrencyProvider]; ok {
			provider = pubRequest.CurrencyProvider
		} else {
			provider = constants.FIXER
		}

		startDate, err := time.Parse("2006-01-02", pubRequest.StartDate)
		if err != nil {
			logger.Sugar().Errorf("unable to parse enddate (format: 'yyyy-mm-dd')', %s", err)
			c.JSON(400, "unable to parse startdate (format: 'yyyy-mm-dd')')")
			return
		}

		endDate, err := time.Parse("2006-01-02", pubRequest.EndDate)
		if err != nil {
			logger.Sugar().Errorf("unable to parse enddate (format: 'yyyy-mm-dd')', %s", err)
			c.JSON(400, "unable to parse enddate (format: 'yyyy-mm-dd')')")
			return
		}

		// check if start date is before end date
		if startDate.After(endDate) {
			logger.Sugar().Errorf("start date is after end date")
			c.JSON(400, "start date is after end date")
			return
		}

		// Retrieve data for all pubs
		publisherInfoList, err := service.RetrieveAllPublisherRows(c.Param("id"), startDate, endDate)
		if err != nil {
			logger.Sugar().Errorf("unable to retrieve all publisher information", err)
			c.JSON(404, "unable to retrieve all publisher information")
			return
		}

		// if currency is USD, return the revnue as is
		var currencyModel *models.Currency
		if strings.ToUpper(c.Param("currency")) == constants.USD {
			currencyModel = &models.Currency{Rate: 1}
		} else {
			currencyModel = currencyService.ReturnRate(provider, c.Param("currency"))
		}

		if currencyModel.Rate == 0 {
			logger.Sugar().Errorf("unable to parse currency:', %s", c.Param("currency"))
			c.JSON(400, "invalid currency")
			return
		}

		for _, publisherInfo := range publisherInfoList {
			publisherInfo.Revenue = publisherInfo.Revenue * currencyModel.Rate
		}
		c.JSON(200, publisherInfoList)
	}
}

// GetPublishers godoc
// @Summary      GetAllPublisherInformation is a route meant to retrieve information on ALL publishers
// @Description  Get all publishers data information
// @Tags         publishers
// @Produce      json
// @Param        currency    path  string true "param for currency [SGD, USD, EUR, ETC...]"
// @Param        publishertimerequest body models.PublisherTimeRequest true "time range for publisher data"
// @Success      200  {array}  models.PublisherInformation
// @Failure      404  {string} string "unable to retrieve all publisher information"
// @Failure      400  {string} string "bad request"
// @Failure      400  {string} string "unable to parse enddate (format: 'yyyy-mm-dd')"
// @Failure      400  {string} string "unable to parse startdate (format: 'yyyy-mm-dd')"
// @Failure      400  {string} string "start date is after end date"
// @Router       /publisher/data/all/{currency} [post]
func GetAllPublisherInformation(service service.IPublisherService, currencyService currency_service.ICurrencyService, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var pubRequest models.PublisherTimeRequest
		c.Header("Content-Type", "application/json")

		if err = c.BindJSON(&pubRequest); err != nil {
			logger.Sugar().Errorf("unable to unmarshal request)', %s", err)
			c.JSON(400, "bad request")
			return
		}

		// check if provider is valid, if not default to fixer
		var provider string
		providerMapping := currencyService.GetProviderMapping()
		if _, ok := providerMapping[pubRequest.CurrencyProvider]; ok {
			provider = pubRequest.CurrencyProvider
		} else {
			provider = constants.FIXER
		}

		startDate, err := time.Parse("2006-01-02", pubRequest.StartDate)
		if err != nil {
			logger.Sugar().Errorf("unable to parse enddate (format: 'yyyy-mm-dd')', %s", err)
			c.JSON(400, "unable to parse startdate (format: 'yyyy-mm-dd')')")
			return
		}

		endDate, err := time.Parse("2006-01-02", pubRequest.EndDate)
		if err != nil {
			logger.Sugar().Errorf("unable to parse enddate (format: 'yyyy-mm-dd')', %s", err)
			c.JSON(400, "unable to parse enddate (format: 'yyyy-mm-dd')')")
			return
		}

		// check if start date is before end date
		if startDate.After(endDate) {
			logger.Sugar().Errorf("start date is after end date")
			c.JSON(400, "start date is after end date")
			return
		}

		// Retrieve data for all pubs
		publisherInfoList, err := service.RetrieveAllPublisherInformation(startDate, endDate)
		if err != nil {
			logger.Sugar().Errorf("unable to retrieve all publisher information", err)
			c.JSON(404, "unable to retrieve all publisher information")
			return
		}

		// if currency is USD, return the revnue as is
		var currencyModel *models.Currency
		if strings.ToUpper(c.Param("currency")) == constants.USD {
			currencyModel = &models.Currency{Rate: 1}
		} else {
			currencyModel = currencyService.ReturnRate(provider, c.Param("currency"))
		}

		if currencyModel.Rate == 0 {
			logger.Sugar().Errorf("unable to parse currency:', %s", c.Param("currency"))
			c.JSON(400, "invalid currency")
			return
		}

		for _, publisherInfo := range publisherInfoList {
			publisherInfo.Revenue = publisherInfo.Revenue * currencyModel.Rate
		}
		c.JSON(200, publisherInfoList)
	}
}

// GetPublishers godoc
// @Summary      GetPublishers is a route meant to fetch all publishers
// @Description  List publishers
// @Tags         publishers
// @Produce      json
// @Success      200  {array}  models.Publisher
// @Failure      404  {string} string "could not retrieve publisher list"
// @Router       /publishers [get]
func GetPublishers(service service.IPublisherService) gin.HandlerFunc {
	return func(c *gin.Context) {
		publishers, err := service.RetrieveAllPublishers()
		if err != nil {
			c.JSON(404, "could not retrieve publishers")
			return
		}
		c.JSON(200, publishers)
	}
}
