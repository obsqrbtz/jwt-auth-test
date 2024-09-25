package controllers

import (
	"jwt-auth-test/database"
	"jwt-auth-test/models"
	"jwt-auth-test/tokens"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateTokens(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.CreateUser(c.IP(), data["email"])

	claims := tokens.CreateUserClaims(user)

	accessToken, err := tokens.CreateAccessToken(&claims, os.Getenv("TOKEN_SECRET"))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not create access token",
		})
	}

	refreshToken, err := tokens.CreateRefreshToken(&user)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not create refresh token",
		})
	}

	user.RefreshToken = refreshToken

	database.DB.Create(&user)

	CreateAccessTokenCookie(c, accessToken)

	refreshTokenBase64 := EncodeRefreshToken(refreshToken)

	return c.JSON(fiber.Map{
		"refreshToken": refreshTokenBase64,
	})
}

func RefreshTokens(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	acccessToken, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return err
	}

	payload := acccessToken.Claims.(jwt.MapClaims)

	var user models.User

	tokenClientId := payload["client_id"].(string)

	database.DB.Where("id = ?", tokenClientId).First(&user)

	if user.ID == "" {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	refreshTokenDecoded, err := DecodeRefreshToken(data["refresh_token"])
	if err != nil {
		return err
	}
	if user.RefreshToken != string(refreshTokenDecoded) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "invalid refresh token",
		})
	}

	requestIp := c.IP()

	if user.IP != requestIp {
		SendEmail("Connected from new IP:"+requestIp, user.Email)
	}

	claims := tokens.CreateUserClaims(user)

	accessToken, err := tokens.CreateAccessToken(&claims, os.Getenv("TOKEN_SECRET"))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not create access token",
		})
	}

	refreshToken, err := tokens.CreateRefreshToken(&user)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not create refresh token",
		})
	}

	user.RefreshToken = refreshToken
	user.IP = requestIp

	database.DB.Save(user)

	CreateAccessTokenCookie(c, accessToken)

	refreshTokenBase64 := EncodeRefreshToken(refreshToken)

	return c.JSON(fiber.Map{
		"refreshToken": refreshTokenBase64,
	})
}

func GetUsers(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return err
	}

	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}
