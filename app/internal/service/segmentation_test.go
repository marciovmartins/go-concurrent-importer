package service

import (
	"go-concurrent-importer/internal/adapter/entity"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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



func TestProcessRecord(t *testing.T) {
	// arrange
	db := newDBFake()
	repo := newSegFakeRepo(db)
	srvc := NewSegmentationService(repo)

	record := []string{
		"1",
		"PATIENT",
		"João Silva",
		`{"age":45,"gender":"M","risk":"high"}`,
	}

	// act
	srvc.ProcessRecord(record)

	// assert
	assert.Equal(t, db.segsCounter, int64(1))
}