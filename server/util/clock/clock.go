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

// 문자열로 주어진 기간 파싱
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

// 스톱워치

// Struct to hold the stopwatch data
type Stopwatch struct {
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

// Start the stopwatch and return the start time
var StartStopwatch = func() time.Time {
	return Now()
}

// Stop the stopwatch, calculate the duration, and return the end time and duration
var StopStopwatch = func(startTime time.Time) (endTime time.Time, duration time.Duration) {
	endTime = Now()
	duration = endTime.Sub(startTime)
	return endTime, duration
}

// Function to save stopwatch data to DB (Stub function - replace with actual DB logic)
var SaveStopwatchData = func(stopwatch Stopwatch) error {
	// Implement your DB saving logic here
	// For now, just print to console (for example purposes)
	fmt.Printf("Saving to DB: StartTime: %s, EndTime: %s, Duration: %s\n", Format(stopwatch.StartTime), Format(stopwatch.EndTime), stopwatch.Duration)
	return nil
}

// Example usage
func ExampleStopwatchUsage() {
	// Start the stopwatch
	startTime := StartStopwatch()

	// Simulate some processing time (sleep for 2 seconds)
	time.Sleep(2 * time.Second)

	// Stop the stopwatch
	endTime, duration := StopStopwatch(startTime)

	// Create the stopwatch data
	stopwatch := Stopwatch{
		StartTime: startTime,
		EndTime:   endTime,
		Duration:  duration,
	}

	// Save the stopwatch data to DB
	err := SaveStopwatchData(stopwatch)
	if err != nil {
		fmt.Println("Error saving stopwatch data:", err)
	}
}

// Function to reset the stopwatch and save the record
var ResetAndSaveStopwatchAtMidnight = func(stopwatch *Stopwatch) {
	for {
		now := Now()
		// Calculate time until next midnight
		midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(24 * time.Hour)
		timeUntilMidnight := time.Until(midnight) // 자정까지 남은 시간

		// Wait until midnight
		time.Sleep(timeUntilMidnight) // 자정까지 대기

		// Stop the stopwatch
		stopwatch.EndTime, stopwatch.Duration = StopStopwatch(stopwatch.StartTime)

		// Save the current stopwatch record
		err := SaveStopwatchData(*stopwatch)
		if err != nil {
			fmt.Println("Error saving stopwatch data:", err)
		}

		// Reset the stopwatch
		stopwatch.StartTime = StartStopwatch()
		stopwatch.EndTime = time.Time{}
		stopwatch.Duration = 0

		fmt.Println("자정이 되어 스톱워치가 초기화되었습니다. 다시 측정을 시작할게요.")
	}
}
