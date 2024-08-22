package router

import (
	"io"
	"log"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"taskbuddy.io/taskbuddy/router/groups"
)

func Initialize(multiWriter io.Writer) *fiber.App {
	f := fiber.New(fiber.Config{
		Immutable: true,
	})

	// CORS middleware
	f.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// logger middleware
	f.Use(logger.New(logger.Config{
		Format:     `[${time}] ${status} ${method} ${host}${path} ${latency}` + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Seoul",
		Output:     multiWriter,
	}))

	// panic recovery middleware
	f.Use(func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				debug.PrintStack()

			}
		}()
		return c.Next()
	})

	groups.NewHelloGroup(f)

	return f
}
