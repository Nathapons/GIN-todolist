package controllers

import (
	"fmt"
	"main/config"
	"main/forms"
	"main/models"
	"main/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary			Get Users.
// @Description		Return users.
// @Tags			User
// @Router			/user [get]
func GetUsers(c *gin.Context) {
	var user []forms.UserDetail
	config.GetDB().Model(&models.User{}).Find(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary			Get Users.
// @Description		Return users.
// @Tags			User
// @Param id path int true "ID of models.User"
// @Router			/user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user forms.UserDetail
	config.GetDB().Model(&models.User{}).First(&user, id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary			Create Users.
// @Description		Return user.
// @Tags			User
// @Param request body forms.User true "User data in JSON format"
// @Router			/user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var email string = user.Email
	if !strings.Contains(email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email data is not email"})
		return
	}

	if user.Username == user.Password {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Username is not same password"})
		return
	}

	password, _ := utils.HashPassword(user.Password)
	user.Password = password

	config.GetDB().Save(&user)
	go utils.SendEmail(c, user.Email, "Create User", "<h1>Create user complete</h1>")
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// @Summary			Update Users.
// @Description		Return new user.
// @Tags			User
// @Param id path int true "ID of models.User"
// @Param request body forms.User true "User data in JSON format"
// @Router			/user/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Update user id = %s", id)})
}

// @Summary			Update Users.
// @Description		Return new user.
// @Tags			User
// @Param id path int true "ID of models.User"
// @Router			/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete user"})
}
