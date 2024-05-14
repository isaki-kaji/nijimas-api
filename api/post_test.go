package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/api/controller"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/service"
	mockservice "github.com/isaki-kaji/nijimas-api/service/mock"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	testPost := randomNewPost()

	testCases := []struct {
		name       string
		body       gin.H
		buildStubs func(service *mockservice.MockPostService)
		check      func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"post_id":        testPost.PostID,
				"uid":            testPost.Uid,
				"main_category":  testPost.MainCategory,
				"post_text":      testPost.PostText,
				"photo_url":      testPost.PhotoUrl,
				"expense":        testPost.Expense,
				"location":       testPost.Location,
				"public_type_no": testPost.PublicTypeNo,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(testPost, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "OK: not required expense and location",
			body: gin.H{
				"post_id":        testPost.PostID,
				"uid":            testPost.Uid,
				"main_category":  testPost.MainCategory,
				"post_text":      testPost.PostText,
				"photo_url":      testPost.PhotoUrl,
				"public_type_no": testPost.PublicTypeNo,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(testPost, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "OK: not required post_text, expense and location",
			body: gin.H{
				"post_id":        testPost.PostID,
				"uid":            testPost.Uid,
				"main_category":  testPost.MainCategory,
				"photo_url":      testPost.PhotoUrl,
				"public_type_no": testPost.PublicTypeNo,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(testPost, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "BadRequest(post_id required)",
			body: gin.H{
				"uid":            testPost.Uid,
				"main_category":  testPost.MainCategory,
				"post_text":      testPost.PostText,
				"photo_url":      testPost.PhotoUrl,
				"public_type_no": testPost.PublicTypeNo,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest(uid required)",
			body: gin.H{
				"post_id":        testPost.PostID,
				"main_category":  testPost.MainCategory,
				"post_text":      testPost.PostText,
				"photo_url":      testPost.PhotoUrl,
				"public_type_no": testPost.PublicTypeNo,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest(public_type_no required)",
			body: gin.H{
				"post_id":       testPost.PostID,
				"uid":           testPost.Uid,
				"main_category": testPost.MainCategory,
				"post_text":     testPost.PostText,
				"photo_url":     testPost.PhotoUrl,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest(expense over limit)",
			body: gin.H{
				"post_id":        testPost.PostID,
				"uid":            testPost.Uid,
				"main_category":  testPost.MainCategory,
				"post_text":      testPost.PostText,
				"expense":        100000000,
				"photo_url":      testPost.PhotoUrl,
				"public_type_no": testPost.PublicTypeNo,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalServerError(other error)",
			body: gin.H{
				"post_id":        testPost.PostID,
				"uid":            testPost.Uid,
				"main_category":  testPost.MainCategory,
				"post_text":      testPost.PostText,
				"expense":        testPost.Expense,
				"photo_url":      testPost.PhotoUrl,
				"public_type_no": testPost.PublicTypeNo,
			},
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(1).Return(db.Post{}, errors.New("other error"))
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postService := mockservice.NewMockPostService(ctrl)
			tc.buildStubs(postService)

			server := NewTestPostServer(t, postService)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/posts"
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, req)
			tc.check(recorder)
		})
	}
}

func TestGetPostsByUid(t *testing.T) {
	testUid := util.RandomUid()
	testPosts := randomNewFullPosts()
	testCases := []struct {
		name       string
		uid        string
		buildStubs func(service *mockservice.MockPostService)
		check      func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			uid:  testUid,
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					GetPostsByUid(gomock.Any(), testUid).
					Times(1).
					Return(testPosts, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				var posts []db.GetPostsByUidRow
				err := json.NewDecoder(recorder.Body).Decode(&posts)
				require.NoError(t, err)

				require.Len(t, posts, len(testPosts))
				for i, post := range posts {
					require.Equal(t, testPosts[i].PostID, post.PostID)
					require.Equal(t, testPosts[i].Uid, post.Uid)
					require.Equal(t, testPosts[i].Username, post.Username)
					require.Equal(t, testPosts[i].MainCategory, post.MainCategory)
					require.Equal(t, testPosts[i].SubCategory, post.SubCategory)
					require.Equal(t, testPosts[i].SubCategory_2, post.SubCategory_2)
					require.Equal(t, testPosts[i].PostText, post.PostText)
					require.Equal(t, testPosts[i].PhotoUrl, post.PhotoUrl)
					require.Equal(t, testPosts[i].Expense, post.Expense)
					require.Equal(t, testPosts[i].Location, post.Location)
					require.Equal(t, testPosts[i].PublicTypeNo, post.PublicTypeNo)
				}
			},
		},
		{
			name: "OK: no posts",
			uid:  testUid,
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					GetPostsByUid(gomock.Any(), testUid).
					Times(1).
					Return([]db.GetPostsByUidRow{}, nil)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Equal(t, "[]", recorder.Body.String())
			},
		},
		{
			name: "InternalServerError(other error)",
			uid:  testUid,
			buildStubs: func(service *mockservice.MockPostService) {
				service.EXPECT().
					GetPostsByUid(gomock.Any(), testUid).
					Times(1).
					Return(nil, errors.New("other error"))
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

			postService := mockservice.NewMockPostService(ctrl)
			tc.buildStubs(postService)

			server := NewTestPostServer(t, postService)
			recorder := httptest.NewRecorder()

			url := "/posts?uid=" + testUid
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, req)
			tc.check(recorder)
		})
	}
}

func randomNewPost() (post db.Post) {
	uuid := uuid.New()

	post = db.Post{
		PostID:       uuid,
		Uid:          util.RandomUid(),
		MainCategory: util.RandomMainCategory(),
		PostText:     util.ToPointerOrNil(util.RandomString(50)),
		PhotoUrl:     util.ToPointerOrNil(util.RandomString(100)),
		Expense:      util.ToPointerOrNil(rand.Int63n(100000)),
		Location:     util.ToPointerOrNil(util.RandomString(20)),
		PublicTypeNo: util.RandomPublicTypeNo(),
	}
	return
}

func randomNewFullPosts() (posts []db.GetPostsByUidRow) {
	n := 5
	for i := 0; i < n; i++ {
		post := db.GetPostsByUidRow{
			PostID:        uuid.New(),
			Username:      util.RandomString(10),
			MainCategory:  util.RandomMainCategory(),
			SubCategory:   util.ToPointerOrNil(util.RandomString(10)),
			SubCategory_2: util.ToPointerOrNil(util.RandomString(10)),
			PostText:      util.ToPointerOrNil(util.RandomString(50)),
			PhotoUrl:      util.ToPointerOrNil(util.RandomString(100)),
			Expense:       util.ToPointerOrNil(rand.Int63n(100000)),
			Location:      util.ToPointerOrNil(util.RandomString(20)),
			PublicTypeNo:  util.RandomPublicTypeNo(),
		}
		posts = append(posts, post)
	}
	return
}

func NewTestPostServer(t *testing.T, postService service.PostService) *Server {
	config, err := util.LoadConfig("..")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	postController := NewTestPostController(postService)
	testRouter := NewTestPostRouter(postController)
	server, err := NewServer(config, testRouter)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	return server
}

func NewTestPostRouter(postController *controller.PostController) *gin.Engine {
	router := gin.Default()

	router.POST("/posts", postController.CreatePost)
	router.GET("/posts", postController.GetPostsByUid)

	return router
}

func NewTestPostController(postService service.PostService) *controller.PostController {
	return controller.NewPostController(postService)
}
