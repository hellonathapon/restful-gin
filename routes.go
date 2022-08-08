package main

import (
	"github.com/hellonathapon/restful-gin/controllers"
)

// parse and mapping routes
func initializeRoutes() {
	router.GET("/", controllers.GetIndex)

	router.GET("/user", controllers.GetUsers)

	router.GET("/article/view/:article_id", controllers.GetArticleByID)
}