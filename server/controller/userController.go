package controller

import (
	"lifelogger/server/di"
	"lifelogger/server/domain/users/dto"
	"lifelogger/server/util/httpHelper"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (a *UserController) CreateUser(c *fiber.Ctx) error {
	reqDto := new(dto.CreateUserReqDto)
	httpHelper.ParseBody(c, reqDto)

	createUserService, cleanup := di.InjectCreateUserService()
	defer cleanup()
	createUserService.CreateUser(*reqDto)

	return c.SendStatus(200)
}
