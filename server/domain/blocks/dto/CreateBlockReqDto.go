package dto

type CreateBlockReqDto struct {
	StartTime          int    `json:"startTime"`
	EndTime            int    `json:"endTime"`
	Duration           int    `json:"duration"`
	Color              string `json:"color"`
	BackgroundImageUrl string `json:"backgroundImageUrl"`
	BlockMemo          string `json:"blockMemo"`
}
