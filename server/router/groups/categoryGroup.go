package groups

import (
	"lifelogger/server/controller"

	"github.com/gofiber/fiber/v2"
)

func NewCategoryGroup(f *fiber.App) {
	categoryController := controller.CategoryController{}

	group := f.Group("/category")
	group.Get("", categoryController.CreateCategory)
}
