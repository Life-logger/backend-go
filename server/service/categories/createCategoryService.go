package categories

import (
	"lifelogger/server/domain/categories/dto"
	model "lifelogger/server/domain/categories/model"
)

type (
	CreateCategoryService interface {
		CreateCategory(dto.CreateCategoryReqDto)
	}

	createCategoryServiceImpl struct {
		categoryRepository model.CategoriesRepository
	}
)

func NewCreateCategoryService(categoryRepository model.CategoriesRepository) CreateCategoryService {
	i := new(createCategoryServiceImpl)
	i.categoryRepository = categoryRepository
	return i
}

func (s *createCategoryServiceImpl) CreateCategory(reqDto dto.CreateCategoryReqDto) {
	category := model.NewCategory(reqDto.UserEmail, reqDto.Color, reqDto.CategoryTitle)
	s.categoryRepository.Save(category)
}
