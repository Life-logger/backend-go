package groups

import (
	"lifelogger/server/controller"

	"github.com/gofiber/fiber/v2"
)

func NewHelloGroup(f *fiber.App) {
	helloController := controller.HelloController{}

	group := f.Group("/hello")
	group.Get("", helloController.Hello)
}
