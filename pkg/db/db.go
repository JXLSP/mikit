package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOptions struct {
	DBName string
	DBPass string
	DBUser string
	DBHost string
}

func (p *PostgresOptions) getDSN() string {
	return fmt.Sprintf(`host=%s user=%s password=%s dbname=%s sslmode=disable`,
		p.DBHost,
		p.DBUser,
		p.DBPass,
		p.DBName,
	)
}

func NewPostgresConnection(opts *PostgresOptions) (*gorm.DB, error) {
	dsn := opts.getDSN()

	// 打开数据库连接
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	// 获取底层SQL数据库连接以配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库连接失败: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期

	return db, nil
}
