package users

import (
	_ "github.com/go-sql-driver/mysql"
)

type User interface {
	UserId() int
	UserEmail() string
	UserName() string
}

type userImpl struct {
	userId    int
	userEmail string
	userName  string
}

// 새로운 사용자를 생성
func NewUser(userId int, userEmail string, userName string) User {
	i := new(userImpl)
	i.userId = userId
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

////////////////////////////////////////////////////////////////
