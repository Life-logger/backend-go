package calendar

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("mysql", "user@tcp(127.0.0.1:3306)/init_db") // 수정
	if err != nil {
		log.Fatalln(err)
	}
}

const DateTimeFormat = "2006-01-02 15:04:05"

type Block struct { // DB에서 불러온 일정 블록 데이터를 담기 위한 구조체
	BlockID            int       `db:"block_id"`
	CategoryID         int       `db:"category_id"`
	StartTime          time.Time `db:"start_time"`
	EndTime            time.Time `db:"end_time"`
	Duration           time.Duration
	Color              string `db:"color"`
	BackgroundImageURL string `db:"background_image_url"`
	BlockMemo          string `db:"block_memo"`
	BlockPin           bool   `db:"block_pin"`
}

type Calendar struct { // 특정 날짜에 대한 대표 블록 ID를 저장하기 위한 구조체
	Date                  time.Time `db:"date"`
	RepresentativeBlockID int       `db:"representative_block_id"`
}

func GetBlocksByDate(c *fiber.Ctx) error { // 사용자가 요청한 날짜에 해당하는 블록 데이터를 가져오는 API handler
	date := c.Params("date") // 요청 경로에서 date 파라미터 추출

	var blocks []Block
	query := `
        SELECT b.background_image_url, b.block_memo
        FROM Blocks b
        JOIN Calendar c ON b.block_id = c.representative_block_id
        WHERE c.date = ?
    ` // Calendar의 특정 날짜(date)에 해당하는 블록을 Block에서 조회하여 가져온다.
	err := db.Select(&blocks, query, date) // 블록 데이터를 가져온다.
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// 성공적으로 데이터를 가져오면 블록 데이터를 JSON 형식으로 반환한다.
	return c.JSON(blocks)
}

func UpdateRepresentativeBlock(c *fiber.Ctx) error { // 특정 날짜의 대표 블록 ID 업데이트하는 API handler
	date := c.Params("date")
	var body struct {
		RepresentativeBlockID int `json:"representative_block_id"`
	}

	if err := c.BodyParser(&body); err != nil { // 요청 본문을 파싱하여 대표 블록 ID를 가져온다.
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	query := `
        UPDATE Calendar 
        SET representative_block_id = ? 
        WHERE date = ?
    ` // Calendar의 representative_block_id를 업데이트 한다.
	_, err := db.Exec(query, body.RepresentativeBlockID, date)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "updated successfully",
	})

	// test
}
