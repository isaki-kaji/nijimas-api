package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/isaki-kaji/nijimas-api/api/controller"
	"github.com/isaki-kaji/nijimas-api/application/service"
	mockservice "github.com/isaki-kaji/nijimas-api/application/service/mock"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	testUser := randomNewUser()

	testCases := []struct {
		name       string
		body       gin.H
		buildStubs func(service *mockservice.MockUserService)
		check      func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"uid":          testUser.Uid,
				"username":     testUser.Username,
				"country_code": testUser.CountryCode,
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(testUser, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "OK: not required country_code",
			body: gin.H{
				"uid":      testUser.Uid,
				"username": testUser.Username,
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(testUser, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			}},
		{
			name: "BadRequest(uid required)",
			body: gin.H{
				"username":     testUser.Username,
				"country_code": testUser.CountryCode,
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest(username required)",
			body: gin.H{
				"uid":          testUser.Uid,
				"country_code": testUser.CountryCode,
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest(username too )",
			body: gin.H{
				"uid":          testUser.Uid,
				"username":     util.RandomString(15),
				"country_code": testUser.CountryCode,
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest(country_code must be 2 characters)",
			body: gin.H{
				"uid":          testUser.Uid,
				"username":     testUser.Username,
				"country_code": util.RandomString(3),
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Conflict(user already exists)",
			body: gin.H{
				"uid":          testUser.Uid,
				"username":     testUser.Username,
				"country_code": testUser.CountryCode,
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, errors.New(util.UserAlreadyExists))
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
		{
			name: "InternalServerError(other error)",
			body: gin.H{
				"uid":          testUser.Uid,
				"username":     testUser.Username,
				"country_code": testUser.CountryCode,
			},
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, errors.New("other error"))
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		}}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mockservice.NewMockUserService(ctrl)
			tc.buildStubs(service)

			server := NewTestUserServer(t, service)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, req)
			tc.check(recorder)
		})

	}
}

func TestGetUser(t *testing.T) {
	testUser := randomNewUser()

	testCases := []struct {
		name       string
		uid        string
		buildStubs func(service *mockservice.MockUserService)
		check      func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			uid:  testUser.Uid,
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(testUser, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NotFound",
			uid:  testUser.Uid,
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(db.User{}, pgx.ErrNoRows)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalServerError(other error)",
			uid:  testUser.Uid,
			buildStubs: func(service *mockservice.MockUserService) {
				service.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(db.User{}, errors.New("other error"))
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		}}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mockservice.NewMockUserService(ctrl)
			tc.buildStubs(service)

			server := NewTestUserServer(t, service)
			recorder := httptest.NewRecorder()

			url := "/users/" + tc.uid
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, req)
			tc.check(recorder)
		})
	}
}

func randomNewUser() (user db.User) {
	user = db.User{
		Uid:         util.RandomUid(),
		Username:    util.RandomString(5),
		CountryCode: util.RandomCountryCode(),
	}
	return
}

func NewTestUserServer(t *testing.T, userService service.UserService) *Server {
	config, err := util.LoadConfig("..")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	userController := NewTestUserController(userService)
	testRouter := NewTestUserRouter(userController)
	server, err := NewServer(config, testRouter)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	return server
}

func NewTestUserRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/users", userController.CreatePost)
	router.GET("/users/:id", userController.GetUserById)

	return router
}

func NewTestUserController(userService service.UserService) *controller.UserController {
	return controller.NewUserController(userService)
}
