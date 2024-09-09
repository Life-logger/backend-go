package calendar

import (
	"context"
	"database/sql"
)

type (
	CalendarRepository interface {
		Save(Calendar)
	}
	calendarRepositoryImpl struct {
		*sql.Conn
		tx  *sql.Tx
		ctx context.Context
	}
)

func NewCalendarRepository(
	conn *sql.Conn,
	tx *sql.Tx,
	ctx context.Context,
) CalendarRepository {
	i := new(calendarRepositoryImpl)
	i.Conn = conn
	i.tx = tx
	i.ctx = ctx
	return i
}

func (s calendarRepositoryImpl) Save(calendar Calendar) {
	sqlStatement := `INSERT INTO CALENDAR
		(representative_block_id)
		VALUEW(?, ?, ?)`

	_, err := s.tx.ExecContext(s.ctx, sqlStatement, calendar.LogDate())
	if err != nil {
		panic(err.Error())
	}
}
