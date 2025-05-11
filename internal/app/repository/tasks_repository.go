package repository

import (
	"context"
	"mikit/internal/app/domain"
	"mikit/internal/app/store"
	"mikit/internal/pkg/models"
)

type tasksRepository struct {
	store store.Store
}

func NewTaskRepository(ds store.Store) domain.TaskRepostiory {
	return &tasksRepository{
		store: ds,
	}
}

func (t *tasksRepository) CreateTask(ctx context.Context, task *models.Tasks) error {
	return nil
}

func (t *tasksRepository) GetTaskByID(ctx context.Context, uuid string) (*models.Tasks, error) {
	return nil, nil
}

func (t *tasksRepository) DeleteTaskByID(ctx context.Context, uuid string) error {
	return nil
}
