package testutil

import (
	"go-concurrent-importer/internal/adapter/entity"
	"sync/atomic"
	"time"
)

type FakeRepository struct {
	db *DbFake
}

func NewFakeRepository(db *DbFake) *FakeRepository {
	return &FakeRepository{
		db,
	}
}

func (s *FakeRepository) Save(data *entity.Segmentation) error {
	lastID := atomic.AddInt64(&s.db.SegsCounter, 1)

	data.ID = lastID
	data.CreatedAt = time.Now().UTC()
	data.UpdatedAt = time.Now().UTC()

	s.db.Segs.Store(lastID, []entity.Segmentation{*data})

	return nil
}

func (s *FakeRepository) SaveBatch(dataSet []*entity.Segmentation) error {
	for _, data := range dataSet {
		lastID := atomic.AddInt64(&s.db.SegsCounter, 1)

		data.ID = lastID
		data.CreatedAt = time.Now().UTC()
		data.UpdatedAt = time.Now().UTC()

		s.db.Segs.Store(lastID, []entity.Segmentation{*data})
	}

	return nil
}
