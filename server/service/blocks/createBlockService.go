package blocks

import (
	"lifelogger/server/domain/blocks/dto"
	model "lifelogger/server/domain/blocks/model"
)

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
