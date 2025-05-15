package app

import (
	"mikit/internal/app/controller"
	"mikit/internal/app/store"

	"github.com/gin-gonic/gin"
)

func InstallRouters(g gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "Not Found",
		})
	})

	t := controller.NewTasksController(&store.Store{})

	g.Group("/v1")
	{
		g.POST("/create", t.CreateTask)
	}
	return nil
}
