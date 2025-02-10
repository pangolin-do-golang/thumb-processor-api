package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/pangolin-do-golang/thumb-processor-api/docs"
	dbAdapter "github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/db"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/sqs"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/config"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/thumb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Thumb Processor API
// @version 0.1.0
// @description Hackathon

// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	databaseAdapter, err := newDatabaseConnection(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	queueAdapter, err := sqs.NewSQSThumbQueue(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	thumbRepository := dbAdapter.NewPostgresThumbRepository(databaseAdapter)
	thumbService := thumb.NewThumbService(thumbRepository, queueAdapter)

	restServer := server.NewRestServer(&server.RestServerOptions{
		ThumService: thumbService,
	})

	restServer.Serve(cfg)
}

func newDatabaseConnection(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DB.GetDNS()), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(
		&dbAdapter.ThumbPostgres{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
