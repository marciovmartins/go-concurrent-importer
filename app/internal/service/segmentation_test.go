package service

import (
	"go-concurrent-importer/internal/adapter/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


type dbFake struct{
	segmentations map[int64][]entity.Segmentation
}

func newDBFake() *dbFake {
	return &dbFake{
		segmentations: make(map[int64][]entity.Segmentation),
	}
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
	lastID := int64(len(s.db.segmentations) + 1)

	data.ID = lastID
	data.CreatedAt = time.Now().UTC()
	data.UpdatedAt = time.Now().UTC()

	s.db.segmentations[lastID] = append(s.db.segmentations[lastID], data)

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
	assert.Equal(t, len(db.segmentations), 1)
}