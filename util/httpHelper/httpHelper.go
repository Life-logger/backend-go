package httpHelper

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseQuery(c *fiber.Ctx, target_dto interface{}) {
	c.QueryParser(target_dto)

	validate := validator.New()
	if err := validate.Struct(target_dto); err != nil {
		log.Println(err.Error())
		panic("[Fail] required parameters empty")
	}
}

func ParseBody(c *fiber.Ctx, target_dto interface{}) {
	if err := c.BodyParser(target_dto); err != nil {
		log.Println(err.Error())
		panic("[Invalid] failed to parse body")
	}

	validate := validator.New()
	if err := validate.Struct(target_dto); err != nil {
		log.Println(err.Error())
		panic("[Fail] required body data empty")
	}
}

func SetHeaderForFileDownload(c *fiber.Ctx, file_name string) {
	c.Set("Access-Control-Expose-Headers", "Content-Disposition,Content-Type,Content-Transfer-Encoding")
	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename="+file_name)
	c.Set("Content-Transfer-Encoding", "binary")
}
