package main

import (
	"github.com/hellonathapon/restful-gin/routes"
)

func initializeRoutes() {
	router.GET("/", routes.HandleIndexRoute)

	router.GET("/article/view/:article_id", routes.HandleGetArticleByID)
}