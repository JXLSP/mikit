package models

import "time"

// Credential 存储认证凭证信息，用于弱口令检测
type Credential struct {
	ID           int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	TaskID       string    `gorm:"column:task_id" json:"task_id"`
	Target       string    `gorm:"column:target" json:"target"`
	Port         int       `gorm:"column:port" json:"port"`
	Service      string    `gorm:"column:service" json:"service"`
	Protocol     string    `gorm:"column:protocol" json:"protocol"`
	Username     string    `gorm:"column:username" json:"username"`
	Password     string    `gorm:"column:password" json:"password"`
	AuthType     string    `gorm:"column:auth_type" json:"auth_type"`
	IsDefault    bool      `gorm:"column:is_default" json:"is_default"`
	IsWeak       bool      `gorm:"column:is_weak" json:"is_weak"`
	Successful   bool      `gorm:"column:successful" json:"successful"`
	FailedReason string    `gorm:"column:failed_reason" json:"failed_reason"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (c *Credential) TableName() string {
	return "credentials"
}
