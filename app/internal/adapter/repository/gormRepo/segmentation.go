package gormRepo

import (
	"go-concurrent-importer/internal/adapter/entity"

	"gorm.io/gorm"
)

type SegmentationGorm struct {
	db *gorm.DB
}

func NewSegmentationGormRepository(db *gorm.DB) *SegmentationGorm {
	return &SegmentationGorm{db: db}
}

func (r *SegmentationGorm) Save(data *entity.Segmentation) error {
	return r.db.Create(data).Error
}

func (r *SegmentationGorm) SaveBatch(dataSet []*entity.Segmentation) error {
    return r.db.CreateInBatches(dataSet, len(dataSet)).Error;
}