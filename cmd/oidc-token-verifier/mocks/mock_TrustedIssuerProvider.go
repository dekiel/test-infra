// Code generated by mockery v2.46.3. DO NOT EDIT.

package mainmocks

import (
	oidc "github.com/kyma-project/test-infra/pkg/oidc"
	mock "github.com/stretchr/testify/mock"
)

// MockTrustedIssuerProvider is an autogenerated mock type for the TrustedIssuerProvider type
type MockTrustedIssuerProvider struct {
	mock.Mock
}

type MockTrustedIssuerProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTrustedIssuerProvider) EXPECT() *MockTrustedIssuerProvider_Expecter {
	return &MockTrustedIssuerProvider_Expecter{mock: &_m.Mock}
}

// GetIssuer provides a mock function with given fields:
func (_m *MockTrustedIssuerProvider) GetIssuer() oidc.Issuer {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIssuer")
	}

	var r0 oidc.Issuer
	if rf, ok := ret.Get(0).(func() oidc.Issuer); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(oidc.Issuer)
	}

	return r0
}

// MockTrustedIssuerProvider_GetIssuer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIssuer'
type MockTrustedIssuerProvider_GetIssuer_Call struct {
	*mock.Call
}

// GetIssuer is a helper method to define mock.On call
func (_e *MockTrustedIssuerProvider_Expecter) GetIssuer() *MockTrustedIssuerProvider_GetIssuer_Call {
	return &MockTrustedIssuerProvider_GetIssuer_Call{Call: _e.mock.On("GetIssuer")}
}

func (_c *MockTrustedIssuerProvider_GetIssuer_Call) Run(run func()) *MockTrustedIssuerProvider_GetIssuer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTrustedIssuerProvider_GetIssuer_Call) Return(_a0 oidc.Issuer) *MockTrustedIssuerProvider_GetIssuer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrustedIssuerProvider_GetIssuer_Call) RunAndReturn(run func() oidc.Issuer) *MockTrustedIssuerProvider_GetIssuer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTrustedIssuerProvider creates a new instance of MockTrustedIssuerProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTrustedIssuerProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTrustedIssuerProvider {
	mock := &MockTrustedIssuerProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}