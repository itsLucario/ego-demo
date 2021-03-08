// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	volatile "github.com/grab/ego/ego/src/go/volatile"
	mock "github.com/stretchr/testify/mock"
)

// HeaderMap is an autogenerated mock type for the HeaderMap type
type HeaderMap struct {
	mock.Mock
}

// AddCopy provides a mock function with given fields: name, value
func (_m *HeaderMap) AddCopy(name string, value string) {
	_m.Called(name, value)
}

// AppendCopy provides a mock function with given fields: name, value
func (_m *HeaderMap) AppendCopy(name string, value string) {
	_m.Called(name, value)
}

// Get provides a mock function with given fields: name
func (_m *HeaderMap) Get(name string) volatile.String {
	ret := _m.Called(name)

	var r0 volatile.String
	if rf, ok := ret.Get(0).(func(string) volatile.String); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(volatile.String)
	}

	return r0
}

// Remove provides a mock function with given fields: name
func (_m *HeaderMap) Remove(name string) {
	_m.Called(name)
}

// SetCopy provides a mock function with given fields: name, value
func (_m *HeaderMap) SetCopy(name string, value string) {
	_m.Called(name, value)
}
