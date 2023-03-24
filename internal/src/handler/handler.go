package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/impactify-api/internal/src/models"
	"github.com/impactify-api/internal/src/service"
)

func GetPublisherByID(service service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		publisher := service.RetrievePublisher(c.Param("id"))
		c.JSON(200, publisher)
	}
}

func GetPublisherInformation(service service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var timePeriod models.PublisherTimeRequest
		c.Header("Content-Type", "application/json")

		if err = c.BindJSON(&timePeriod); err != nil {
			fmt.Printf("Error: %s", err)
			c.JSON(404, "bad request")
			return
		}
		if err != nil {
			panic(err)
		}
		startDate, err := time.Parse("2006-01-02", timePeriod.StartDate)
		if err != nil {
			panic(err)
		}

		endDate, err := time.Parse("2006-01-02", timePeriod.EndDate)
		if err != nil {
			panic(err)
		}

		publisher := service.RetrievePublisherRevenue(c.Param("id"), c.Param("currency"), startDate, endDate)
		c.JSON(200, publisher)
	}
}

func GetPublishers(service service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		publishers := service.RetrieveAllPublishers()
		c.JSON(200, publishers)
	}
}
