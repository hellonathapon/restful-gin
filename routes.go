package main

import "github.com/hellonathapon/restful-gin/routes"

func initializeRoutes() {
	router.GET("/", routes.HandleIndexRoute)
}