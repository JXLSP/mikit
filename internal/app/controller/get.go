package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *TasksController) GetTask(ctx *gin.Context) {
	uuid := ctx.Params.ByName("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "参数是必须的",
		})
		return
	}

	task, err := ctrl.TUsecase.GetTaskByID(ctx, uuid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": task,
	})
}
