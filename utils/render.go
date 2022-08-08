package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, data gin.H, templateName string) {

	// check for accepted format and return respectively
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}