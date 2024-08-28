package controller

import (
	"lifelogger/server/util"
	"lifelogger/server/util/cookies"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct{}

func (a *LoginController) Login(c *fiber.Ctx) error {
	code := c.Query("code")
	tokenGetter := util.NewTokenGetter()
	accessToken, refreshToken := tokenGetter.GetAccessToken(code) // Access Token, Refresh Token 가져오기
	nickname, email, err := tokenGetter.GetUserInfo(accessToken)  // Nickname 가져오기
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Faild to retrieve user info")
	}
	c.Cookie(cookies.MakeCookies("refreshToken", refreshToken)) // Refresh Token을 쿠키에 저장
	c.Cookie(cookies.MakeCookies("accessToken", accessToken))   // Access Token을 쿠키에 저장
	return c.SendString("hi, " + nickname + "!! \nYour email: " + email + "\nYour refreshToken: " + refreshToken)
	// return c.SendString(refreshtoken)

}

func (a *LoginController) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Query("refreshToken")
	tokenGetter := util.NewTokenGetter()
	newAccessToken := tokenGetter.GetNewAccessToken(refreshToken)
	c.Cookie(cookies.MakeCookies("accessToken", newAccessToken))
	return c.SendStatus(200)
}
