package route

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
	"github.com/isaki-kaji/nijimas-api/api/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRouter),
)

func NewRouter(
	UserController *controller.UserController,
	PostController *controller.PostController,
	FavoriteController *controller.FavoriteController,
	SummaryController *controller.SummaryController,
	PostSearchController *controller.PostSearchController,
	FollowController *controller.FollowController,
	FollowRequestController *controller.FollowRequestController,
	SubCategoryController *controller.SubCategoryController,
	AuthClient *auth.Client,
) *gin.Engine {
	router := gin.Default()
	authRouter := router.Group("/").Use(middleware.LoggingMiddleware(), middleware.AuthMiddleware(AuthClient))

	NewUserRouter(router, authRouter, UserController)
	NewPostRouter(router, authRouter, PostController)
	NewFavoriteRouter(router, authRouter, FavoriteController)
	NewSummaryRouter(router, authRouter, SummaryController)
	NewPostSearchRouter(router, authRouter, PostSearchController)
	NewFollowRouter(router, authRouter, FollowController)
	NewFollowRequestRouter(router, authRouter, FollowRequestController)
	NewSubCategoryRouter(router, authRouter, SubCategoryController)

	return router
}
