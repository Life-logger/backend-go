package groups

import (
	"lifelogger/server/controller"

	"github.com/gofiber/fiber/v2"
)

func NewBlockGroup(f *fiber.App) {
	blockController := controller.BlockController{}

	group := f.Group("/block")
	group.Get("", blockController.CreateBlock)
}
