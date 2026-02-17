package model

import "time"

type Segmentation struct {
	ID               uint64    `gorm:"primaryKey"`
	UserID           int64     `gorm:"not null"`
	Type 			 string    `gorm:"not null"`
	Name 			 string    `gorm:"not null"`
	Data             any       `gorm:"type:jsonb;not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}
