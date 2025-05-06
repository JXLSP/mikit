package report

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlReporter struct{}

// Connect 尝试使用提供的凭据连接MySQL数据库，检测弱口令
func (r *MysqlReporter) Connect(target, username, password string) (bool, error) {
	// 构建MySQL连接字符串: username:password@tcp(host:port)/database
	connStr := fmt.Sprintf("%s:%s@tcp(%s)", username, password, target)

	// 设置连接超时
	db, err := sqlx.ConnectContext(context.Background(), "mysql", connStr)
	if err != nil {
		// 连接失败，可能是凭据错误或网络问题
		return false, fmt.Errorf("连接MySQL失败: %v", err)
	}
	defer db.Close()

	// 尝试执行简单查询验证连接
	if err := db.Ping(); err != nil {
		return false, fmt.Errorf("MySQL连接验证失败: %v", err)
	}

	// 连接成功，表示使用这些凭据可以访问数据库
	return true, nil
}

// Unauthorized 检测MySQL数据库是否存在未授权访问漏洞
func (r *MysqlReporter) Unauthorized(target string) (bool, error) {
	// 尝试使用空密码和常见默认用户名连接
	defaultUsers := []string{"root", "mysql", "admin", ""}

	for _, user := range defaultUsers {
		// 尝试空密码
		connStr := fmt.Sprintf("%s:@tcp(%s)", user, target)

		// 设置短超时，避免长时间等待
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		db, err := sqlx.ConnectContext(ctx, "mysql", connStr)
		if err == nil {
			db.Close()
			// 如果能够连接成功，说明存在未授权访问
			return true, fmt.Errorf("发现MySQL未授权访问漏洞，使用用户名: %s 和空密码可以连接", user)
		}
	}

	return false, nil
}
