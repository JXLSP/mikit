package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
)

type Store struct {
	ds *gorm.DB
}
