package domain

import (
	"context"

	"mikit/internal/pkg/models"
)

type TaskRepostiory interface {
	CreateTask(ctx context.Context, task *models.Tasks) error
	GetTaskByID(ctx context.Context, uuid string) (*models.Tasks, error)
	DeleteTaskByID(ctx context.Context, uuid string) error
}

type TaskUsecase interface {
	CreateTask(ctx context.Context, task *models.Tasks) error
	GetTaskByID(ctx context.Context, uuid string) (*models.Tasks, error)
	DeleteTaskByID(ctx context.Context, uuid string) error
}
