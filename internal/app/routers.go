package app

import "github.com/gin-gonic/gin"

func InstallRouters(g gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "Not Found",
		})
	})

	g.Group("/v1")
	{
	}
	return nil
}
