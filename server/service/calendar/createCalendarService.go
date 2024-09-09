package calendar

import (
	"lifelogger/server/domain/calendar/dto"
	model "lifelogger/server/domain/calendar/model"
)

type (
	CreateCalendarService interface {
		CreateCalendar(dto.CreateCalendarReqDto)
	}

	createCalendarServiceImpl struct {
		calendarRepository model.CalendarRepository
	}
)

func NewCreateCalendarService(calendarRepository model.CalendarRepository) CreateCalendarService {
	i := new(createCalendarServiceImpl)
	i.calendarRepository = calendarRepository
	return i
}

func (s *createCalendarServiceImpl) CreateCalendar(reqDto dto.CreateCalendarReqDto) {
	calendar := model.NewCalendar(reqDto.RepresentativeBlockId)
	s.calendarRepository.Save(calendar)
}
