// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	context "github.com/grab/ego/egofilters/http/security/context"
	mock "github.com/stretchr/testify/mock"
)

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// OnComplete provides a mock function with given fields: _a0
func (_m *Callbacks) OnComplete(_a0 context.AuthResponse) {
	_m.Called(_a0)
}
