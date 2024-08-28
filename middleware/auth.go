package middleware

import (
	"lifelogger/server/util"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware는 사용자의 로그인 상태를 확인하는 미들웨어입니다.
func AuthMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("accessToken")
	log.Printf("Access Token: %s", accessToken)
	if accessToken == "" {
		// 쿠키에 액세스 토큰이 없으면 로그인 페이지로 리다이렉트
		log.Println("No access token found, redirecting to login")
		return c.Redirect("/login")
	}

	tokenGetter := util.NewTokenGetter()
	_, _, err := tokenGetter.GetUserInfo(accessToken)
	if err != nil {
		// 액세스 토큰이 유효하지 않으면 로그인 페이지로 리다이렉트
		return c.Redirect("/login")
	}

	return c.Next()
}
