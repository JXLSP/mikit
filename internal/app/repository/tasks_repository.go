package repository

import (
	"context"

	"mikit/internal/app/domain"
	"mikit/internal/app/store"
	"mikit/internal/pkg/models"
)

type tasksRepository struct {
	store *store.Store
}

// NewTaskRepository 创建任务仓库实例，注入Store依赖
func NewTaskRepository(ds *store.Store) domain.TaskRepostiory {
	return &tasksRepository{
		store: ds,
	}
}

// CreateTask 创建任务
func (t *tasksRepository) CreateTask(ctx context.Context, task *models.Tasks) error {
	// 使用Store的DB方法获取数据库连接
	return t.store.DB().Create(task).Error
}

// GetTaskByID 根据ID获取任务
func (t *tasksRepository) GetTaskByID(ctx context.Context, uuid string) (*models.Tasks, error) {
	var task models.Tasks
	err := t.store.DB().Where("task_id = ?", uuid).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTaskByID 根据ID删除任务
func (t *tasksRepository) DeleteTaskByID(ctx context.Context, uuid string) error {
	return t.store.DB().Where("task_id = ?", uuid).Delete(&models.Tasks{}).Error
}
