package controllers

import (
	"fmt"
	"main/config"
	"main/forms"
	"main/models"
	"main/utils"
	"net/http"
	"strconv"
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
// @Success 201 {object} models.User "Successful response"
// Failure 400 {object} gin.H  "error response"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is not same password"})
		return
	}

	password, _ := utils.HashPassword(user.Password)
	user.Password = password

	config.GetDB().Save(&user)
	userId := fmt.Sprintf("%d", user.Id)
	activationLink := fmt.Sprintf("http://localhost:8000/api/v1/user/activate/%s", userId)
	body := fmt.Sprintf("<h1>Create user complete</h1><p><a href=\"%s\">Click here</a> to activate</p>", activationLink)
	go utils.SendEmail(c, user.Email, "Create User", body)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// @Summary			Update Users.
// @Description		Return new user.
// @Tags			User
// @Param id path int true "ID of models.User"
// @Param request body forms.User true "User data in JSON format"
// @Router			/user/{id} [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var userData forms.UserDetail
	var myUser models.User
	var err error

	if err = c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err = config.GetDB().First(&myUser, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	myUser.FirstName = userData.FirstName
	myUser.LastName = userData.LastName
	config.GetDB().Save(&myUser)

	c.JSON(http.StatusOK, gin.H{"user": myUser})
}

// @Summary			Update Users.
// @Description		Return new user.
// @Tags			User
// @Param id path int true "ID of models.User"
// @Router			/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	config.GetDB().Delete(&models.User{}, id)
	c.Status(http.StatusNoContent)
}


func ActiveUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	config.GetDB().First(&user, id)
	user.IsActive = true
	config.GetDB().Save(&user)

	htmlContent := fmt.Sprintf("<h1>Active user %s complete!</h1>", user.GetFullName())
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, htmlContent)
}