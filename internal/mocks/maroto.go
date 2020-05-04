// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bytes "bytes"
import color "github.com/johnfercher/maroto/pkg/color"
import consts "github.com/johnfercher/maroto/pkg/consts"
import mock "github.com/stretchr/testify/mock"

import props "github.com/johnfercher/maroto/pkg/props"

// Maroto is an autogenerated mock type for the Maroto type
type Maroto struct {
	mock.Mock
}

// Barcode provides a mock function with given fields: code, prop
func (_m *Maroto) Barcode(code string, prop ...props.Barcode) error {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, code)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...props.Barcode) error); ok {
		r0 = rf(code, prop...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Base64Image provides a mock function with given fields: base64, extension, prop
func (_m *Maroto) Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) error {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, base64, extension)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, consts.Extension, ...props.Rect) error); ok {
		r0 = rf(base64, extension, prop...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Col provides a mock function with given fields: closure
func (_m *Maroto) Col(width uint, closure func()) {
	_m.Called(closure)
}

// ColSpace provides a mock function with given fields:
func (_m *Maroto) ColSpace(uint) {
	_m.Called()
}

// ColSpaces provides a mock function with given fields: qtd
func (_m *Maroto) ColSpaces(qtd int) {
	_m.Called(qtd)
}

// FileImage provides a mock function with given fields: filePathName, prop
func (_m *Maroto) FileImage(filePathName string, prop ...props.Rect) error {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, filePathName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...props.Rect) error); ok {
		r0 = rf(filePathName, prop...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBorder provides a mock function with given fields:
func (_m *Maroto) GetBorder() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetCurrentOffset provides a mock function with given fields:
func (_m *Maroto) GetCurrentOffset() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetCurrentPage provides a mock function with given fields:
func (_m *Maroto) GetCurrentPage() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetPageMargins provides a mock function with given fields:
func (_m *Maroto) GetPageMargins() (float64, float64, float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 float64
	if rf, ok := ret.Get(2).(func() float64); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(float64)
	}

	var r3 float64
	if rf, ok := ret.Get(3).(func() float64); ok {
		r3 = rf()
	} else {
		r3 = ret.Get(3).(float64)
	}

	return r0, r1, r2, r3
}

// GetPageSize provides a mock function with given fields:
func (_m *Maroto) GetPageSize() (float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// Line provides a mock function with given fields: spaceHeight
func (_m *Maroto) Line(spaceHeight float64) {
	_m.Called(spaceHeight)
}

// Output provides a mock function with given fields:
func (_m *Maroto) Output() (bytes.Buffer, error) {
	ret := _m.Called()

	var r0 bytes.Buffer
	if rf, ok := ret.Get(0).(func() bytes.Buffer); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bytes.Buffer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OutputFileAndClose provides a mock function with given fields: filePathName
func (_m *Maroto) OutputFileAndClose(filePathName string) error {
	ret := _m.Called(filePathName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filePathName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QrCode provides a mock function with given fields: code, prop
func (_m *Maroto) QrCode(code string, prop ...props.Rect) {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, code)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// RegisterFooter provides a mock function with given fields: closure
func (_m *Maroto) RegisterFooter(closure func()) {
	_m.Called(closure)
}

// RegisterHeader provides a mock function with given fields: closure
func (_m *Maroto) RegisterHeader(closure func()) {
	_m.Called(closure)
}

// Row provides a mock function with given fields: height, closure
func (_m *Maroto) Row(height float64, closure func()) {
	_m.Called(height, closure)
}

// SetBackgroundColor provides a mock function with given fields: _a0
func (_m *Maroto) SetBackgroundColor(_a0 color.Color) {
	_m.Called(_a0)
}

// SetTextColor provides a mock function with given fields: _a0
func (_m *Maroto) SetTextColor(_a0 color.Color) {
	_m.Called(_a0)
}

// SetBorder provides a mock function with given fields: on
func (_m *Maroto) SetBorder(on bool) {
	_m.Called(on)
}

// SetPageMargins provides a mock function with given fields: left, top, right
func (_m *Maroto) SetPageMargins(left float64, top float64, right float64) {
	_m.Called(left, top, right)
}

// Signature provides a mock function with given fields: label, prop
func (_m *Maroto) Signature(label string, prop ...props.Font) {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, label)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// TableList provides a mock function with given fields: header, contents, prop
func (_m *Maroto) TableList(header []string, contents [][]string, prop ...props.TableList) {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, header, contents)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Text provides a mock function with given fields: text, prop
func (_m *Maroto) Text(text string, prop ...props.Text) {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, text)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Text provides a mock function with given fields: text, prop
func (_m *Maroto) TextWithLink(text string, link int, c color.Color, prop ...props.Text) {
	_va := make([]interface{}, len(prop))
	for _i := range prop {
		_va[_i] = prop[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, text)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}
