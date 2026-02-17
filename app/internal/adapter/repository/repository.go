package repository

import "go-concurrent-importer/internal/adapter/entity"

type Segmentation interface {
	Save(entity.Segmentation) error
}
