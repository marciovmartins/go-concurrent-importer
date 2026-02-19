package container

import (
	"go-concurrent-importer/config"
	"go-concurrent-importer/internal/adapter/database"
	"go-concurrent-importer/internal/adapter/repository/gormRepo"
	"go-concurrent-importer/internal/service"
)

 type CliApp struct {
	SegmentationService *service.SegmentationService
 }

 func NewCliApp(cfg *config.Config) (*CliApp, error) {
	db, err := database.GetGormDB(cfg.Database)
	if err != nil {

	}
	segRepo := gormRepo.NewSegmentationGormRepository(db)

	segService := service.NewSegmentationService(segRepo)

	return &CliApp{
		SegmentationService: segService,
	}, nil
 }
