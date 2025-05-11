package controller

import (
	"mikit/internal/app/domain"
	"mikit/internal/app/repository"
	"mikit/internal/app/store"
	"mikit/internal/app/usecase"
)

type TasksController struct {
	TUsecase domain.TaskUsecase
	ds       store.Store
}

func NewTasksController(ds store.Store) *TasksController {
	repo := repository.NewTaskRepository(ds)
	return &TasksController{
		TUsecase: usecase.NewTasksUsecase(repo),
		ds:       ds,
	}
}
