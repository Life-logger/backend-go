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
	accessToken, _ := tokenGetter.GetAccessToken(code)    // Access Token, Refresh Token 가져오기
	nickname, err := tokenGetter.GetUserInfo(accessToken) // Nickname 가져오기
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Faild to retrieve user info")
	}
	c.Cookie(cookies.MakeCookies("accessToken", accessToken)) // Access Token을 쿠키에 저장
	return c.SendString("hi, " + nickname + "!!")
	// return c.SendString(refreshtoken)

}

func (a *LoginController) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Query("refreshToken")
	tokenGetter := util.NewTokenGetter()
	newAccessToken := tokenGetter.RefreshAccessToken(refreshToken)
	c.Cookie(cookies.MakeCookies("accessToken", newAccessToken))
	return c.SendStatus(200)
}
