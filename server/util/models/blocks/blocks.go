package blocks

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatalln(err)
	}
}

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

type Blocks struct {
	BlockID            int           `db:"block_id"`
	StartTime          time.Time     `db:"start_time"`
	EndTime            time.Time     `db:"end_time"`
	Duration           time.Duration `db:"duration"`
	Color              string        `db:"color"`
	BackgroundImageURL string        `db:"background_image_url"`
	BlockMemo          string        `db:"block_memo"`
	BlockPin           bool          `db:"block_pin"`
}

// Struct to hold the stopwatch duration data
// type Stopwatch struct {
// 	// StartTime Blocks.StartTime
// 	// EndTime   time.Time
// 	Duration time.Duration
// }

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

// 스톱워치 데이터 Blocks DB에 저장
// Function to save stopwatch data to DB (Stub function - replace with actual DB logic)
var SaveStopwatchData = func(blocks Blocks) error {
	// Implement your DB saving logic here
	if db == nil {
		return fmt.Errorf("DB connection is not initialized")
	}

	query := `
        INSERT INTO Blocks (start_time, end_time, color, background_image_url, block_memo, block_pin) 
        VALUES (?, ?, ?, ?, ?, ?)
    `

	_, err := db.Exec(query, blocks.StartTime, blocks.EndTime, blocks.Color, blocks.BackgroundImageURL, blocks.BlockMemo, blocks.BlockPin)
	if err != nil {
		return fmt.Errorf("데이터 삽입 실패: %v", err)
	}
	// For now, just print to console (for example purposes)
	fmt.Printf("Saving to DB: StartTime: %s, EndTime: %s, Duration: %s\n", Format(blocks.StartTime), Format(blocks.EndTime), blocks.Duration)
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
	stopwatch := Blocks{
		StartTime:          startTime,
		EndTime:            endTime,
		Duration:           duration,
		Color:              "#FFF",
		BackgroundImageURL: "",
		BlockMemo:          "",
		BlockPin:           false,
	}

	// Save the stopwatch data to DB
	err := SaveStopwatchData(stopwatch)
	if err != nil {
		fmt.Println("Error saving stopwatch data:", err)
	}
}

// Function to reset the stopwatch and save the record
var ResetAndSaveStopwatchAtMidnight = func(stopwatch *Blocks) {
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
