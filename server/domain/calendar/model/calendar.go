package calendar

type Calendar interface {
	LogDate() string
	RepresentativeBlockId() int
}

type calendarImpl struct {
	logDate               string
	representativeBlockId int
}

func NewCalendar(representativeBlockId int) Calendar {
	i := new(calendarImpl)
	i.representativeBlockId = representativeBlockId
	return i
}

func (c calendarImpl) LogDate() string {
	return c.logDate
}

func (c calendarImpl) RepresentativeBlockId() int {
	return c.representativeBlockId
}
