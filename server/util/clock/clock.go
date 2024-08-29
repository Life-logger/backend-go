package clock

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"
const startingDate = "2024-01-01T00:00:00+09:00"

var Now = func() time.Time {
	return time.Now()
}

var ParseDatetime = func(datetime string) time.Time { // 문자열로 주어진 시간을 time.Time(시간 형식)으로 변환
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return time.Time{}
		// panic(err)
	}
	return t.In(time.Now().Location())
}

var Format = func(t time.Time) string { // time.Time(시간)을 문자열로 변환
	emptyTimeString := "0001-01-01T00:00:00Z"
	result := t.Format(time.RFC3339)
	if result == emptyTimeString {
		return ""
	}
	return result
}

var AddDuration = func(t time.Time, years int, months int, days int) time.Time {
	// Add years and months first
	t = t.AddDate(years, months, 0)
	// Then add days
	t = t.AddDate(0, 0, days)
	return t
}

var ParseDuration = func(durationStr string) (years, months, days int, err error) {
	parts := strings.Split(durationStr, ":")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid duration format")
	}

	years, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid year value")
	}

	months, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid month value")
	}

	days, err = strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid day value")
	}

	return years, months, days, nil
}

var GetNextDay = func(datetime string) string {
	parsedTime, error := time.Parse(time.DateOnly, datetime)
	if error != nil {
		return ""
	}
	parsedTime = parsedTime.AddDate(0, 0, 1)
	return parsedTime.Format(time.DateOnly)
}

var CalculateSlot = func(time time.Time) int {
	hrs := time.Sub(ParseDatetime(startingDate)).Hours()
	hrs *= 2

	return int(hrs)
}
