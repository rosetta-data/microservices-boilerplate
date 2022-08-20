// Code generated by mockery v2.14.0. DO NOT EDIT.

package service

import (
	context "context"
	domain "microservices-boilerplate/internal/serviceA/domain"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, item
func (_m *Service) Create(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	ret := _m.Called(ctx, item)

	var r0 *domain.ItemA
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ItemA) *domain.ItemA); ok {
		r0 = rf(ctx, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ItemA)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ItemA) error); ok {
		r1 = rf(ctx, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Service) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Service) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.ItemA
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.ItemA); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.ItemA)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByID provides a mock function with given fields: ctx, id
func (_m *Service) GetOneByID(ctx context.Context, id string) (*domain.ItemA, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.ItemA
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.ItemA); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ItemA)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, item
func (_m *Service) Update(ctx context.Context, id string, item *domain.ItemA) error {
	ret := _m.Called(ctx, id, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.ItemA) error); ok {
		r0 = rf(ctx, id, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
