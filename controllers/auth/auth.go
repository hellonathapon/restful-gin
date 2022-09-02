package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellonathapon/restful-gin/models"
	"gorm.io/gorm"
)

type UserLogin struct {
	Username string `json:"username"`;
	Password string `json:"password"`;
}

func Login(c *gin.Context) {
	var userCredentials UserLogin;
	//* extract and validate login credentials
	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	//* perform finding by username
	result := models.DB.First(&user, "username = ?", userCredentials.Username)

	//* if username not found in DB.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Username not found"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Register(c *gin.Context) {
	fmt.Println("DEBUG: Register handler");
}