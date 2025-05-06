package report

import (
	"context"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type MssqlReporter struct{}

// Connect 尝试使用提供的凭据连接MSSQL数据库，检测弱口令
func (r *MssqlReporter) Connect(target, username, password string) (bool, error) {
	// 构建MSSQL连接字符串: sqlserver://username:password@host:port?database=dbname
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s", username, password, target)

	// 设置连接超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "sqlserver", connStr)
	if err != nil {
		// 连接失败，可能是凭据错误或网络问题
		return false, fmt.Errorf("连接MSSQL失败: %v", err)
	}
	defer db.Close()

	// 尝试执行简单查询验证连接
	if err := db.Ping(); err != nil {
		return false, fmt.Errorf("MSSQL连接验证失败: %v", err)
	}

	return true, nil
}

// Unauthorized 检测MSSQL数据库是否存在未授权访问漏洞
func (r *MssqlReporter) Unauthorized(target string) (bool, error) {
	// 尝试使用空密码和常见默认用户名连接
	defaultUsers := []string{"sa", "admin", "mssql", ""}

	for _, user := range defaultUsers {
		// 尝试空密码
		connStr := fmt.Sprintf("sqlserver://%s:@%s", user, target)

		// 设置短超时，避免长时间等待
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		db, err := sqlx.ConnectContext(ctx, "sqlserver", connStr)
		if err == nil {
			db.Close()
			// 如果能够连接成功，说明存在未授权访问
			return true, fmt.Errorf("发现MSSQL未授权访问漏洞，使用用户名: %s 和空密码可以连接", user)
		}
	}

	return false, nil
}
