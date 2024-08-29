package groups

import (
	// "lifelogger/server/controller"
	"lifelogger/middleware"
	"lifelogger/server/controller"

	"github.com/gofiber/fiber/v2"
)

func NewLoginGroup(f *fiber.App) {
	loginController := controller.LoginController{}

	group := f.Group("/login")
    group.Use(middleware.AuthMiddleware)

	group.Get("/callback", loginController.Login)
	group.Get("/refresh", loginController.RefreshToken)
}
