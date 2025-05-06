package report

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PgReporter struct{}

func (r *PgReporter) Connect(target, username, password string) (bool, error) {
	connStr := "postgres://" + username + ":" + password + "@" + target + "?sslmode=disable"

	// 设置连接超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", connStr)
	if err != nil {
		return false, err
	}
	defer db.Close()

	// 测试连接
	if err := db.PingContext(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func (r *PgReporter) Unauthorized(target string) (bool, error) {
	defaultUsers := []string{"postgres", "root", "admin"}

	for _, user := range defaultUsers {
		connStr := "postgres://" + user + ":" + user + "@" + target + "?sslmode=disable"

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		db, err := sqlx.ConnectContext(ctx, "postgres", connStr)
		if err == nil {
			db.Close()
			return false, fmt.Errorf("unauthorized access detected for user: %s", user)
		}
	}
	return true, nil
}
