package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/impactify-api/docs"
	"github.com/impactify-api/internal/src/config"
	"github.com/impactify-api/internal/src/handler"
	currency_service "github.com/impactify-api/internal/src/service/currency"
	providers "github.com/impactify-api/internal/src/service/currency/providers"
	publisher_service "github.com/impactify-api/internal/src/service/publisher"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Hello World")
	r := gin.Default()
	logger, _ := zap.NewProduction()

	config := config.InitConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Dbname)
	fmt.Println("dsn: ", dsn)
	// time.Sleep(40 * time.Second)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		logger.Fatal("Unable to open mysql connection", zap.Error(err))
		fmt.Printf("Unable to open mysql connection: %s", err)
		panic("Unable to open mysql connection")
	}

	db.SetConnMaxLifetime(time.Minute * 13)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		logger.Fatal("Unable to ping mysql", zap.Error(err))
		panic("Unable to ping mysql")
	}

	publisherService := publisher_service.NewService(db)
	currencyService := currency_service.NewCurrencyService()

	fixerProvider := providers.NewFixerProvider(config)
	exchangeRateProvider := providers.NewExchangeRateProvider(config)

	// Add providers to currency service
	currencyService.AddToCurrencyProviders(fixerProvider)
	currencyService.AddToCurrencyProviders(exchangeRateProvider)

	// InsertPubs(db)
	// logger.Info("Inserted publishers...")
	// InsertPubsInfo(db)
	// logger.Info("Inserted publishers info...")
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/publisher/data/:id/:currency", handler.GetPublisherInformation(publisherService, currencyService, logger))
	r.POST("/publisher/data/all/:currency", handler.GetAllPublisherInformation(publisherService, currencyService, logger))
	r.GET("/publisher/:id", handler.GetPublisherByID(publisherService))
	r.GET("/publishers", handler.GetPublishers(publisherService))
	r.Run()
}

func InsertPubs(db *sql.DB) {
	for i := 0; i < 1000000; i++ {
		_, err := db.Exec("INSERT INTO publishers (id, name) VALUES (?, ?)", i, gofakeit.Company())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func InsertPubsInfo(db *sql.DB) {
	for i := 0; i < 1000000; i++ {
		_, err := db.Exec("INSERT INTO publisher_info (publisher_id, requests, impressions, clicks, revenue, date_created) VALUES (?, ?, ?, ?, ?, ?)", i, gofakeit.Number(1, 100000000), gofakeit.Number(1, 100000000), gofakeit.Number(1, 100000000), gofakeit.Number(1, 100000000), gofakeit.Date())
		if err != nil {
			fmt.Println(err)
		}
	}
}
