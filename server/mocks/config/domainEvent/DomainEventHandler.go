// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	domainEvent "lifelogger/server/config/domainEvent"

	mock "github.com/stretchr/testify/mock"
)

// DomainEventHandler is an autogenerated mock type for the DomainEventHandler type
type DomainEventHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: event
func (_m *DomainEventHandler) Handle(event domainEvent.DomainEvent) {
	_m.Called(event)
}

// NewDomainEventHandler creates a new instance of DomainEventHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDomainEventHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *DomainEventHandler {
	mock := &DomainEventHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
