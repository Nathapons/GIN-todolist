package controllers

import (
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

// @Summary			Get Users.
// @Description		Return users.
// @Tags			User
// @Router			/user [get]
func GetUsers(c *gin.Context) {
	m := gomail.NewMessage()
	m.SetHeader("From", "nuthaponm79@gmail.com")
	m.SetHeader("To", "nuthaponm79@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
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

// @Summary			Send Email.
// @Tags			Test
// @Router			/send-email [get]
func TestEmail(c *gin.Context) {
	go utils.SendEmail(c, "nuthaponm79@gmail.com", "Test send mail", "<h1>This is blog application</h1>")
	c.String(http.StatusOK, "Email sent successfully!")
}
