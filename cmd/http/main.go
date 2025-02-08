package main

import (
	"log"

	_ "github.com/pangolin-do-golang/thumb-processor-api/docs"
	dbAdapter "github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/db"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Thumb processor worker
// @version 0.1.0
// @description Hackathon

// @host localhost:8080
// @BasePath /
func main() {
	_, err := initDb()
	if err != nil {
		panic(err)
	}

	/**

	_ := dbAdapter.NewPostgresThumbRepository(db)

	_ := thumb.NewThumbService()

	_ := dbAdapter.NewPostgresThumbRepository(db)

	**/

	restServer := server.NewRestServer(&server.RestServerOptions{})

	restServer.Serve()
}

func initDb() (*gorm.DB, error) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

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
