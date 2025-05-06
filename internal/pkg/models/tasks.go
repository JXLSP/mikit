package models

import "time"

type Tasks struct {
	ID                int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	TaskID            string    `gorm:"column:task_id" json:"task_id"`
	TaskName          string    `gorm:"column:task_name;not null" json:"task_name"`
	TaskType          string    `gorm:"column:task_type" json:"task_type"`
	Targets           string    `gorm:"column:targets" json:"targets"`
	ConcurrencyNumber int       `gorm:"column:concurrency_number;default:10" json:"concurrency_number"`
	DepthNumber       int       `gorm:"column:depth_number;default:3" json:"depth_number"`
	Status            string    `gorm:"column:status;default:'pending'" json:"status"` // pending, running, completed, failed
	Progress          float64   `gorm:"column:progress;default:0" json:"progress"`     // 0-100
	StartTime         time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime           time.Time `gorm:"column:end_time" json:"end_time"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (t *Tasks) TableName() string {
	return "tasks"
}

type TaskResult struct {
	ID             int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	TaskID         string    `gorm:"column:task_id" json:"task_id"`
	Target         string    `gorm:"column:target" json:"target"`
	Port           int       `gorm:"column:port" json:"port"`
	Service        string    `gorm:"column:service" json:"service"`
	Protocol       string    `gorm:"column:protocol" json:"protocol"`
	VulnType       string    `gorm:"column:vuln_type" json:"vuln_type"`   // weak_password, unauthorized_access, etc.
	RiskLevel      string    `gorm:"column:risk_level" json:"risk_level"` // high, medium, low, info
	AuthType       string    `gorm:"column:auth_type" json:"auth_type"`   // basic, digest, form, etc.
	Username       string    `gorm:"column:username" json:"username"`
	Password       string    `gorm:"column:password" json:"password"`
	ExploitDetails string    `gorm:"column:exploit_details" json:"exploit_details"` // JSON string with detailed info
	DBType         string    `gorm:"column:db_type" json:"db_type"`
	DatabaseName   string    `gorm:"column:database_name" json:"database_name"`
	XCoordinate    float64   `gorm:"column:x_coordinate" json:"x_coordinate"`
	YCoordinate    float64   `gorm:"column:y_coordinate" json:"y_coordinate"`
	City           string    `gorm:"column:city" json:"city"`
	Country        string    `gorm:"column:country" json:"country"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (tr *TaskResult) TableName() string {
	return "task_result"
}
