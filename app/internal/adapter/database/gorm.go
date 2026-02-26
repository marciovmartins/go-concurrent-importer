package database

import (
	"fmt"
	"go-concurrent-importer/config"
	"go-concurrent-importer/internal/adapter/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGormDB(cfg config.Database) (*gorm.DB, error) {
	strConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Pass,
		cfg.DB,
		cfg.Port,
		cfg.SSLmode,
	)

	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failure on database connection: %w", err)
	}

	// TODO: alter automigrate to goMigrate
	err = db.Statement.AutoMigrate(&model.Segmentation{})
	if err != nil {
		return nil, fmt.Errorf("failure on automigrate models: %w", err)
	}

	return db, nil
}
