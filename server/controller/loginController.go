package controller

import (
	"github.com/gofiber/fiber/v2"
	"taskbuddy.io/taskbuddy/server/util"
	"taskbuddy.io/taskbuddy/server/util/cookies"
)

type LoginController struct{}

func (a *LoginController) Login(c *fiber.Ctx) error {
	code := c.Query("code") // 인가 코드
	tokenGetter := util.NewTokenGetter()
	accessToken, refreshToken := tokenGetter.GetAccessToken(code)
	c.Cookie(cookies.MakeCookies("accessToken", accessToken)) // 쿠키에 저장
	return c.SendString(refreshToken)                         // 프론트에 보냄
}

func (a *LoginController) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Query("refreshToken")
	tokenGetter := util.NewTokenGetter()
	newAccessToken := tokenGetter.GetNewAccessToken(refreshToken)
	c.Cookie(cookies.MakeCookies("accessToken", newAccessToken))
	return c.SendStatus(200)
}
