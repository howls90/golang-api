package routes

import (
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func Init() {
	Router.Use(middlewares.ErrorHandler)

	CreateUrlMappingsV1()
	CreateUrlMappingsV2()

	Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
}
