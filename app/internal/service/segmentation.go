package service

import (
	"fmt"
	"go-concurrent-importer/internal/adapter/entity"
	"go-concurrent-importer/internal/adapter/repository"
	"strconv"
	"strings"
	"time"
)

type Segmentation struct {
	segRepo repository.Segmentation
}

func NewSegmentation(segRepo repository.Segmentation) *Segmentation {
	return &Segmentation{
		segRepo,
	}
}

func (s *Segmentation) ProcessBatch(records [][]string) ([]*entity.Segmentation, []error) {
	var segmentations []*entity.Segmentation
	var errors []error

	for i, record := range records {
		seg, err := mapRecordToSegmentationEntity(record)
		if err != nil {
			errors = append(errors, fmt.Errorf("linha %d: %w", i+1, err))
			continue
		}
		segmentations = append(segmentations, seg)
	}

	s.segRepo.SaveBatch(segmentations)

	return segmentations, errors
}

func mapRecordToSegmentationEntity(record []string) (*entity.Segmentation, error) {
	if len(record) < 4 {
		return nil, fmt.Errorf("invalid columns: %v", record)
	}

	userID, _ := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
	typ := strings.TrimSpace(record[1])
	name := strings.TrimSpace(record[2])
	dataStr := strings.TrimSpace(record[3])

	return &entity.Segmentation{
		UserID:    int64(userID),
		Type:      typ,
		Name:      name,
		Data:      dataStr,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
