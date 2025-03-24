package router

import (
	"net/http"
	"usermanager/controller"
	"usermanager/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	r.GET("/hi", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	user := r.Group("/user")
	{
		user.GET("/info/:id", controller.UserController{}.GetUser)
		user.POST("/put", controller.UserController{}.PutUser)
	}
	return r
}
