// Code generated by mockery v2.39.1. DO NOT EDIT.

package pipelinesmocks

import (
	context "context"

	build "github.com/microsoft/azure-devops-go-api/azuredevops/v7/build"

	mock "github.com/stretchr/testify/mock"
)

// MockBuildClient is an autogenerated mock type for the BuildClient type
type MockBuildClient struct {
	mock.Mock
}

type MockBuildClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBuildClient) EXPECT() *MockBuildClient_Expecter {
	return &MockBuildClient_Expecter{mock: &_m.Mock}
}

// GetBuildLogLines provides a mock function with given fields: ctx, args
func (_m *MockBuildClient) GetBuildLogLines(ctx context.Context, args build.GetBuildLogLinesArgs) (*[]string, error) {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetBuildLogLines")
	}

	var r0 *[]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildLogLinesArgs) (*[]string, error)); ok {
		return rf(ctx, args)
	}
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildLogLinesArgs) *[]string); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, build.GetBuildLogLinesArgs) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBuildClient_GetBuildLogLines_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBuildLogLines'
type MockBuildClient_GetBuildLogLines_Call struct {
	*mock.Call
}

// GetBuildLogLines is a helper method to define mock.On call
//   - ctx context.Context
//   - args build.GetBuildLogLinesArgs
func (_e *MockBuildClient_Expecter) GetBuildLogLines(ctx interface{}, args interface{}) *MockBuildClient_GetBuildLogLines_Call {
	return &MockBuildClient_GetBuildLogLines_Call{Call: _e.mock.On("GetBuildLogLines", ctx, args)}
}

func (_c *MockBuildClient_GetBuildLogLines_Call) Run(run func(ctx context.Context, args build.GetBuildLogLinesArgs)) *MockBuildClient_GetBuildLogLines_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(build.GetBuildLogLinesArgs))
	})
	return _c
}

func (_c *MockBuildClient_GetBuildLogLines_Call) Return(_a0 *[]string, _a1 error) *MockBuildClient_GetBuildLogLines_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBuildClient_GetBuildLogLines_Call) RunAndReturn(run func(context.Context, build.GetBuildLogLinesArgs) (*[]string, error)) *MockBuildClient_GetBuildLogLines_Call {
	_c.Call.Return(run)
	return _c
}

// GetBuildLogs provides a mock function with given fields: ctx, args
func (_m *MockBuildClient) GetBuildLogs(ctx context.Context, args build.GetBuildLogsArgs) (*[]build.BuildLog, error) {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetBuildLogs")
	}

	var r0 *[]build.BuildLog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildLogsArgs) (*[]build.BuildLog, error)); ok {
		return rf(ctx, args)
	}
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildLogsArgs) *[]build.BuildLog); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]build.BuildLog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, build.GetBuildLogsArgs) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBuildClient_GetBuildLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBuildLogs'
type MockBuildClient_GetBuildLogs_Call struct {
	*mock.Call
}

// GetBuildLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - args build.GetBuildLogsArgs
func (_e *MockBuildClient_Expecter) GetBuildLogs(ctx interface{}, args interface{}) *MockBuildClient_GetBuildLogs_Call {
	return &MockBuildClient_GetBuildLogs_Call{Call: _e.mock.On("GetBuildLogs", ctx, args)}
}

func (_c *MockBuildClient_GetBuildLogs_Call) Run(run func(ctx context.Context, args build.GetBuildLogsArgs)) *MockBuildClient_GetBuildLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(build.GetBuildLogsArgs))
	})
	return _c
}

func (_c *MockBuildClient_GetBuildLogs_Call) Return(_a0 *[]build.BuildLog, _a1 error) *MockBuildClient_GetBuildLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBuildClient_GetBuildLogs_Call) RunAndReturn(run func(context.Context, build.GetBuildLogsArgs) (*[]build.BuildLog, error)) *MockBuildClient_GetBuildLogs_Call {
	_c.Call.Return(run)
	return _c
}

// GetBuildTimeline provides a mock function with given fields: ctx, args
func (_m *MockBuildClient) GetBuildTimeline(ctx context.Context, args build.GetBuildTimelineArgs) (*build.Timeline, error) {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetBuildTimeline")
	}

	var r0 *build.Timeline
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildTimelineArgs) (*build.Timeline, error)); ok {
		return rf(ctx, args)
	}
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildTimelineArgs) *build.Timeline); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*build.Timeline)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, build.GetBuildTimelineArgs) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBuildClient_GetBuildTimeline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBuildTimeline'
type MockBuildClient_GetBuildTimeline_Call struct {
	*mock.Call
}

// GetBuildTimeline is a helper method to define mock.On call
//   - ctx context.Context
//   - args build.GetBuildTimelineArgs
func (_e *MockBuildClient_Expecter) GetBuildTimeline(ctx interface{}, args interface{}) *MockBuildClient_GetBuildTimeline_Call {
	return &MockBuildClient_GetBuildTimeline_Call{Call: _e.mock.On("GetBuildTimeline", ctx, args)}
}

func (_c *MockBuildClient_GetBuildTimeline_Call) Run(run func(ctx context.Context, args build.GetBuildTimelineArgs)) *MockBuildClient_GetBuildTimeline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(build.GetBuildTimelineArgs))
	})
	return _c
}

func (_c *MockBuildClient_GetBuildTimeline_Call) Return(_a0 *build.Timeline, _a1 error) *MockBuildClient_GetBuildTimeline_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBuildClient_GetBuildTimeline_Call) RunAndReturn(run func(context.Context, build.GetBuildTimelineArgs) (*build.Timeline, error)) *MockBuildClient_GetBuildTimeline_Call {
	_c.Call.Return(run)
	return _c
}

// GetBuilds provides a mock function with given fields: ctx, args
func (_m *MockBuildClient) GetBuilds(ctx context.Context, args build.GetBuildsArgs) (*build.GetBuildsResponseValue, error) {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetBuilds")
	}

	var r0 *build.GetBuildsResponseValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildsArgs) (*build.GetBuildsResponseValue, error)); ok {
		return rf(ctx, args)
	}
	if rf, ok := ret.Get(0).(func(context.Context, build.GetBuildsArgs) *build.GetBuildsResponseValue); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*build.GetBuildsResponseValue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, build.GetBuildsArgs) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBuildClient_GetBuilds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBuilds'
type MockBuildClient_GetBuilds_Call struct {
	*mock.Call
}

// GetBuilds is a helper method to define mock.On call
//   - ctx context.Context
//   - args build.GetBuildsArgs
func (_e *MockBuildClient_Expecter) GetBuilds(ctx interface{}, args interface{}) *MockBuildClient_GetBuilds_Call {
	return &MockBuildClient_GetBuilds_Call{Call: _e.mock.On("GetBuilds", ctx, args)}
}

func (_c *MockBuildClient_GetBuilds_Call) Run(run func(ctx context.Context, args build.GetBuildsArgs)) *MockBuildClient_GetBuilds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(build.GetBuildsArgs))
	})
	return _c
}

func (_c *MockBuildClient_GetBuilds_Call) Return(_a0 *build.GetBuildsResponseValue, _a1 error) *MockBuildClient_GetBuilds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBuildClient_GetBuilds_Call) RunAndReturn(run func(context.Context, build.GetBuildsArgs) (*build.GetBuildsResponseValue, error)) *MockBuildClient_GetBuilds_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBuildClient creates a new instance of MockBuildClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBuildClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBuildClient {
	mock := &MockBuildClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
