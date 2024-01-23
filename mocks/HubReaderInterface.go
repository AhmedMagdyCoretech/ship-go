// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"
)

// HubReaderInterface is an autogenerated mock type for the HubReaderInterface type
type HubReaderInterface struct {
	mock.Mock
}

type HubReaderInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *HubReaderInterface) EXPECT() *HubReaderInterface_Expecter {
	return &HubReaderInterface_Expecter{mock: &_m.Mock}
}

// AllowWaitingForTrust provides a mock function with given fields: ski
func (_m *HubReaderInterface) AllowWaitingForTrust(ski string) bool {
	ret := _m.Called(ski)

	if len(ret) == 0 {
		panic("no return value specified for AllowWaitingForTrust")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ski)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HubReaderInterface_AllowWaitingForTrust_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllowWaitingForTrust'
type HubReaderInterface_AllowWaitingForTrust_Call struct {
	*mock.Call
}

// AllowWaitingForTrust is a helper method to define mock.On call
//   - ski string
func (_e *HubReaderInterface_Expecter) AllowWaitingForTrust(ski interface{}) *HubReaderInterface_AllowWaitingForTrust_Call {
	return &HubReaderInterface_AllowWaitingForTrust_Call{Call: _e.mock.On("AllowWaitingForTrust", ski)}
}

func (_c *HubReaderInterface_AllowWaitingForTrust_Call) Run(run func(ski string)) *HubReaderInterface_AllowWaitingForTrust_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubReaderInterface_AllowWaitingForTrust_Call) Return(_a0 bool) *HubReaderInterface_AllowWaitingForTrust_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HubReaderInterface_AllowWaitingForTrust_Call) RunAndReturn(run func(string) bool) *HubReaderInterface_AllowWaitingForTrust_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteSKIConnected provides a mock function with given fields: ski
func (_m *HubReaderInterface) RemoteSKIConnected(ski string) {
	_m.Called(ski)
}

// HubReaderInterface_RemoteSKIConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteSKIConnected'
type HubReaderInterface_RemoteSKIConnected_Call struct {
	*mock.Call
}

// RemoteSKIConnected is a helper method to define mock.On call
//   - ski string
func (_e *HubReaderInterface_Expecter) RemoteSKIConnected(ski interface{}) *HubReaderInterface_RemoteSKIConnected_Call {
	return &HubReaderInterface_RemoteSKIConnected_Call{Call: _e.mock.On("RemoteSKIConnected", ski)}
}

func (_c *HubReaderInterface_RemoteSKIConnected_Call) Run(run func(ski string)) *HubReaderInterface_RemoteSKIConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubReaderInterface_RemoteSKIConnected_Call) Return() *HubReaderInterface_RemoteSKIConnected_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubReaderInterface_RemoteSKIConnected_Call) RunAndReturn(run func(string)) *HubReaderInterface_RemoteSKIConnected_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteSKIDisconnected provides a mock function with given fields: ski
func (_m *HubReaderInterface) RemoteSKIDisconnected(ski string) {
	_m.Called(ski)
}

// HubReaderInterface_RemoteSKIDisconnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteSKIDisconnected'
type HubReaderInterface_RemoteSKIDisconnected_Call struct {
	*mock.Call
}

// RemoteSKIDisconnected is a helper method to define mock.On call
//   - ski string
func (_e *HubReaderInterface_Expecter) RemoteSKIDisconnected(ski interface{}) *HubReaderInterface_RemoteSKIDisconnected_Call {
	return &HubReaderInterface_RemoteSKIDisconnected_Call{Call: _e.mock.On("RemoteSKIDisconnected", ski)}
}

func (_c *HubReaderInterface_RemoteSKIDisconnected_Call) Run(run func(ski string)) *HubReaderInterface_RemoteSKIDisconnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubReaderInterface_RemoteSKIDisconnected_Call) Return() *HubReaderInterface_RemoteSKIDisconnected_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubReaderInterface_RemoteSKIDisconnected_Call) RunAndReturn(run func(string)) *HubReaderInterface_RemoteSKIDisconnected_Call {
	_c.Call.Return(run)
	return _c
}

// ServicePairingDetailUpdate provides a mock function with given fields: ski, detail
func (_m *HubReaderInterface) ServicePairingDetailUpdate(ski string, detail *api.ConnectionStateDetail) {
	_m.Called(ski, detail)
}

// HubReaderInterface_ServicePairingDetailUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServicePairingDetailUpdate'
type HubReaderInterface_ServicePairingDetailUpdate_Call struct {
	*mock.Call
}

// ServicePairingDetailUpdate is a helper method to define mock.On call
//   - ski string
//   - detail *api.ConnectionStateDetail
func (_e *HubReaderInterface_Expecter) ServicePairingDetailUpdate(ski interface{}, detail interface{}) *HubReaderInterface_ServicePairingDetailUpdate_Call {
	return &HubReaderInterface_ServicePairingDetailUpdate_Call{Call: _e.mock.On("ServicePairingDetailUpdate", ski, detail)}
}

func (_c *HubReaderInterface_ServicePairingDetailUpdate_Call) Run(run func(ski string, detail *api.ConnectionStateDetail)) *HubReaderInterface_ServicePairingDetailUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*api.ConnectionStateDetail))
	})
	return _c
}

