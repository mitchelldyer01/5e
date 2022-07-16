package db

import (
	"os"

	"github.com/mitchelldyer01/characters-5e/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func New() *Repository {
	return &Repository{
		DB: Init(),
	}
}

func Init() *gorm.DB {
	url, exists := os.LookupEnv("DB_URL")
	if !exists {
		logrus.Fatalf("${DB_URL} missing")
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Failed connecting to DB: %s", err)
	}

	db.AutoMigrate(
		&models.Character{},
	)

	return db
}
