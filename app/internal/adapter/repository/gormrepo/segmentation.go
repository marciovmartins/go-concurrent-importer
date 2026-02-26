package gormrepo

import (
	"go-concurrent-importer/internal/adapter/entity"

	"gorm.io/gorm"
)

type Segmentation struct {
	db *gorm.DB
}

func NewSegmentation(db *gorm.DB) *Segmentation {
	return &Segmentation{db: db}
}

func (r *Segmentation) Save(data *entity.Segmentation) error {
	return r.db.Create(data).Error
}

func (r *Segmentation) SaveBatch(dataSet []*entity.Segmentation) error {
	return r.db.CreateInBatches(dataSet, len(dataSet)).Error
}
