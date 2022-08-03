package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hellonathapon/restful-gin/models"
	"github.com/hellonathapon/restful-gin/utils"
)

func HandleIndexRoute(c *gin.Context) {
	articles := models.GetAllArticles()

	utils.Render(c, gin.H{
		"title": "Articles",
		"payload": articles,
	}, "index.html")
}