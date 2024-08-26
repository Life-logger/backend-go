package groups

import (
	"github.com/gofiber/fiber/v2"
	"taskbuddy.io/taskbuddy/server/controller"
)

func NewHelloGroup(f *fiber.App) {
	helloController := controller.HelloController{}

	group := f.Group("/hello")
	group.Get("", helloController.Hello)
}
