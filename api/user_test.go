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
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
	mockservice "github.com/isaki-kaji/nijimas-api/service/mock"
	"github.com/isaki-kaji/nijimas-api/util"
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
			}}, {
			name: "NG: BadRequest(uid required)",
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
			}}, {
			name: "NG: BadRequest(uid required)",
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
			name: "NG: BadRequest(username required)",
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
			name: "NG: BadRequest(username character limit exceeded)",
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
			name: "NG: BadRequest(country_code must be 2 characters)",
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
		}, {
			name: "NG: Conflict(user already exists)",
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
		}, {
			name: "NG: Conflict(other error)",
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

func randomNewUser() (user db.User) {
	user = db.User{
		Uid:         util.RandomUid(),
		Username:    util.RandomString(5),
		CountryCode: util.RandomCountryCode(),
	}
	return
}

func NewTestUserServer(t *testing.T, userService domain.UserService) *Server {
	config, err := util.LoadConfig("..")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	userController := NewTestUserController(userService)
	testRouter := NewTestRouter(userController)
	server, err := NewServer(config, testRouter)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	return server
}

func NewTestRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/users", userController.Create)
	router.GET("/users/:id", userController.Get)

	return router
}

func NewTestUserController(userService domain.UserService) *controller.UserController {
	return controller.NewUserController(userService)
}
