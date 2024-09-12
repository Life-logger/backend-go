package blocks

import (
	"lifelogger/server/domain/blocks/dto"
	model "lifelogger/server/domain/blocks/model"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05" // 시간 포맷 상수

var Now = func() time.Time {
	return time.Now() // 현재 시간을 반환하는 함수
}

type (
	CreateBlockService interface {
		CreateBlock(dto.CreateBlockReqDto)
	}

	createBlockServiceImpl struct {
		blockRepository model.BlocksRepository
	}
)

func NewCreateBlockService(blockRepository model.BlocksRepository) CreateBlockService {
	i := new(createBlockServiceImpl)
	i.blockRepository = blockRepository
	return i
}

func (s *createBlockServiceImpl) CreateBlock(reqDto dto.CreateBlockReqDto) {
	block := model.NewBlock(reqDto.StartTime, reqDto.EndTime, reqDto.Duration, reqDto.Color, reqDto.BackgroundImageUrl, reqDto.BlockMemo)
	s.blockRepository.Save(block)
}
