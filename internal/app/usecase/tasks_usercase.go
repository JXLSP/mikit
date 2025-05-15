package usecase

import (
	"context"

	"mikit/internal/app/domain"
	"mikit/internal/pkg/models"
)

type tasksUsercase struct {
	repo domain.TaskRepostiory
}

// NewTasksUsecase 创建任务用例实例，注入Repository依赖
func NewTasksUsecase(repo domain.TaskRepostiory) domain.TaskUsecase {
	return &tasksUsercase{
		repo: repo,
	}
}

// CreateTask 创建任务
func (t *tasksUsercase) CreateTask(ctx context.Context, task *models.Tasks) error {
	// 调用Repository层的方法创建任务
	return t.repo.CreateTask(ctx, task)
}

// GetTaskByID 根据ID获取任务
func (t *tasksUsercase) GetTaskByID(ctx context.Context, uuid string) (*models.Tasks, error) {
	// 调用Repository层的方法获取任务
	return t.repo.GetTaskByID(ctx, uuid)
}

// DeleteTaskByID 根据ID删除任务
func (t *tasksUsercase) DeleteTaskByID(ctx context.Context, uuid string) error {
	// 调用Repository层的方法删除任务
	return t.repo.DeleteTaskByID(ctx, uuid)
}
