// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	core "github.com/mechiko/maroto/v2/pkg/core"
	entity "github.com/mechiko/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	node "github.com/johnfercher/go-tree/node"

	props "github.com/mechiko/maroto/v2/pkg/props"
)

// Col is an autogenerated mock type for the Col type
type Col struct {
	mock.Mock
}

type Col_Expecter struct {
	mock *mock.Mock
}

func (_m *Col) EXPECT() *Col_Expecter {
	return &Col_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: components
func (_m *Col) Add(components ...core.Component) core.Col {
	_va := make([]interface{}, len(components))
	for _i := range components {
		_va[_i] = components[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 core.Col
	if rf, ok := ret.Get(0).(func(...core.Component) core.Col); ok {
		r0 = rf(components...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Col)
		}
	}

	return r0
}

// Col_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Col_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - components ...core.Component
func (_e *Col_Expecter) Add(components ...interface{}) *Col_Add_Call {
	return &Col_Add_Call{Call: _e.mock.On("Add",
		append([]interface{}{}, components...)...)}
}

func (_c *Col_Add_Call) Run(run func(components ...core.Component)) *Col_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Component, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(core.Component)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Col_Add_Call) Return(_a0 core.Col) *Col_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Col_Add_Call) RunAndReturn(run func(...core.Component) core.Col) *Col_Add_Call {
	_c.Call.Return(run)
	return _c
}

// GetHeight provides a mock function with given fields: provider, cell
func (_m *Col) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	ret := _m.Called(provider, cell)

	if len(ret) == 0 {
		panic("no return value specified for GetHeight")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func(core.Provider, *entity.Cell) float64); ok {
		r0 = rf(provider, cell)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Col_GetHeight_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeight'
type Col_GetHeight_Call struct {
	*mock.Call
}

// GetHeight is a helper method to define mock.On call
//   - provider core.Provider
//   - cell *entity.Cell
func (_e *Col_Expecter) GetHeight(provider interface{}, cell interface{}) *Col_GetHeight_Call {
	return &Col_GetHeight_Call{Call: _e.mock.On("GetHeight", provider, cell)}
}

func (_c *Col_GetHeight_Call) Run(run func(provider core.Provider, cell *entity.Cell)) *Col_GetHeight_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(core.Provider), args[1].(*entity.Cell))
	})
	return _c
}

func (_c *Col_GetHeight_Call) Return(_a0 float64) *Col_GetHeight_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Col_GetHeight_Call) RunAndReturn(run func(core.Provider, *entity.Cell) float64) *Col_GetHeight_Call {
	_c.Call.Return(run)
	return _c
}

// GetSize provides a mock function with given fields:
func (_m *Col) GetSize() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSize")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Col_GetSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSize'
type Col_GetSize_Call struct {
	*mock.Call
}

// GetSize is a helper method to define mock.On call
func (_e *Col_Expecter) GetSize() *Col_GetSize_Call {
	return &Col_GetSize_Call{Call: _e.mock.On("GetSize")}
}

func (_c *Col_GetSize_Call) Run(run func()) *Col_GetSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Col_GetSize_Call) Return(_a0 int) *Col_GetSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Col_GetSize_Call) RunAndReturn(run func() int) *Col_GetSize_Call {
	_c.Call.Return(run)
	return _c
}

// GetStructure provides a mock function with given fields:
func (_m *Col) GetStructure() *node.Node[core.Structure] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStructure")
	}

	var r0 *node.Node[core.Structure]
	if rf, ok := ret.Get(0).(func() *node.Node[core.Structure]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Node[core.Structure])
		}
	}

	return r0
}

// Col_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type Col_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *Col_Expecter) GetStructure() *Col_GetStructure_Call {
	return &Col_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *Col_GetStructure_Call) Run(run func()) *Col_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Col_GetStructure_Call) Return(_a0 *node.Node[core.Structure]) *Col_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Col_GetStructure_Call) RunAndReturn(run func() *node.Node[core.Structure]) *Col_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// Render provides a mock function with given fields: provider, cell, createCell
func (_m *Col) Render(provider core.Provider, cell entity.Cell, createCell bool) {
	_m.Called(provider, cell, createCell)
}

// Col_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type Col_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - provider core.Provider
//   - cell entity.Cell
//   - createCell bool
func (_e *Col_Expecter) Render(provider interface{}, cell interface{}, createCell interface{}) *Col_Render_Call {
	return &Col_Render_Call{Call: _e.mock.On("Render", provider, cell, createCell)}
}

func (_c *Col_Render_Call) Run(run func(provider core.Provider, cell entity.Cell, createCell bool)) *Col_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(core.Provider), args[1].(entity.Cell), args[2].(bool))
	})
	return _c
}

func (_c *Col_Render_Call) Return() *Col_Render_Call {
	_c.Call.Return()
	return _c
}

func (_c *Col_Render_Call) RunAndReturn(run func(core.Provider, entity.Cell, bool)) *Col_Render_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfig provides a mock function with given fields: config
func (_m *Col) SetConfig(config *entity.Config) {
	_m.Called(config)
}

// Col_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type Col_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//   - config *entity.Config
func (_e *Col_Expecter) SetConfig(config interface{}) *Col_SetConfig_Call {
	return &Col_SetConfig_Call{Call: _e.mock.On("SetConfig", config)}
}

func (_c *Col_SetConfig_Call) Run(run func(config *entity.Config)) *Col_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Config))
	})
	return _c
}

func (_c *Col_SetConfig_Call) Return() *Col_SetConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *Col_SetConfig_Call) RunAndReturn(run func(*entity.Config)) *Col_SetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// WithStyle provides a mock function with given fields: style
func (_m *Col) WithStyle(style *props.Cell) core.Col {
	ret := _m.Called(style)

	if len(ret) == 0 {
		panic("no return value specified for WithStyle")
	}

	var r0 core.Col
	if rf, ok := ret.Get(0).(func(*props.Cell) core.Col); ok {
		r0 = rf(style)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Col)
		}
	}

	return r0
}

// Col_WithStyle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithStyle'
type Col_WithStyle_Call struct {
	*mock.Call
}

// WithStyle is a helper method to define mock.On call
//   - style *props.Cell
func (_e *Col_Expecter) WithStyle(style interface{}) *Col_WithStyle_Call {
	return &Col_WithStyle_Call{Call: _e.mock.On("WithStyle", style)}
}

func (_c *Col_WithStyle_Call) Run(run func(style *props.Cell)) *Col_WithStyle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*props.Cell))
	})
	return _c
}

func (_c *Col_WithStyle_Call) Return(_a0 core.Col) *Col_WithStyle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Col_WithStyle_Call) RunAndReturn(run func(*props.Cell) core.Col) *Col_WithStyle_Call {
	_c.Call.Return(run)
	return _c
}

// NewCol creates a new instance of Col. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCol(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Col {
	mock := &Col{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
