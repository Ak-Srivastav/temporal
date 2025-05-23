// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go
//
// Generated by this command:
//
//	mockgen -copyright_file ../../../LICENSE -package events -source cache.go -destination events_cache_mock.go
//

// Package events is a generated GoMock package.
package events

import (
	context "context"
	reflect "reflect"

	history "go.temporal.io/api/history/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
	isgomock struct{}
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

// DeleteEvent mocks base method.
func (m *MockCache) DeleteEvent(key EventKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteEvent", key)
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockCacheMockRecorder) DeleteEvent(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockCache)(nil).DeleteEvent), key)
}

// GetEvent mocks base method.
func (m *MockCache) GetEvent(ctx context.Context, shardID int32, key EventKey, firstEventID int64, branchToken []byte) (*history.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", ctx, shardID, key, firstEventID, branchToken)
	ret0, _ := ret[0].(*history.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockCacheMockRecorder) GetEvent(ctx, shardID, key, firstEventID, branchToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockCache)(nil).GetEvent), ctx, shardID, key, firstEventID, branchToken)
}

// PutEvent mocks base method.
func (m *MockCache) PutEvent(key EventKey, event *history.HistoryEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutEvent", key, event)
}

// PutEvent indicates an expected call of PutEvent.
func (mr *MockCacheMockRecorder) PutEvent(key, event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutEvent", reflect.TypeOf((*MockCache)(nil).PutEvent), key, event)
}
