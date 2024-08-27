package controller

import (
	"github.com/gofiber/fiber/v2"
	"taskbuddy.io/taskbuddy/server/util"
	"taskbuddy.io/taskbuddy/server/util/cookies"
)

type LoginController struct{}

func (a *LoginController) Login(c *fiber.Ctx) error {
	code := c.Query("code")
	tokenGetter := util.NewTokenGetter()
	accessToken, refreshToken := tokenGetter.GetAccessToken(code)
	c.Cookie(cookies.MakeCookies("accessToken", accessToken))
	return c.SendString(refreshToken)
}

func (a *LoginController) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Query("refreshToken")
	tokenGetter := util.NewTokenGetter()
	newAccessToken := tokenGetter.RefreshAccessToken(refreshToken)
	c.Cookie(cookies.MakeCookies("accessToken", newAccessToken))
	return c.SendStatus(200)
}
