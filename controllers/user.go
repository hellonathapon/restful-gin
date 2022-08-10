package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hellonathapon/restful-gin/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(c *gin.Context) {
	//* extract user param ID
	if ID, err := strconv.Atoi(c.Param("user_id")); err == nil {
		//* find ID in DB
		var user models.User
		models.DB.First(&user, ID)

		c.JSON(http.StatusOK, gin.H{"data": user})
	}else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func CreateUser(c *gin.Context) {
	//* extract user credentials and validate
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//* response error no credentials attempt
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//* check in DB.
	result := models.DB.Create(&user)

	//* Error on process
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully created"})
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User

	//* 1. bind user credentials to user model
	//* 2. use Update function to update the entire row
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a user object
	var value models.User

	if err := models.DB.Where("ID = ?", c.Param("user_id")).First(&value).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	
	value.Username = user.Username
	value.Email = user.Email
	value.Password = user.Password

	result := models.DB.Save(&value)


	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
	}
}

func DeleteUser(c *gin.Context) {
	var user models.User
	
	if err := models.DB.Where("ID = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this ID not found"})
		return
	}

	result := models.DB.Delete(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
	}
}