// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	dto "lifelogger/server/domain/blocks/dto"

	mock "github.com/stretchr/testify/mock"
)

// CreateBlockService is an autogenerated mock type for the CreateBlockService type
type CreateBlockService struct {
	mock.Mock
}

// CreateBlock provides a mock function with given fields: _a0
func (_m *CreateBlockService) CreateBlock(_a0 dto.CreateBlockReqDto) {
	_m.Called(_a0)
}

// NewCreateBlockService creates a new instance of CreateBlockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateBlockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateBlockService {
	mock := &CreateBlockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
