package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	// load HTML templates
	router.LoadHTMLGlob("templates/*")

	// parse routes
	initializeRoutes()

	router.Run()

}