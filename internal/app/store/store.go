package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once          sync.Once
	storeInstance *Store
)

// Store 提供数据库访问功能的结构体
type Store struct {
	ds *gorm.DB
}

// NewStore 创建并返回Store实例，使用单例模式确保只有一个实例
func NewStore(instance *gorm.DB) *Store {
	once.Do(func() {
		storeInstance = &Store{
			ds: instance,
		}
	})
	return storeInstance
}

// DB 返回底层数据库连接
func (s *Store) DB() *gorm.DB {
	return s.ds
}

// WithTx 在事务中执行函数
func (s *Store) WithTx(fn func(tx *gorm.DB) error) error {
	return s.ds.Transaction(fn)
}
