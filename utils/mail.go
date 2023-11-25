package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendEmail(c *gin.Context, emailTo, subject, body string) {
	envMap, err := godotenv.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	EMAIL_HOST := envMap["EMAIL_HOST"]
	EMAIL_HOST_USER := envMap["EMAIL_HOST_USER"]
	EMAIL_PORT, _ := strconv.Atoi(envMap["EMAIL_PORT"])
	EMAIL_HOST_PASSWORD := envMap["EMAIL_HOST_PASSWORD"]

	dialer := gomail.NewDialer(EMAIL_HOST, EMAIL_PORT, EMAIL_HOST_USER, EMAIL_HOST_PASSWORD)

	sender := gomail.NewMessage()
	sender.SetHeader("From", EMAIL_HOST_USER)
	sender.SetHeader("To", emailTo)
	sender.SetHeader("Subject", subject)

	sender.SetBody("text/html", body)

	if err := dialer.DialAndSend(sender); err != nil {
		fmt.Println(err)
		return
	}
}
