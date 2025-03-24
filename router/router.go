package router

import (
	"net/http"
	"usermanager/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
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
