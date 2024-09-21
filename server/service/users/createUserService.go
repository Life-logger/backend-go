package users

import (
	"fmt"
	"lifelogger/server/domain/users/dto"
	model "lifelogger/server/domain/users/model"
)

type (
	CreateUserService interface {
		CreateUser(dto.CreateUserReqDto) dto.UserDto
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

func (s *createUserServiceImpl) CreateUser(reqDto dto.CreateUserReqDto) dto.UserDto {
	fmt.Println(reqDto, "reqdto")
	user := model.NewUser(reqDto.UserEmail, reqDto.UserName)
	s.UserRepository.Save(user)
	fmt.Println(user, "user")
	return user.ToDto()
}
