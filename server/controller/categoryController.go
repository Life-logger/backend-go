package controller

import (
	"lifelogger/server/di"
	"lifelogger/server/domain/categories/dto"
	"lifelogger/server/util/httpHelper"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct{}

func (a *CategoryController) CreateCategory(c *fiber.Ctx) error {
	reqDto := new(dto.CreateCategoryReqDto)
	httpHelper.ParseBody(c, reqDto)

	createCategoryService, cleanup := di.InjectCreateCategoryService()
	defer cleanup()
	createCategoryService.CreateCategory(*reqDto)

	return c.SendStatus(200)
}
