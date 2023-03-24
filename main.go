package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/impactify-api/internal/src/config"
	"github.com/impactify-api/internal/src/handler"
	"github.com/impactify-api/internal/src/service"
)

func main() {
	fmt.Println("Hello World")
	r := gin.Default()

	config := config.InitConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Dbname)
	fmt.Println("dsn: ", dsn)
	// time.Sleep(40 * time.Second)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Unable to open mysql connection: %s", err)
		panic("Unable to open mysql connection")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic("Unable to ping mysql")
	}

	publisherService := service.NewService(db)

	r.POST("/publisher/:id/:currency", handler.GetPublisherInformation(publisherService))
	r.POST("/getAllPublishers/:currency", handler.GetPublishers(publisherService))
	r.GET("/publishers/:id", handler.GetPublisherByID(publisherService))
	r.GET("/publishers", handler.GetPublishers(publisherService))
	r.Run()
}
