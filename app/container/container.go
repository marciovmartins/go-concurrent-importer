package container

import (
	"go-concurrent-importer/config"
	"go-concurrent-importer/internal/adapter/database"
	"go-concurrent-importer/internal/adapter/repository/gormrepo"
	"go-concurrent-importer/internal/service"
)

type Container struct {
	SegmentationService *service.Segmentation
}

func New(cfg *config.Config) (*Container, error) {
	db, err := database.GetGormDB(cfg.Database)
	if err != nil {

	}
	segRepo := gormrepo.NewSegmentation(db)

	segService := service.NewSegmentation(segRepo)

	return &Container{
		SegmentationService: segService,
	}, nil
}
