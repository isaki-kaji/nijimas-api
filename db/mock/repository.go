// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/isaki-kaji/nijimas-api/db/sqlc (interfaces: Repository)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateMainCategory mocks base method.
func (m *MockRepository) CreateMainCategory(arg0 context.Context, arg1 string) (db.MainCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMainCategory", arg0, arg1)
	ret0, _ := ret[0].(db.MainCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMainCategory indicates an expected call of CreateMainCategory.
func (mr *MockRepositoryMockRecorder) CreateMainCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMainCategory", reflect.TypeOf((*MockRepository)(nil).CreateMainCategory), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockRepository) CreatePost(arg0 context.Context, arg1 db.CreatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockRepositoryMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockRepository)(nil).CreatePost), arg0, arg1)
}

// CreatePostSubCategory mocks base method.
func (m *MockRepository) CreatePostSubCategory(arg0 context.Context, arg1 db.CreatePostSubCategoryParams) (db.PostSubcategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostSubCategory", arg0, arg1)
	ret0, _ := ret[0].(db.PostSubcategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostSubCategory indicates an expected call of CreatePostSubCategory.
func (mr *MockRepositoryMockRecorder) CreatePostSubCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostSubCategory", reflect.TypeOf((*MockRepository)(nil).CreatePostSubCategory), arg0, arg1)
}

// CreatePostTx mocks base method.
func (m *MockRepository) CreatePostTx(arg0 context.Context, arg1 db.CreatePostTxParam) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostTx", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostTx indicates an expected call of CreatePostTx.
func (mr *MockRepositoryMockRecorder) CreatePostTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostTx", reflect.TypeOf((*MockRepository)(nil).CreatePostTx), arg0, arg1)
}

// CreateSubCategory mocks base method.
func (m *MockRepository) CreateSubCategory(arg0 context.Context, arg1 string) (db.SubCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubCategory", arg0, arg1)
	ret0, _ := ret[0].(db.SubCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubCategory indicates an expected call of CreateSubCategory.
func (mr *MockRepositoryMockRecorder) CreateSubCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubCategory", reflect.TypeOf((*MockRepository)(nil).CreateSubCategory), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), arg0, arg1)
}

// GetFollowUsers mocks base method.
func (m *MockRepository) GetFollowUsers(arg0 context.Context, arg1 string) ([]db.GetFollowUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.GetFollowUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowUsers indicates an expected call of GetFollowUsers.
func (mr *MockRepositoryMockRecorder) GetFollowUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowUsers", reflect.TypeOf((*MockRepository)(nil).GetFollowUsers), arg0, arg1)
}

// GetMainCategories mocks base method.
func (m *MockRepository) GetMainCategories(arg0 context.Context) ([]db.MainCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainCategories", arg0)
	ret0, _ := ret[0].([]db.MainCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMainCategories indicates an expected call of GetMainCategories.
func (mr *MockRepositoryMockRecorder) GetMainCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainCategories", reflect.TypeOf((*MockRepository)(nil).GetMainCategories), arg0)
}

// GetMainCategory mocks base method.
func (m *MockRepository) GetMainCategory(arg0 context.Context, arg1 string) (db.MainCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainCategory", arg0, arg1)
	ret0, _ := ret[0].(db.MainCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMainCategory indicates an expected call of GetMainCategory.
func (mr *MockRepositoryMockRecorder) GetMainCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainCategory", reflect.TypeOf((*MockRepository)(nil).GetMainCategory), arg0, arg1)
}

// GetPostById mocks base method.
func (m *MockRepository) GetPostById(arg0 context.Context, arg1 uuid.UUID) (db.GetPostByIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostById", arg0, arg1)
	ret0, _ := ret[0].(db.GetPostByIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostById indicates an expected call of GetPostById.
func (mr *MockRepositoryMockRecorder) GetPostById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostById", reflect.TypeOf((*MockRepository)(nil).GetPostById), arg0, arg1)
}

// GetPostSubCategoryByPostId mocks base method.
func (m *MockRepository) GetPostSubCategoryByPostId(arg0 context.Context, arg1 uuid.UUID) ([]db.PostSubcategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostSubCategoryByPostId", arg0, arg1)
	ret0, _ := ret[0].([]db.PostSubcategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostSubCategoryByPostId indicates an expected call of GetPostSubCategoryByPostId.
func (mr *MockRepositoryMockRecorder) GetPostSubCategoryByPostId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostSubCategoryByPostId", reflect.TypeOf((*MockRepository)(nil).GetPostSubCategoryByPostId), arg0, arg1)
}

// GetPostsByCategory mocks base method.
func (m *MockRepository) GetPostsByCategory(arg0 context.Context, arg1 db.GetPostsByCategoryParams) ([]db.GetPostsByCategoryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.GetPostsByCategoryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByCategory indicates an expected call of GetPostsByCategory.
func (mr *MockRepositoryMockRecorder) GetPostsByCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByCategory", reflect.TypeOf((*MockRepository)(nil).GetPostsByCategory), arg0, arg1)
}

// GetPostsByFollowing mocks base method.
func (m *MockRepository) GetPostsByFollowing(arg0 context.Context, arg1 string) ([]db.GetPostsByFollowingRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByFollowing", arg0, arg1)
	ret0, _ := ret[0].([]db.GetPostsByFollowingRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByFollowing indicates an expected call of GetPostsByFollowing.
func (mr *MockRepositoryMockRecorder) GetPostsByFollowing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByFollowing", reflect.TypeOf((*MockRepository)(nil).GetPostsByFollowing), arg0, arg1)
}

// GetPostsBySubCategory mocks base method.
func (m *MockRepository) GetPostsBySubCategory(arg0 context.Context, arg1 db.GetPostsBySubCategoryParams) ([]db.GetPostsBySubCategoryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsBySubCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.GetPostsBySubCategoryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsBySubCategory indicates an expected call of GetPostsBySubCategory.
func (mr *MockRepositoryMockRecorder) GetPostsBySubCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsBySubCategory", reflect.TypeOf((*MockRepository)(nil).GetPostsBySubCategory), arg0, arg1)
}

// GetPostsByUid mocks base method.
func (m *MockRepository) GetPostsByUid(arg0 context.Context, arg1 string) ([]db.GetPostsByUidRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByUid", arg0, arg1)
	ret0, _ := ret[0].([]db.GetPostsByUidRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByUid indicates an expected call of GetPostsByUid.
func (mr *MockRepositoryMockRecorder) GetPostsByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByUid", reflect.TypeOf((*MockRepository)(nil).GetPostsByUid), arg0, arg1)
}

// GetSubCategory mocks base method.
func (m *MockRepository) GetSubCategory(arg0 context.Context, arg1 string) (db.SubCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubCategory", arg0, arg1)
	ret0, _ := ret[0].(db.SubCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubCategory indicates an expected call of GetSubCategory.
func (mr *MockRepositoryMockRecorder) GetSubCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubCategory", reflect.TypeOf((*MockRepository)(nil).GetSubCategory), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockRepository) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockRepositoryMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockRepository)(nil).UpdatePost), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockRepository) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepositoryMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepository)(nil).UpdateUser), arg0, arg1)
}