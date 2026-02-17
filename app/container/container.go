package container

import (
	"go-concurrent-importer/internal/adapter/entity"
	"go-concurrent-importer/internal/service"
	"sync"
	"sync/atomic"
	"time"
)


type dbFake struct{
	segs        sync.Map // key: int64 | value: []entity.Segmentation
	segsCounter int64
}

func newDBFake() *dbFake {
	return &dbFake{}
}

type segFakeRepo struct{
	db *dbFake
}

func newSegFakeRepo(db *dbFake) *segFakeRepo{
	return &segFakeRepo{
		db,
	}
}

func (s *segFakeRepo) Save(data entity.Segmentation) error {
	lastID := atomic.AddInt64(&s.db.segsCounter, 1)

	data.ID = lastID
	data.CreatedAt = time.Now().UTC()
	data.UpdatedAt = time.Now().UTC()

	s.db.segs.Store(lastID, []entity.Segmentation{data})

	return nil
}


 type CliApp struct {
	SegmentationService *service.SegmentationService
 }

 func NewCliApp() (*CliApp, error) {
	db := newDBFake()
	segRepo := newSegFakeRepo(db)

	// db := database.GetGormDB()
	// segRepo := gormRepo.NewSegmentationGormRepository(db)
	segService := service.NewSegmentationService(segRepo)

	return &CliApp{
		SegmentationService: segService,
	}, nil
 }
