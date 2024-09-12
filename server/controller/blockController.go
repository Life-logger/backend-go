package controller

import (
	"lifelogger/server/di"
	"lifelogger/server/domain/blocks/dto"
	"lifelogger/server/util/httpHelper"

	"github.com/gofiber/fiber/v2"
)

type BlockController struct{}

func (a *BlockController) CreateBlock(c *fiber.Ctx) error {
	reqDto := new(dto.CreateBlockReqDto)
	httpHelper.ParseBody(c, reqDto)

	createBlockService, cleanup := di.InjectCreateBlockService()
	defer cleanup()
	createBlockService.CreateBlock(*reqDto)

	return c.SendStatus(200)
}
