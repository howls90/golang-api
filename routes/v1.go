package routes

import (
	"api/controllers"
	"api/middlewares"
)

func CreateUrlMappingsV1() {

	v1 := Router.Group("/api/v1")
	{
		test := v1.Group("/tests")
		{
			testCtl := new(controllers.TestCtrl)
			test.GET("/hello", testCtl.Hello)
		}

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
			users.GET("/:user", usersCtl.Show)
			users.DELETE("/:user", usersCtl.Delete)
			users.POST("/", usersCtl.Store)

			users.POST("/file", usersCtl.Csv)
		}

		posts := v1.Group("/posts")
		{
			posts.Use(middlewares.Verify)

			postsCtl := new(controllers.PostsCtrl)
			posts.GET("/", postsCtl.Index)

			posts.GET("/:post", postsCtl.Show)
			posts.DELETE("/:post", postsCtl.Delete)
			posts.POST("/", postsCtl.Store)

			commentsCtl := new(controllers.CommentsCtrl)
			posts.GET("/:post/comments", commentsCtl.Index)
			posts.GET("/:post/comments/:comment", commentsCtl.Show)
			posts.DELETE("/:post/comments/:comment", commentsCtl.Delete)
			posts.POST("/:post/comments", commentsCtl.Store)
		}

	}
}
