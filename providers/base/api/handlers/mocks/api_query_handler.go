// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/slinky/providers/types"
)

// APIQueryHandler is an autogenerated mock type for the APIQueryHandler type
type APIQueryHandler[K comparable, V interface{}] struct {
	mock.Mock
}

// Query provides a mock function with given fields: ctx, ids, responseCh
func (_m *APIQueryHandler[K, V]) Query(ctx context.Context, ids []K, responseCh chan<- types.GetResponse[K, V]) {
	_m.Called(ctx, ids, responseCh)
}

// NewAPIQueryHandler creates a new instance of APIQueryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIQueryHandler[K comparable, V interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *APIQueryHandler[K, V] {
	mock := &APIQueryHandler[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}