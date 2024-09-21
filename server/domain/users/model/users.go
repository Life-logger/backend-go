package users

import (
	_ "github.com/go-sql-driver/mysql"
	"lifelogger/server/domain/users/dto"
)

type User interface {
	UserId() int
	UserEmail() string
	UserName() string
	ToDto() dto.UserDto
}

type userImpl struct {
	userId    int
	userEmail string
	userName  string
}

// 새로운 사용자를 생성
func NewUser(userEmail string, userName string) User {
	i := new(userImpl)
	i.userEmail = userEmail
	i.userName = userName
	return i
}

func (c userImpl) UserId() int {
	return c.userId
}

func (c userImpl) UserEmail() string {
	return c.userEmail
}

func (c userImpl) UserName() string {
	return c.userName
}

func (c userImpl) ToDto() dto.UserDto {
	return dto.UserDto{
		UserEmail: c.userEmail,
		UserName:  c.userName,
	}

}

////////////////////////////////////////////////////////////////
