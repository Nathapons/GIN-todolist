package controllers

import (
	"main/config"
	"main/forms"
	"main/middleware"
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a login endpoint
// @Summary Login
// @Description Perform user login
// @Accept json
// @Produce json
// @Router /login [post]
func Login(c *gin.Context) {
	var form forms.LoginForm
	var user, queryUser models.User

	if c.ShouldBind(&form) == nil {
		if err := config.GetDB().Where("username = ?", form.Username).First(&queryUser).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "error": "Unknown username"})
		} else if utils.CheckPasswordHash(user.Password, queryUser.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "unauthorized", "error": "invalid password"})
		} else if user.IsActive {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "unauthorized", "error": "user is not active"})
		} else {
			token := middleware.JwtSign(queryUser)
			c.JSON(http.StatusOK, gin.H{"message": "login complete", "token": token})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to bind"})
	}
}
