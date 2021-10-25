// Code generated by MockGen. DO NOT EDIT.
// Source: producer.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockpublisher is a mock of publisher interface.
type Mockpublisher struct {
	ctrl     *gomock.Controller
	recorder *MockpublisherMockRecorder
}

// MockpublisherMockRecorder is the mock recorder for Mockpublisher.
type MockpublisherMockRecorder struct {
	mock *Mockpublisher
}

// NewMockpublisher creates a new mock instance.
func NewMockpublisher(ctrl *gomock.Controller) *Mockpublisher {
	mock := &Mockpublisher{ctrl: ctrl}
	mock.recorder = &MockpublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockpublisher) EXPECT() *MockpublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *Mockpublisher) Publish(topic string, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", topic, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockpublisherMockRecorder) Publish(topic, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*Mockpublisher)(nil).Publish), topic, body)
}
