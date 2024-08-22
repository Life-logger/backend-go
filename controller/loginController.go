package controller

import (
	"github.com/gofiber/fiber/v2"
	"taskbuddy.io/taskbuddy/util"
)

type LoginController struct{}

func (a *LoginController) Login(c *fiber.Ctx) error {
	code := c.Query("code")
	tokenGetter := util.NewTokenGetter()
	token := tokenGetter.GetAccessToken(code)
	return c.SendString(token)
}
