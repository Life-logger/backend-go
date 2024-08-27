package groups

import (
	"github.com/gofiber/fiber/v2"
	"taskbuddy.io/taskbuddy/server/controller"
)

func NewLoginGroup(f *fiber.App) {
	loginController := controller.LoginController{}

	group := f.Group("/login")
	group.Get("/callback", loginController.Login)
	group.Get("/refresh", loginController.RefreshToken)
}
