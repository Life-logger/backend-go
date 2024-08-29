package groups

import (
	"lifelogger/server/controller"

	"github.com/gofiber/fiber/v2"
)

func NewHomeGroup(f *fiber.App) {
	homeController := &controller.HomeController{}

	home := f.Group("/home")
	// home.Use(middleware.AuthMiddleware())
	home.Get("/home", homeController.Index) // 메인 페이지 렌더링

	// home.Get("/web/login.html", homeController.Index)

}
