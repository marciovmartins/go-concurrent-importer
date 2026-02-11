package entity

import "time"

type Segmentation struct {
	ID 				 int64
	UserID           int64
	Type 			 string
	Name 			 string
	Data 			 string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}