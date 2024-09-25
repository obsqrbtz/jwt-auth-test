package controllers

import (
	"encoding/base64"
	"net/smtp"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SendEmail(message string, email string) error {
	from := os.Getenv("SMTP_SENDER")
	password := os.Getenv("SMTP_PASSWORD")

	to := []string{
		email,
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))

	return err
}

func CreateAccessTokenCookie(c *fiber.Ctx, accessToken string) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
}

func EncodeRefreshToken(refreshToken string) string {
	return (base64.StdEncoding.EncodeToString([]byte(refreshToken)))
}

func DecodeRefreshToken(refreshToken string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(refreshToken)
	return string(decoded), err
}
