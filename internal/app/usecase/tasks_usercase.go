package usecase

import (
	"context"
	"mikit/internal/app/domain"
	"mikit/internal/pkg/models"
)

type tasksUsercase struct {
	repo domain.TaskRepostiory
}

func NewTasksUsecase(repo domain.TaskRepostiory) domain.TaskUsecase {
	return &tasksUsercase{
		repo: repo,
	}
}

func (t *tasksUsercase) CreateTask(ctx context.Context, task *models.Tasks) error {
	return nil
}

func (t *tasksUsercase) GetTaskByID(ctx context.Context, uuid string) (*models.Tasks, error) {
	return nil, nil
}

func (t *tasksUsercase) DeleteTaskByID(ctx context.Context, uuid string) error {
	return nil
}
