package users

import (
	"lifelogger/server/domain/users/dto"
	model "lifelogger/server/domain/users/model"
)

type (
	CreateUserService interface {
		CreateUser(dto.CreateUserReqDto)
	}

	createUserServiceImpl struct {
		UserRepository model.UsersRepository
	}
)

func NewCreateUserService(UserRepository model.UsersRepository) CreateUserService {
	i := new(createUserServiceImpl)
	i.UserRepository = UserRepository
	return i
}

func (s *createUserServiceImpl) CreateUser(reqDto dto.CreateUserReqDto) {
	user := model.NewUser(reqDto.UserId, reqDto.UserEmail, reqDto.UserName)
	s.UserRepository.Save(user)
}
