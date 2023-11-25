package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary			Get Users.
// @Description		Return users.
// @Tags			User
// @Router			/user [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

// @Summary			Create Users.
// @Description		Return user.
// @Tags			User
// @Param request body forms.User true "User data in JSON format"
// @Router			/user [post]
func CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create user"})
}

// @Summary			Update Users.
// @Description		Return new user.
// @Tags			User
// @Param id path int true "ID of models.User"
// @Param request body forms.User true "User data in JSON format"
// @Router			/user/{id} [put]
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update user"})
}

// @Summary			Update Users.
// @Description		Return new user.
// @Tags			User
// @Param id path int true "ID of models.User"
// @Router			/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete user"})
}
