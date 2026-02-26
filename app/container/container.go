package container

import (
	"go-concurrent-importer/config"
	"go-concurrent-importer/internal/adapter/database"
	"go-concurrent-importer/internal/adapter/repository/gormRepo"
	"go-concurrent-importer/internal/service"
)

type Container struct {
	SegmentationService *service.SegmentationService
}

func New(cfg *config.Config) (*Container, error) {
	db, err := database.GetGormDB(cfg.Database)
	if err != nil {

	}
	segRepo := gormRepo.NewSegmentationGormRepository(db)

	segService := service.NewSegmentationService(segRepo)

	return &Container{
		SegmentationService: segService,
	}, nil
}
