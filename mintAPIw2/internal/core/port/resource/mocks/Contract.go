// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "mintapi/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// Contract is an autogenerated mock type for the Contract type
type Contract struct {
	mock.Mock
}

// CreateEvent provides a mock function with given fields: ctx, event
func (_m *Contract) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	ret := _m.Called(ctx, event)

	var r0 *domain.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Event) (*domain.Event, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Event) *domain.Event); ok {
		r0 = rf(ctx, event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Event) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTicket provides a mock function with given fields: ctx, event
func (_m *Contract) CreateTicket(ctx context.Context, event *domain.Ticket) (*domain.Ticket, error) {
	ret := _m.Called(ctx, event)

	var r0 *domain.Ticket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Ticket) (*domain.Ticket, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Ticket) *domain.Ticket); ok {
		r0 = rf(ctx, event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Ticket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Ticket) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContract interface {
	mock.TestingT
	Cleanup(func())
}

// NewContract creates a new instance of Contract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContract(t mockConstructorTestingTNewContract) *Contract {
	mock := &Contract{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
