// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/connect/v2/x/marketmap/types"
)

// PriceAggregator is an autogenerated mock type for the PriceAggregator type
type PriceAggregator struct {
	mock.Mock
}

type PriceAggregator_Expecter struct {
	mock *mock.Mock
}

func (_m *PriceAggregator) EXPECT() *PriceAggregator_Expecter {
	return &PriceAggregator_Expecter{mock: &_m.Mock}
}

// AggregatePrices provides a mock function with no fields
func (_m *PriceAggregator) AggregatePrices() {
	_m.Called()
}

// PriceAggregator_AggregatePrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AggregatePrices'
type PriceAggregator_AggregatePrices_Call struct {
	*mock.Call
}

// AggregatePrices is a helper method to define mock.On call
func (_e *PriceAggregator_Expecter) AggregatePrices() *PriceAggregator_AggregatePrices_Call {
	return &PriceAggregator_AggregatePrices_Call{Call: _e.mock.On("AggregatePrices")}
}

func (_c *PriceAggregator_AggregatePrices_Call) Run(run func()) *PriceAggregator_AggregatePrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PriceAggregator_AggregatePrices_Call) Return() *PriceAggregator_AggregatePrices_Call {
	_c.Call.Return()
	return _c
}

func (_c *PriceAggregator_AggregatePrices_Call) RunAndReturn(run func()) *PriceAggregator_AggregatePrices_Call {
	_c.Run(run)
	return _c
}

// GetPrices provides a mock function with no fields
func (_m *PriceAggregator) GetPrices() map[string]*big.Float {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPrices")
	}

	var r0 map[string]*big.Float
	if rf, ok := ret.Get(0).(func() map[string]*big.Float); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*big.Float)
		}
	}

	return r0
}

// PriceAggregator_GetPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrices'
type PriceAggregator_GetPrices_Call struct {
	*mock.Call
}

// GetPrices is a helper method to define mock.On call
func (_e *PriceAggregator_Expecter) GetPrices() *PriceAggregator_GetPrices_Call {
	return &PriceAggregator_GetPrices_Call{Call: _e.mock.On("GetPrices")}
}

func (_c *PriceAggregator_GetPrices_Call) Run(run func()) *PriceAggregator_GetPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PriceAggregator_GetPrices_Call) Return(_a0 map[string]*big.Float) *PriceAggregator_GetPrices_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PriceAggregator_GetPrices_Call) RunAndReturn(run func() map[string]*big.Float) *PriceAggregator_GetPrices_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with no fields
func (_m *PriceAggregator) Reset() {
	_m.Called()
}

// PriceAggregator_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type PriceAggregator_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *PriceAggregator_Expecter) Reset() *PriceAggregator_Reset_Call {
	return &PriceAggregator_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *PriceAggregator_Reset_Call) Run(run func()) *PriceAggregator_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PriceAggregator_Reset_Call) Return() *PriceAggregator_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *PriceAggregator_Reset_Call) RunAndReturn(run func()) *PriceAggregator_Reset_Call {
	_c.Run(run)
	return _c
}

// SetProviderPrices provides a mock function with given fields: provider, prices
func (_m *PriceAggregator) SetProviderPrices(provider string, prices map[string]*big.Float) {
	_m.Called(provider, prices)
}

// PriceAggregator_SetProviderPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetProviderPrices'
type PriceAggregator_SetProviderPrices_Call struct {
	*mock.Call
}

// SetProviderPrices is a helper method to define mock.On call
//   - provider string
//   - prices map[string]*big.Float
func (_e *PriceAggregator_Expecter) SetProviderPrices(provider interface{}, prices interface{}) *PriceAggregator_SetProviderPrices_Call {
	return &PriceAggregator_SetProviderPrices_Call{Call: _e.mock.On("SetProviderPrices", provider, prices)}
}

func (_c *PriceAggregator_SetProviderPrices_Call) Run(run func(provider string, prices map[string]*big.Float)) *PriceAggregator_SetProviderPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(map[string]*big.Float))
	})
	return _c
}

func (_c *PriceAggregator_SetProviderPrices_Call) Return() *PriceAggregator_SetProviderPrices_Call {
	_c.Call.Return()
	return _c
}

func (_c *PriceAggregator_SetProviderPrices_Call) RunAndReturn(run func(string, map[string]*big.Float)) *PriceAggregator_SetProviderPrices_Call {
	_c.Run(run)
	return _c
}

// UpdateMarketMap provides a mock function with given fields: _a0
func (_m *PriceAggregator) UpdateMarketMap(_a0 types.MarketMap) {
	_m.Called(_a0)
}

// PriceAggregator_UpdateMarketMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMarketMap'
type PriceAggregator_UpdateMarketMap_Call struct {
	*mock.Call
}

// UpdateMarketMap is a helper method to define mock.On call
//   - _a0 types.MarketMap
func (_e *PriceAggregator_Expecter) UpdateMarketMap(_a0 interface{}) *PriceAggregator_UpdateMarketMap_Call {
	return &PriceAggregator_UpdateMarketMap_Call{Call: _e.mock.On("UpdateMarketMap", _a0)}
}

func (_c *PriceAggregator_UpdateMarketMap_Call) Run(run func(_a0 types.MarketMap)) *PriceAggregator_UpdateMarketMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.MarketMap))
	})
	return _c
}

func (_c *PriceAggregator_UpdateMarketMap_Call) Return() *PriceAggregator_UpdateMarketMap_Call {
	_c.Call.Return()
	return _c
}

func (_c *PriceAggregator_UpdateMarketMap_Call) RunAndReturn(run func(types.MarketMap)) *PriceAggregator_UpdateMarketMap_Call {
	_c.Run(run)
	return _c
}

// NewPriceAggregator creates a new instance of PriceAggregator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceAggregator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceAggregator {
	mock := &PriceAggregator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
