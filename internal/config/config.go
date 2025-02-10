package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type S3 struct {
	Bucket string `env:"S3_BUCKET"`
}

type SQS struct {
	QueueURL string `env:"SQS_QUEUE_URL"`
}

type Config struct {
	S3  S3
	SQS SQS
	DB  Database
	API API
}

type API struct {
	Port string `env:"API_PORT"`
}

func Load() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}

type Database struct {
	User     string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
}

func (db *Database) GetDNS() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
