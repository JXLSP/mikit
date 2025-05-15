package controller

import (
	"mikit/internal/app/domain"
	"mikit/internal/app/repository"
	"mikit/internal/app/store"
	"mikit/internal/app/usecase"
)

type TasksController struct {
	TUsecase domain.TaskUsecase
}

// NewTasksController 创建任务控制器实例，注入Store依赖
func NewTasksController(ds *store.Store) *TasksController {
	// 创建Repository实例，注入Store依赖
	repo := repository.NewTaskRepository(ds)
	// 创建UseCase实例，注入Repository依赖
	return &TasksController{
		TUsecase: usecase.NewTasksUsecase(repo),
	}
}
