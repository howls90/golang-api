package routes

import (
	"api/controllers"
	"api/middlewares"
)

func CreateUrlMappingsV2() {

	v1 := Router.Group("/api/v2")
	{
		auth := v1.Group("/auth")
		{
			loginCtl := new(controllers.AuthCtrl)
			auth.POST("/login", loginCtl.Login)
		}

		users := v1.Group("/users")
		{
			users.Use(middlewares.Verify)

			usersCtl := new(controllers.UsersCtrl)
			users.GET("/", usersCtl.Index)
			users.GET("/:id", usersCtl.Show)
			users.DELETE("/:id", usersCtl.Delete)
			users.POST("/", usersCtl.Store)
		}

	}
}
