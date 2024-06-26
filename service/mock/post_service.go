// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/isaki-kaji/nijimas-api/service (interfaces: PostService)

// Package mockservice is a generated GoMock package.
package mockservice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	service "github.com/isaki-kaji/nijimas-api/service"
)

// MockPostService is a mock of PostService interface.
type MockPostService struct {
	ctrl     *gomock.Controller
	recorder *MockPostServiceMockRecorder
}

// MockPostServiceMockRecorder is the mock recorder for MockPostService.
type MockPostServiceMockRecorder struct {
	mock *MockPostService
}

// NewMockPostService creates a new mock instance.
func NewMockPostService(ctrl *gomock.Controller) *MockPostService {
	mock := &MockPostService{ctrl: ctrl}
	mock.recorder = &MockPostServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostService) EXPECT() *MockPostServiceMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPostService) CreatePost(arg0 context.Context, arg1 service.CreatePostRequest) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostServiceMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostService)(nil).CreatePost), arg0, arg1)
}

// GetPostsByUid mocks base method.
func (m *MockPostService) GetPostsByUid(arg0 context.Context, arg1 db.GetPostsByUidParams) ([]service.PostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByUid", arg0, arg1)
	ret0, _ := ret[0].([]service.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByUid indicates an expected call of GetPostsByUid.
func (mr *MockPostServiceMockRecorder) GetPostsByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByUid", reflect.TypeOf((*MockPostService)(nil).GetPostsByUid), arg0, arg1)
}
