// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LoggingInterface is an autogenerated mock type for the LoggingInterface type
type LoggingInterface struct {
	mock.Mock
}

type LoggingInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *LoggingInterface) EXPECT() *LoggingInterface_Expecter {
	return &LoggingInterface_Expecter{mock: &_m.Mock}
}

// Debug provides a mock function with given fields: args
func (_m *LoggingInterface) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type LoggingInterface_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Debug(args ...interface{}) *LoggingInterface_Debug_Call {
	return &LoggingInterface_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{}, args...)...)}
}

func (_c *LoggingInterface_Debug_Call) Run(run func(args ...interface{})) *LoggingInterface_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Debug_Call) Return() *LoggingInterface_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Debug_Call) RunAndReturn(run func(...interface{})) *LoggingInterface_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Debugf provides a mock function with given fields: format, args
func (_m *LoggingInterface) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Debugf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugf'
type LoggingInterface_Debugf_Call struct {
	*mock.Call
}

// Debugf is a helper method to define mock.On call
//   - format string
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Debugf(format interface{}, args ...interface{}) *LoggingInterface_Debugf_Call {
	return &LoggingInterface_Debugf_Call{Call: _e.mock.On("Debugf",
		append([]interface{}{format}, args...)...)}
}

func (_c *LoggingInterface_Debugf_Call) Run(run func(format string, args ...interface{})) *LoggingInterface_Debugf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Debugf_Call) Return() *LoggingInterface_Debugf_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Debugf_Call) RunAndReturn(run func(string, ...interface{})) *LoggingInterface_Debugf_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: args
func (_m *LoggingInterface) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type LoggingInterface_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Error(args ...interface{}) *LoggingInterface_Error_Call {
	return &LoggingInterface_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{}, args...)...)}
}

func (_c *LoggingInterface_Error_Call) Run(run func(args ...interface{})) *LoggingInterface_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Error_Call) Return() *LoggingInterface_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Error_Call) RunAndReturn(run func(...interface{})) *LoggingInterface_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Errorf provides a mock function with given fields: format, args
func (_m *LoggingInterface) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Errorf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorf'
type LoggingInterface_Errorf_Call struct {
	*mock.Call
}

// Errorf is a helper method to define mock.On call
//   - format string
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Errorf(format interface{}, args ...interface{}) *LoggingInterface_Errorf_Call {
	return &LoggingInterface_Errorf_Call{Call: _e.mock.On("Errorf",
		append([]interface{}{format}, args...)...)}
}

func (_c *LoggingInterface_Errorf_Call) Run(run func(format string, args ...interface{})) *LoggingInterface_Errorf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Errorf_Call) Return() *LoggingInterface_Errorf_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Errorf_Call) RunAndReturn(run func(string, ...interface{})) *LoggingInterface_Errorf_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: args
func (_m *LoggingInterface) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type LoggingInterface_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Info(args ...interface{}) *LoggingInterface_Info_Call {
	return &LoggingInterface_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{}, args...)...)}
}

func (_c *LoggingInterface_Info_Call) Run(run func(args ...interface{})) *LoggingInterface_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Info_Call) Return() *LoggingInterface_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Info_Call) RunAndReturn(run func(...interface{})) *LoggingInterface_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Infof provides a mock function with given fields: format, args
func (_m *LoggingInterface) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Infof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infof'
type LoggingInterface_Infof_Call struct {
	*mock.Call
}

// Infof is a helper method to define mock.On call
//   - format string
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Infof(format interface{}, args ...interface{}) *LoggingInterface_Infof_Call {
	return &LoggingInterface_Infof_Call{Call: _e.mock.On("Infof",
		append([]interface{}{format}, args...)...)}
}

func (_c *LoggingInterface_Infof_Call) Run(run func(format string, args ...interface{})) *LoggingInterface_Infof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Infof_Call) Return() *LoggingInterface_Infof_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Infof_Call) RunAndReturn(run func(string, ...interface{})) *LoggingInterface_Infof_Call {
	_c.Call.Return(run)
	return _c
}

// Trace provides a mock function with given fields: args
func (_m *LoggingInterface) Trace(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Trace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trace'
type LoggingInterface_Trace_Call struct {
	*mock.Call
}

// Trace is a helper method to define mock.On call
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Trace(args ...interface{}) *LoggingInterface_Trace_Call {
	return &LoggingInterface_Trace_Call{Call: _e.mock.On("Trace",
		append([]interface{}{}, args...)...)}
}

func (_c *LoggingInterface_Trace_Call) Run(run func(args ...interface{})) *LoggingInterface_Trace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Trace_Call) Return() *LoggingInterface_Trace_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Trace_Call) RunAndReturn(run func(...interface{})) *LoggingInterface_Trace_Call {
	_c.Call.Return(run)
	return _c
}

// Tracef provides a mock function with given fields: format, args
func (_m *LoggingInterface) Tracef(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LoggingInterface_Tracef_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tracef'
type LoggingInterface_Tracef_Call struct {
	*mock.Call
}

// Tracef is a helper method to define mock.On call
//   - format string
//   - args ...interface{}
func (_e *LoggingInterface_Expecter) Tracef(format interface{}, args ...interface{}) *LoggingInterface_Tracef_Call {
	return &LoggingInterface_Tracef_Call{Call: _e.mock.On("Tracef",
		append([]interface{}{format}, args...)...)}
}

func (_c *LoggingInterface_Tracef_Call) Run(run func(format string, args ...interface{})) *LoggingInterface_Tracef_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *LoggingInterface_Tracef_Call) Return() *LoggingInterface_Tracef_Call {
	_c.Call.Return()
	return _c
}

func (_c *LoggingInterface_Tracef_Call) RunAndReturn(run func(string, ...interface{})) *LoggingInterface_Tracef_Call {
	_c.Call.Return(run)
	return _c
}

// NewLoggingInterface creates a new instance of LoggingInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoggingInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoggingInterface {
	mock := &LoggingInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
