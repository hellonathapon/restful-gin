package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hellonathapon/restful-gin/models"
	"github.com/hellonathapon/restful-gin/utils"
)

func GetIndex(c *gin.Context) {
	articles := models.GetAllArticles()

	utils.Render(c, gin.H{
		"title": "Articles",
		"payload": articles,
	}, "index.html")
}

func GetArticleByID(c *gin.Context) {
	//* Check if the article ID from client is valid
	//* by converting it from string to int
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := models.GetArticleByID(articleID); err == nil {
			utils.Render(c, gin.H{
				"title": article.Title,
				"payload": article,
			}, "article.html")
		}else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	}else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}