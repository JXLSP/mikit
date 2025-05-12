package controller

import (
	"mikit/internal/pkg/models"
	"mikit/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func (ctrl *TasksController) CreateTask(ctx *gin.Context) {
	var (
		req *types.CreatedTaskRequest
		tk  *models.Tasks
	)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误",
			"code": 400,
		})
		return
	}
	_ = copier.Copy(&tk, req)
	err := ctrl.TUsecase.CreateTask(ctx, tk)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "创建成功",
		"code": 200,
	})
}
