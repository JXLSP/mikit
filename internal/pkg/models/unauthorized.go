package models

import "time"

// UnauthorizedAccess 存储未授权访问检测的详细信息
type UnauthorizedAccess struct {
	ID             int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	TaskID         string    `gorm:"column:task_id" json:"task_id"`
	Target         string    `gorm:"column:target" json:"target"`
	Port           int       `gorm:"column:port" json:"port"`
	Service        string    `gorm:"column:service" json:"service"`
	Protocol       string    `gorm:"column:protocol" json:"protocol"`
	AccessPath     string    `gorm:"column:access_path" json:"access_path"`
	AccessMethod   string    `gorm:"column:access_method" json:"access_method"`
	ResponseCode   int       `gorm:"column:response_code" json:"response_code"`
	ResponseSize   int64     `gorm:"column:response_size" json:"response_size"`
	ExposedData    string    `gorm:"column:exposed_data" json:"exposed_data"`
	ExploitDetails string    `gorm:"column:exploit_details" json:"exploit_details"`
	RiskLevel      string    `gorm:"column:risk_level" json:"risk_level"`
	Remediation    string    `gorm:"column:remediation" json:"remediation"`
	Screenshot     string    `gorm:"column:screenshot" json:"screenshot"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *UnauthorizedAccess) TableName() string {
	return "unauthorized_access"
}
