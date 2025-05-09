// Code generated by mockery v2.46.3. DO NOT EDIT.

package oidcmocks

import (
	oidc "github.com/kyma-project/test-infra/pkg/oidc"
	mock "github.com/stretchr/testify/mock"
)

// MockTokenProcessorOption is an autogenerated mock type for the TokenProcessorOption type
type MockTokenProcessorOption struct {
	mock.Mock
}

type MockTokenProcessorOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTokenProcessorOption) EXPECT() *MockTokenProcessorOption_Expecter {
	return &MockTokenProcessorOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockTokenProcessorOption) Execute(_a0 *oidc.TokenProcessor) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*oidc.TokenProcessor) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTokenProcessorOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockTokenProcessorOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *oidc.TokenProcessor
func (_e *MockTokenProcessorOption_Expecter) Execute(_a0 interface{}) *MockTokenProcessorOption_Execute_Call {
	return &MockTokenProcessorOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockTokenProcessorOption_Execute_Call) Run(run func(_a0 *oidc.TokenProcessor)) *MockTokenProcessorOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*oidc.TokenProcessor))
	})
	return _c
}

func (_c *MockTokenProcessorOption_Execute_Call) Return(_a0 error) *MockTokenProcessorOption_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTokenProcessorOption_Execute_Call) RunAndReturn(run func(*oidc.TokenProcessor) error) *MockTokenProcessorOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTokenProcessorOption creates a new instance of MockTokenProcessorOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTokenProcessorOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTokenProcessorOption {
	mock := &MockTokenProcessorOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
