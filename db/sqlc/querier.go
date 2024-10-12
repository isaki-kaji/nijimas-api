// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateFavorite(ctx context.Context, arg CreateFavoriteParams) (Favorite, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreatePostSubCategory(ctx context.Context, arg CreatePostSubCategoryParams) (PostSubcategory, error)
	CreateSubCategory(ctx context.Context, arg CreateSubCategoryParams) (SubCategory, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteFavorite(ctx context.Context, arg DeleteFavoriteParams) (Favorite, error)
	DeletePost(ctx context.Context, postID uuid.UUID) error
	DeletePostSubCategory(ctx context.Context, arg DeletePostSubCategoryParams) error
	DeleteSubCategory(ctx context.Context, categoryID uuid.UUID) error
	GetDailyActivitySummaryByMonth(ctx context.Context, arg GetDailyActivitySummaryByMonthParams) ([]GetDailyActivitySummaryByMonthRow, error)
	GetExpenseSummaryByMonth(ctx context.Context, arg GetExpenseSummaryByMonthParams) ([]GetExpenseSummaryByMonthRow, error)
	GetFavorite(ctx context.Context, arg GetFavoriteParams) (Favorite, error)
	GetFollowUsers(ctx context.Context, uid string) ([]GetFollowUsersRow, error)
	GetOwnPosts(ctx context.Context, uid string) ([]GetOwnPostsRow, error)
	GetPostById(ctx context.Context, postID uuid.UUID) (GetPostByIdRow, error)
	GetPostsByUid(ctx context.Context, arg GetPostsByUidParams) ([]GetPostsByUidRow, error)
	GetSubCategoryByName(ctx context.Context, categoryName string) (SubCategory, error)
	GetSubCategorySummaryByMonth(ctx context.Context, arg GetSubCategorySummaryByMonthParams) ([]GetSubCategorySummaryByMonthRow, error)
	GetUser(ctx context.Context, uid string) (User, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
