package middleware

import (
	"fmt"
	"lifelogger/server/util"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware는 사용자의 로그인 상태를 확인하는 미들웨어입니다.
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !isExcludedUrl(c) {
			accessToken := c.Cookies("accessToken")
			log.Printf("Access Token: %s", accessToken)
			if accessToken == "" {
				// 쿠키에 액세스 토큰이 없으면 로그인 페이지로 리다이렉트
				log.Println("No access token found")
				panic("No access token found 안녕")
			}

			tokenGetter := util.NewTokenGetter()
			_, _, err := tokenGetter.GetUserInfo(accessToken)
			if err != nil {
				// 액세스 토큰이 유효하지 않으면 로그인 페이지로 리다이렉트
				panic("No access token found")
			}
		}
		return c.Next()
	}
}

func isExcludedUrl(c *fiber.Ctx) bool {
	excludedUri := []string{"user"}
	for _, uri := range excludedUri {
		fmt.Println(string(c.Request().RequestURI()), uri)
		if strings.Contains(string(c.Request().RequestURI()), uri) {
			return true
		}
	}
	return false
}
