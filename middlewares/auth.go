package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) {
	fmt.Println("hello there :)")
}