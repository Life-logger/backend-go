package groups

import (
	"lifelogger/server/controller"

	"github.com/gofiber/fiber/v2"
)

func NewUserGroup(f *fiber.App) {
	userController := controller.UserController{}

	group := f.Group("/user")
	group.Get("", userController.CreateUser)
}
