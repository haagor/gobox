// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/haagor/gobox/katas/birthdayGreetings/identity/adapter (interfaces: DBAdapter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/haagor/gobox/katas/birthdayGreetings/identity/entity"
	reflect "reflect"
	time "time"
)

// MockDBAdapter is a mock of DBAdapter interface
type MockDBAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockDBAdapterMockRecorder
}

// MockDBAdapterMockRecorder is the mock recorder for MockDBAdapter
type MockDBAdapterMockRecorder struct {
	mock *MockDBAdapter
}

// NewMockDBAdapter creates a new mock instance
func NewMockDBAdapter(ctrl *gomock.Controller) *MockDBAdapter {
	mock := &MockDBAdapter{ctrl: ctrl}
	mock.recorder = &MockDBAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBAdapter) EXPECT() *MockDBAdapterMockRecorder {
	return m.recorder
}

// GetFriendsByBirthDate mocks base method
func (m *MockDBAdapter) GetFriendsByBirthDate(arg0 time.Time) []entity.Friend {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendsByBirthDate", arg0)
	ret0, _ := ret[0].([]entity.Friend)
	return ret0
}

// GetFriendsByBirthDate indicates an expected call of GetFriendsByBirthDate
func (mr *MockDBAdapterMockRecorder) GetFriendsByBirthDate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendsByBirthDate", reflect.TypeOf((*MockDBAdapter)(nil).GetFriendsByBirthDate), arg0)
}

// getAllFriends mocks base method
func (m *MockDBAdapter) getAllFriends() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "getAllFriends")
}

// getAllFriends indicates an expected call of getAllFriends
func (mr *MockDBAdapterMockRecorder) getAllFriends() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAllFriends", reflect.TypeOf((*MockDBAdapter)(nil).getAllFriends))
}

// setFriend mocks base method
func (m *MockDBAdapter) setFriend(arg0 entity.Friend) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setFriend", arg0)
}

// setFriend indicates an expected call of setFriend
func (mr *MockDBAdapterMockRecorder) setFriend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setFriend", reflect.TypeOf((*MockDBAdapter)(nil).setFriend), arg0)
}