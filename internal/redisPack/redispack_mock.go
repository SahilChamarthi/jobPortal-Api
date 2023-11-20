// Code generated by MockGen. DO NOT EDIT.
// Source: intilizeRedis.go
//
// Generated by this command:
//
//	mockgen -source intilizeRedis.go -destination redispack_mock.go -package redispack
//
// Package redispack is a generated GoMock package.
package redispack

import (
	model "project/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// CheckRedisKey mocks base method.
func (m *MockCache) CheckRedisKey(key string) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRedisKey", key)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRedisKey indicates an expected call of CheckRedisKey.
func (mr *MockCacheMockRecorder) CheckRedisKey(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRedisKey", reflect.TypeOf((*MockCache)(nil).CheckRedisKey), key)
}

// SetRedisKey mocks base method.
func (m *MockCache) SetRedisKey(key string, jobData model.Job) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRedisKey", key, jobData)
}

// SetRedisKey indicates an expected call of SetRedisKey.
func (mr *MockCacheMockRecorder) SetRedisKey(key, jobData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRedisKey", reflect.TypeOf((*MockCache)(nil).SetRedisKey), key, jobData)
}
