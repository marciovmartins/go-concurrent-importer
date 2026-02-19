package service

import (
	"fmt"
	"go-concurrent-importer/internal/adapter/entity"
	"go-concurrent-importer/internal/adapter/repository"
	"strconv"
	"strings"
	"time"
)

type SegmentationService struct {
	segRepo repository.Segmentation
}

func NewSegmentationService(segRepo repository.Segmentation) *SegmentationService {
	return &SegmentationService{
		segRepo,
	}
}

func (s *SegmentationService) ProcessBatch(records [][]string) ([]*entity.Segmentation, []error) {
    var segmentations []*entity.Segmentation
    var errors []error
    
    for i, record := range records {
		// TODO: ProcessRecord ONLY for validation in `dev time`
        seg, err := s.processRecord(record)
        if err != nil {
            errors = append(errors, fmt.Errorf("linha %d: %w", i+1, err))
            continue
        }
        segmentations = append(segmentations, seg)
    }

    return segmentations, errors
}

func (s *SegmentationService) processRecord(record []string) (*entity.Segmentation, error) {

	if len(record) < 4 {
		return nil, fmt.Errorf("invalid columns: %v", record)
	}

	userID, _ := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
	typ := strings.TrimSpace(record[1])
	name := strings.TrimSpace(record[2])
	dataStr := strings.TrimSpace(record[3])

	segmentation := &entity.Segmentation{
		UserID:           int64(userID),
		Type: 			  typ,
		Name: 			  name,
		Data:             dataStr,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := s.segRepo.Save(segmentation)
	if err != nil {
		return nil, err
	}

	return segmentation, nil
}