package controller

import (
	"github.com/gofiber/fiber/v2"
)

// HomeController는 홈 페이지 요청을 처리하는 구조체입니다.
type HomeController struct{}

// Index는 메인 페이지를 렌더링합니다.
func (hc *HomeController) Index(c *fiber.Ctx) error {
	// 액세스 토큰이 없으면 로그인 페이지로 리디렉션
	// 메인 페이지의 HTML 파일을 렌더링합니다.
	return c.SendFile("./web/index.html")
}
