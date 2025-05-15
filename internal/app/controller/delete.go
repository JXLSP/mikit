package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteTask 删除任务
func (tc *TasksController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if err := tc.TUsecase.DeleteTaskByID(c.Request.Context(), taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "任务删除成功"})
}
