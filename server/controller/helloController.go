package controller

import (
	"lifelogger/server/di"
	"lifelogger/server/domain/hello/dto"
	"lifelogger/server/util/httpHelper"

	"github.com/gofiber/fiber/v2"
)

type HelloController struct{}

func (a *HelloController) Hello(c *fiber.Ctx) error {
	reqDto := new(dto.GetHelloReqDto)
	httpHelper.ParseQuery(c, reqDto)

	getHelloService := di.InjectGetHelloService()
	return c.SendString(getHelloService.GetHello(*reqDto))
}
