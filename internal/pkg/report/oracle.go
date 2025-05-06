package report

import (
	"context"
	"fmt"
	"time"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

type OracleReporter struct{}

// Connect 尝试使用提供的凭据连接Oracle数据库，检测弱口令
func (r *OracleReporter) Connect(target, username, password string) (bool, error) {
	// 构建Oracle连接字符串: username/password@host:port/service_name
	connStr := fmt.Sprintf("%s/%s@%s", username, password, target)

	// 设置连接超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "godror", connStr)
	if err != nil {
		// 连接失败，可能是凭据错误或网络问题
		return false, fmt.Errorf("连接Oracle失败: %v", err)
	}
	defer db.Close()

	// 尝试执行简单查询验证连接
	if err := db.Ping(); err != nil {
		return false, fmt.Errorf("Oracle连接验证失败: %v", err)
	}

	return true, nil
}

// Unauthorized 检测Oracle数据库是否存在未授权访问漏洞
func (r *OracleReporter) Unauthorized(target string) (bool, error) {
	// 尝试使用空密码和常见默认用户名连接
	defaultUsers := []string{"system", "sys", "oracle", "admin", ""}
	defaultPasswords := []string{"", "oracle", "password", "admin", "manager"}

	for _, user := range defaultUsers {
		for _, pass := range defaultPasswords {
			// 尝试默认用户名和密码组合
			connStr := fmt.Sprintf("%s/%s@%s", user, pass, target)

			// 设置短超时，避免长时间等待
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			db, err := sqlx.ConnectContext(ctx, "godror", connStr)
			if err == nil {
				db.Close()
				// 如果能够连接成功，说明存在未授权访问或弱口令
				return true, fmt.Errorf("发现Oracle未授权访问漏洞，使用用户名: %s 和密码: %s 可以连接", user, pass)
			}
		}
	}

	return false, nil
}
