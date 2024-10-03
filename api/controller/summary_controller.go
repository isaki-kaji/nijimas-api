package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/application/service"
)

type SummaryController struct {
	service service.SummaryService
}

func NewSummaryController(service service.SummaryService) *SummaryController {
	return &SummaryController{service: service}
}

func (s *SummaryController) GetMonthlySummary(ctx *gin.Context) {
	yearParam := ctx.Param("year")
	monthParam := ctx.Param("month")

	year, err := strconv.Atoi(yearParam)
	if err != nil {
		err = apperror.BadPathParam.Wrap(err, "year must be a number")
		ctx.JSON(http.StatusBadRequest, apperror.ErrorResponse(ctx, err))
		return
	}
	if year < 2000 || year > 2100 {
		err = apperror.BadPathParam.Wrap(ErrInvalidYear, "year must be between 2000 and 2100")
		ctx.JSON(http.StatusBadRequest, apperror.ErrorResponse(ctx, err))
		return
	}

	month, err := strconv.Atoi(monthParam)
	if err != nil {
		err = apperror.BadPathParam.Wrap(err, "month must be a number")
		ctx.JSON(http.StatusBadRequest, apperror.ErrorResponse(ctx, err))
		return
	}
	if month < 1 || month > 12 {
		err = apperror.BadPathParam.Wrap(ErrInvalidMonth, "month must be between 1 and 12")
		ctx.JSON(http.StatusBadRequest, apperror.ErrorResponse(ctx, err))
		return
	}

	uid, err := checkUid(ctx)
	if err != nil {
		return
	}

	summary, err := s.service.GetMonthlySummary(ctx, uid, year, month)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, summary)
}
