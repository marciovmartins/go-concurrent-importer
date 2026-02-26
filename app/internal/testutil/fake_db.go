package testutil

import "sync"

type DbFake struct {
	Segs        sync.Map // key: int64 | value: []entity.Segmentation
	SegsCounter int64
}

func NewDbFake() *DbFake {
	return &DbFake{}
}
