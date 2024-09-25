package routes

import (
	"jwt-auth-test/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/auth")
	api.Get("/get-users", controllers.GetUsers)
	api.Post("/create-tokens", controllers.CreateTokens)
	api.Post("/refresh-tokens", controllers.RefreshTokens)

}
