package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hellonathapon/restful-gin/models"

	_ "github.com/denisenkom/go-mssqldb"
)
var router *gin.Engine

func main() {
	router = gin.Default()

	// db pooling connection
	models.DbConn()

	// load HTML templates
	router.LoadHTMLGlob("templates/*")

	// parse routes
	initializeRoutes()

	router.Run()

}