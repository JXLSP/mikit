package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTask 获取任务详情
func (tc *TasksController) GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task, err := tc.TUsecase.GetTaskByID(c.Request.Context(), taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
