package controller

import (
	"lifelogger/server/di"
	"lifelogger/server/domain/calendar/dto"
	"lifelogger/server/util/httpHelper"

	"github.com/gofiber/fiber/v2"
)

type CalendarController struct{}

func (a *CalendarController) CreateCalendar(c *fiber.Ctx) error {
	reqDto := new(dto.CreateCalendarReqDto)
	httpHelper.ParseBody(c, reqDto)

	createCalendarService, cleanup := di.InjectCreateCalendarService()
	defer cleanup()
	createCalendarService.CreateCalendar(*reqDto)

	return c.SendStatus(200)
}
