package groups

import (
	"lifelogger/middleware"
	"lifelogger/server/controller"

	"github.com/gofiber/fiber/v2"
)

func NewHomeGroup(f *fiber.App) {
	homeController := &controller.HomeController{}

	home := f.Group("/home")
	home.Use(middleware.AuthMiddleware)
	home.Get("/web/login.html", homeController.Index)

}
