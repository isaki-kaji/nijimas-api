package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/application/service"
)

type SubCategoryController struct {
	service service.SubCategoryService
}

func NewSubCategoryController(service service.SubCategoryService) *SubCategoryController {
	return &SubCategoryController{service: service}
}

func (c *SubCategoryController) GetUserUsedSubCategories(ctx *gin.Context) {
	ownUid, err := checkUid(ctx)
	if err != nil {
		return
	}

	frs, err := c.service.GetUserUsedSubCategories(ctx, ownUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, frs)
}
