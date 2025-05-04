package models

import "time"

type Tasks struct {
	ID        int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	TaskID    string    `gorm:"column:task_id" json:"task_id"`
	TaskName  string    `gorm:"column:task_name;not null" json:"task_name"`
	TaskType  string    `gorm:"column:task_type" json:"task_type"`
	Targets   string    `gorm:"column:targets" json:"targets"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (t *Tasks) TableName() string {
	return "tasks"
}

type TaskResult struct {
	ID           int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	TaskID       string    `gorm:"column:task_id" json:"task_id"`
	Target       string    `gorm:"column:target" json:"target"`
	DBType       string    `gorm:"column:db_type" json:"db_type"`
	DatabaseName string    `gorm:"column:database_name" json:"database_name"`
	XCoordinate  float64   `gorm:"column:x_coordinate" json:"x_coordinate"`
	YCoordinate  float64   `gorm:"column:y_coordinate" json:"y_coordinate"`
	City         string    `gorm:"column:city" json:"city"`
	Country      string    `gorm:"column:country" json:"country"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (tr *TaskResult) TableName() string {
	return "task_result"
}
