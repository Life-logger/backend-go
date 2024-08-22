package controller

import (
	"github.com/gofiber/fiber/v2"
	"taskbuddy.io/taskbuddy/di"
	"taskbuddy.io/taskbuddy/domain/hello/dto"
	"taskbuddy.io/taskbuddy/util/httpHelper"
)

type HelloController struct{}

func (a *HelloController) Hello(c *fiber.Ctx) error {
	reqDto := new(dto.GetHelloReqDto)
	httpHelper.ParseQuery(c, reqDto)

	getHelloService := di.InjectGetHelloService()
	return c.SendString(getHelloService.GetHello(*reqDto))
}
