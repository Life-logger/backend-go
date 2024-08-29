package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func MakeCookies(key, value string) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Path = "/"
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(time.Hour * 24)
	return cookie
}
