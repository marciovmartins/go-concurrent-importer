package container

import (
	"go-concurrent-importer/internal/adapter/database"
	"go-concurrent-importer/internal/adapter/repository/gormRepo"
	"go-concurrent-importer/internal/service"
)


 type CliApp struct {
	SegmentationService *service.SegmentationService
 }

 func NewCliApp() (*CliApp, error) {
	db := database.GetGormDB()
	segRepo := gormRepo.NewSegmentationGormRepository(db)
	segService := service.NewSegmentationService(segRepo)

	return &CliApp{
		SegmentationService: segService,
	}, nil
 }
