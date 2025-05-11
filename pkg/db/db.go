package db

import (
	"fmt"

	"gorm.io/gorm"
)

type PostgresOptions struct {
	DBName string
	DBPass string
	DBUser string
	DBHost string
}

func (p *PostgresOptions) getDSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s`,
		p.DBUser,
		p.DBPass,
		p.DBHost,
		p.DBName,
	)
}

func NewPostgresConnection(opts *PostgresOptions) (*gorm.DB, error) {
	return nil, nil
}
