package main

import (
	"github.com/hellonathapon/restful-gin/controllers"
)

// parse and mapping routes
func initializeRoutes() {
	router.GET("/", controllers.GetIndex)

	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:user_id", controllers.GetUserByID)
	router.POST("/user", controllers.CreateUser)

	router.GET("/article/view/:article_id", controllers.GetArticleByID)
}