func (_c *HubReaderInterface_ServicePairingDetailUpdate_Call) Return() *HubReaderInterface_ServicePairingDetailUpdate_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubReaderInterface_ServicePairingDetailUpdate_Call) RunAndReturn(run func(string, *api.ConnectionStateDetail)) *HubReaderInterface_ServicePairingDetailUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// ServiceShipIDUpdate provides a mock function with given fields: ski, shipID
func (_m *HubReaderInterface) ServiceShipIDUpdate(ski string, shipID string) {
	_m.Called(ski, shipID)
}

// HubReaderInterface_ServiceShipIDUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServiceShipIDUpdate'
type HubReaderInterface_ServiceShipIDUpdate_Call struct {
	*mock.Call
}

// ServiceShipIDUpdate is a helper method to define mock.On call
//   - ski string
//   - shipID string
func (_e *HubReaderInterface_Expecter) ServiceShipIDUpdate(ski interface{}, shipID interface{}) *HubReaderInterface_ServiceShipIDUpdate_Call {
	return &HubReaderInterface_ServiceShipIDUpdate_Call{Call: _e.mock.On("ServiceShipIDUpdate", ski, shipID)}
}

func (_c *HubReaderInterface_ServiceShipIDUpdate_Call) Run(run func(ski string, shipID string)) *HubReaderInterface_ServiceShipIDUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *HubReaderInterface_ServiceShipIDUpdate_Call) Return() *HubReaderInterface_ServiceShipIDUpdate_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubReaderInterface_ServiceShipIDUpdate_Call) RunAndReturn(run func(string, string)) *HubReaderInterface_ServiceShipIDUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// SetupRemoteDevice provides a mock function with given fields: ski, writeI
func (_m *HubReaderInterface) SetupRemoteDevice(ski string, writeI api.ShipConnectionDataWriterInterface) api.ShipConnectionDataReaderInterface {
	ret := _m.Called(ski, writeI)

	if len(ret) == 0 {
		panic("no return value specified for SetupRemoteDevice")
	}

	var r0 api.ShipConnectionDataReaderInterface
	if rf, ok := ret.Get(0).(func(string, api.ShipConnectionDataWriterInterface) api.ShipConnectionDataReaderInterface); ok {
		r0 = rf(ski, writeI)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.ShipConnectionDataReaderInterface)
		}
	}

	return r0
}

// HubReaderInterface_SetupRemoteDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetupRemoteDevice'
type HubReaderInterface_SetupRemoteDevice_Call struct {
	*mock.Call
}

// SetupRemoteDevice is a helper method to define mock.On call
//   - ski string
//   - writeI api.ShipConnectionDataWriterInterface
func (_e *HubReaderInterface_Expecter) SetupRemoteDevice(ski interface{}, writeI interface{}) *HubReaderInterface_SetupRemoteDevice_Call {
	return &HubReaderInterface_SetupRemoteDevice_Call{Call: _e.mock.On("SetupRemoteDevice", ski, writeI)}
}

func (_c *HubReaderInterface_SetupRemoteDevice_Call) Run(run func(ski string, writeI api.ShipConnectionDataWriterInterface)) *HubReaderInterface_SetupRemoteDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(api.ShipConnectionDataWriterInterface))
	})
	return _c
}

func (_c *HubReaderInterface_SetupRemoteDevice_Call) Return(_a0 api.ShipConnectionDataReaderInterface) *HubReaderInterface_SetupRemoteDevice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HubReaderInterface_SetupRemoteDevice_Call) RunAndReturn(run func(string, api.ShipConnectionDataWriterInterface) api.ShipConnectionDataReaderInterface) *HubReaderInterface_SetupRemoteDevice_Call {
	_c.Call.Return(run)
	return _c
}

// VisibleMDNSRecordsUpdated provides a mock function with given fields: entries
func (_m *HubReaderInterface) VisibleMDNSRecordsUpdated(entries []*api.MdnsEntry) {
	_m.Called(entries)
}

// HubReaderInterface_VisibleMDNSRecordsUpdated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VisibleMDNSRecordsUpdated'
type HubReaderInterface_VisibleMDNSRecordsUpdated_Call struct {
	*mock.Call
}

// VisibleMDNSRecordsUpdated is a helper method to define mock.On call
//   - entries []*api.MdnsEntry
func (_e *HubReaderInterface_Expecter) VisibleMDNSRecordsUpdated(entries interface{}) *HubReaderInterface_VisibleMDNSRecordsUpdated_Call {
	return &HubReaderInterface_VisibleMDNSRecordsUpdated_Call{Call: _e.mock.On("VisibleMDNSRecordsUpdated", entries)}
}

func (_c *HubReaderInterface_VisibleMDNSRecordsUpdated_Call) Run(run func(entries []*api.MdnsEntry)) *HubReaderInterface_VisibleMDNSRecordsUpdated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*api.MdnsEntry))
	})
	return _c
}

func (_c *HubReaderInterface_VisibleMDNSRecordsUpdated_Call) Return() *HubReaderInterface_VisibleMDNSRecordsUpdated_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubReaderInterface_VisibleMDNSRecordsUpdated_Call) RunAndReturn(run func([]*api.MdnsEntry)) *HubReaderInterface_VisibleMDNSRecordsUpdated_Call {
	_c.Call.Return(run)
	return _c
}

// NewHubReaderInterface creates a new instance of HubReaderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHubReaderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *HubReaderInterface {
	mock := &HubReaderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
