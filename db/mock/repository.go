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

// CreateFavorite mocks base method.
func (m *MockRepository) CreateFavorite(arg0 context.Context, arg1 db.CreateFavoriteParams) (db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFavorite indicates an expected call of CreateFavorite.
func (mr *MockRepositoryMockRecorder) CreateFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFavorite", reflect.TypeOf((*MockRepository)(nil).CreateFavorite), arg0, arg1)
}

// CreateFollow mocks base method.
func (m *MockRepository) CreateFollow(arg0 context.Context, arg1 db.CreateFollowParams) (db.Follow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFollow", arg0, arg1)
	ret0, _ := ret[0].(db.Follow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFollow indicates an expected call of CreateFollow.
func (mr *MockRepositoryMockRecorder) CreateFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFollow", reflect.TypeOf((*MockRepository)(nil).CreateFollow), arg0, arg1)
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
func (m *MockRepository) CreateSubCategory(arg0 context.Context, arg1 db.CreateSubCategoryParams) (db.SubCategory, error) {
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

// DeleteFavorite mocks base method.
func (m *MockRepository) DeleteFavorite(arg0 context.Context, arg1 db.DeleteFavoriteParams) (db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFavorite indicates an expected call of DeleteFavorite.
func (mr *MockRepositoryMockRecorder) DeleteFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavorite", reflect.TypeOf((*MockRepository)(nil).DeleteFavorite), arg0, arg1)
}

// DeleteFollow mocks base method.
func (m *MockRepository) DeleteFollow(arg0 context.Context, arg1 db.DeleteFollowParams) (db.Follow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFollow", arg0, arg1)
	ret0, _ := ret[0].(db.Follow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFollow indicates an expected call of DeleteFollow.
func (mr *MockRepositoryMockRecorder) DeleteFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFollow", reflect.TypeOf((*MockRepository)(nil).DeleteFollow), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockRepository) DeletePost(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockRepositoryMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockRepository)(nil).DeletePost), arg0, arg1)
}

// DeletePostSubCategory mocks base method.
func (m *MockRepository) DeletePostSubCategory(arg0 context.Context, arg1 db.DeletePostSubCategoryParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostSubCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostSubCategory indicates an expected call of DeletePostSubCategory.
func (mr *MockRepositoryMockRecorder) DeletePostSubCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostSubCategory", reflect.TypeOf((*MockRepository)(nil).DeletePostSubCategory), arg0, arg1)
}

// DeleteSubCategory mocks base method.
func (m *MockRepository) DeleteSubCategory(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubCategory indicates an expected call of DeleteSubCategory.
func (mr *MockRepositoryMockRecorder) DeleteSubCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubCategory", reflect.TypeOf((*MockRepository)(nil).DeleteSubCategory), arg0, arg1)
}

// GetDailyActivitySummaryByMonth mocks base method.
func (m *MockRepository) GetDailyActivitySummaryByMonth(arg0 context.Context, arg1 db.GetDailyActivitySummaryByMonthParams) ([]db.GetDailyActivitySummaryByMonthRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDailyActivitySummaryByMonth", arg0, arg1)
	ret0, _ := ret[0].([]db.GetDailyActivitySummaryByMonthRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDailyActivitySummaryByMonth indicates an expected call of GetDailyActivitySummaryByMonth.
func (mr *MockRepositoryMockRecorder) GetDailyActivitySummaryByMonth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDailyActivitySummaryByMonth", reflect.TypeOf((*MockRepository)(nil).GetDailyActivitySummaryByMonth), arg0, arg1)
}

// GetExpenseSummaryByMonth mocks base method.
func (m *MockRepository) GetExpenseSummaryByMonth(arg0 context.Context, arg1 db.GetExpenseSummaryByMonthParams) ([]db.GetExpenseSummaryByMonthRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpenseSummaryByMonth", arg0, arg1)
	ret0, _ := ret[0].([]db.GetExpenseSummaryByMonthRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpenseSummaryByMonth indicates an expected call of GetExpenseSummaryByMonth.
func (mr *MockRepositoryMockRecorder) GetExpenseSummaryByMonth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpenseSummaryByMonth", reflect.TypeOf((*MockRepository)(nil).GetExpenseSummaryByMonth), arg0, arg1)
}

// GetFavorite mocks base method.
func (m *MockRepository) GetFavorite(arg0 context.Context, arg1 db.GetFavoriteParams) (db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorite indicates an expected call of GetFavorite.
func (mr *MockRepositoryMockRecorder) GetFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorite", reflect.TypeOf((*MockRepository)(nil).GetFavorite), arg0, arg1)
}

// GetFollow mocks base method.
func (m *MockRepository) GetFollow(arg0 context.Context, arg1 db.GetFollowParams) (db.Follow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollow", arg0, arg1)
	ret0, _ := ret[0].(db.Follow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollow indicates an expected call of GetFollow.
func (mr *MockRepositoryMockRecorder) GetFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollow", reflect.TypeOf((*MockRepository)(nil).GetFollow), arg0, arg1)
}

// GetFollowCount mocks base method.
func (m *MockRepository) GetFollowCount(arg0 context.Context, arg1 db.GetFollowCountParams) (db.GetFollowCountRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowCount", arg0, arg1)
	ret0, _ := ret[0].(db.GetFollowCountRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowCount indicates an expected call of GetFollowCount.
func (mr *MockRepositoryMockRecorder) GetFollowCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowCount", reflect.TypeOf((*MockRepository)(nil).GetFollowCount), arg0, arg1)
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

// GetOwnPosts mocks base method.
func (m *MockRepository) GetOwnPosts(arg0 context.Context, arg1 string) ([]db.GetOwnPostsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.GetOwnPostsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwnPosts indicates an expected call of GetOwnPosts.
func (mr *MockRepositoryMockRecorder) GetOwnPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnPosts", reflect.TypeOf((*MockRepository)(nil).GetOwnPosts), arg0, arg1)
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

// GetPostsByMainCategory mocks base method.
func (m *MockRepository) GetPostsByMainCategory(arg0 context.Context, arg1 db.GetPostsByMainCategoryParams) ([]db.GetPostsByMainCategoryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByMainCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.GetPostsByMainCategoryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByMainCategory indicates an expected call of GetPostsByMainCategory.
func (mr *MockRepositoryMockRecorder) GetPostsByMainCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByMainCategory", reflect.TypeOf((*MockRepository)(nil).GetPostsByMainCategory), arg0, arg1)
}

// GetPostsByMainCategoryAndSubCategory mocks base method.
func (m *MockRepository) GetPostsByMainCategoryAndSubCategory(arg0 context.Context, arg1 db.GetPostsByMainCategoryAndSubCategoryParams) ([]db.GetPostsByMainCategoryAndSubCategoryRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByMainCategoryAndSubCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.GetPostsByMainCategoryAndSubCategoryRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByMainCategoryAndSubCategory indicates an expected call of GetPostsByMainCategoryAndSubCategory.
func (mr *MockRepositoryMockRecorder) GetPostsByMainCategoryAndSubCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByMainCategoryAndSubCategory", reflect.TypeOf((*MockRepository)(nil).GetPostsByMainCategoryAndSubCategory), arg0, arg1)
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
func (m *MockRepository) GetPostsByUid(arg0 context.Context, arg1 db.GetPostsByUidParams) ([]db.GetPostsByUidRow, error) {
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

// GetPostsCount mocks base method.
func (m *MockRepository) GetPostsCount(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsCount", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsCount indicates an expected call of GetPostsCount.
func (mr *MockRepositoryMockRecorder) GetPostsCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsCount", reflect.TypeOf((*MockRepository)(nil).GetPostsCount), arg0, arg1)
}

// GetSubCategoryByName mocks base method.
func (m *MockRepository) GetSubCategoryByName(arg0 context.Context, arg1 string) (db.SubCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubCategoryByName", arg0, arg1)
	ret0, _ := ret[0].(db.SubCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubCategoryByName indicates an expected call of GetSubCategoryByName.
func (mr *MockRepositoryMockRecorder) GetSubCategoryByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubCategoryByName", reflect.TypeOf((*MockRepository)(nil).GetSubCategoryByName), arg0, arg1)
}

// GetSubCategorySummaryByMonth mocks base method.
func (m *MockRepository) GetSubCategorySummaryByMonth(arg0 context.Context, arg1 db.GetSubCategorySummaryByMonthParams) ([]db.GetSubCategorySummaryByMonthRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubCategorySummaryByMonth", arg0, arg1)
	ret0, _ := ret[0].([]db.GetSubCategorySummaryByMonthRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubCategorySummaryByMonth indicates an expected call of GetSubCategorySummaryByMonth.
func (mr *MockRepositoryMockRecorder) GetSubCategorySummaryByMonth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubCategorySummaryByMonth", reflect.TypeOf((*MockRepository)(nil).GetSubCategorySummaryByMonth), arg0, arg1)
}

// GetTimelinePosts mocks base method.
func (m *MockRepository) GetTimelinePosts(arg0 context.Context, arg1 string) ([]db.GetTimelinePostsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimelinePosts", arg0, arg1)
	ret0, _ := ret[0].([]db.GetTimelinePostsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimelinePosts indicates an expected call of GetTimelinePosts.
func (mr *MockRepositoryMockRecorder) GetTimelinePosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimelinePosts", reflect.TypeOf((*MockRepository)(nil).GetTimelinePosts), arg0, arg1)
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
