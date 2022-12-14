// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	structures "rimeks.ru/services/pkg/app/structures"
)

// MockOngoingMaintenance is a mock of OngoingMaintenance interface.
type MockOngoingMaintenance struct {
	ctrl     *gomock.Controller
	recorder *MockOngoingMaintenanceMockRecorder
}

// MockOngoingMaintenanceMockRecorder is the mock recorder for MockOngoingMaintenance.
type MockOngoingMaintenanceMockRecorder struct {
	mock *MockOngoingMaintenance
}

// NewMockOngoingMaintenance creates a new mock instance.
func NewMockOngoingMaintenance(ctrl *gomock.Controller) *MockOngoingMaintenance {
	mock := &MockOngoingMaintenance{ctrl: ctrl}
	mock.recorder = &MockOngoingMaintenanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOngoingMaintenance) EXPECT() *MockOngoingMaintenanceMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockOngoingMaintenance) Clear() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockOngoingMaintenanceMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockOngoingMaintenance)(nil).Clear))
}

// CreateOrder mocks base method.
func (m *MockOngoingMaintenance) CreateOrder(order structures.OrderList) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOngoingMaintenanceMockRecorder) CreateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOngoingMaintenance)(nil).CreateOrder), order)
}

// DeleteOrder mocks base method.
func (m *MockOngoingMaintenance) DeleteOrder(order structures.OrderList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockOngoingMaintenanceMockRecorder) DeleteOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockOngoingMaintenance)(nil).DeleteOrder), order)
}

// GetAllOrders mocks base method.
func (m *MockOngoingMaintenance) GetAllOrders() ([]structures.OrderList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrders")
	ret0, _ := ret[0].([]structures.OrderList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrders indicates an expected call of GetAllOrders.
func (mr *MockOngoingMaintenanceMockRecorder) GetAllOrders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrders", reflect.TypeOf((*MockOngoingMaintenance)(nil).GetAllOrders))
}

// GetAllOrdersByServiceMarketID mocks base method.
func (m *MockOngoingMaintenance) GetAllOrdersByServiceMarketID(id int) ([]structures.OrderData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrdersByServiceMarketID", id)
	ret0, _ := ret[0].([]structures.OrderData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrdersByServiceMarketID indicates an expected call of GetAllOrdersByServiceMarketID.
func (mr *MockOngoingMaintenanceMockRecorder) GetAllOrdersByServiceMarketID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrdersByServiceMarketID", reflect.TypeOf((*MockOngoingMaintenance)(nil).GetAllOrdersByServiceMarketID), id)
}

// UpdateOrder mocks base method.
func (m *MockOngoingMaintenance) UpdateOrder(order structures.OrderList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockOngoingMaintenanceMockRecorder) UpdateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockOngoingMaintenance)(nil).UpdateOrder), order)
}
