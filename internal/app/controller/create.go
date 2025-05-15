package controller

import (
	"mikit/internal/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (tc *TasksController) CreateTask(c *gin.Context) {
	var task models.Tasks
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TUsecase.CreateTask(c.Request.Context(), &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task_id": task.TaskID})
}
