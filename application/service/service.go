package service

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewPostService),
	fx.Provide(NewFavoriteService),
	fx.Provide(NewSummaryService),
	fx.Provide(NewPostSearchService),
	fx.Provide(NewFollowService),
	fx.Provide(NewFollowRequestService),
)
