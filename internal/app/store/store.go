package store

import "gorm.io/gorm"

type Store struct {
	ds *gorm.DB
}
