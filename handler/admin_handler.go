package handler

import (
	"net/http"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/service"
	"github.com/Amierza/TedXBackend/utils"
	"github.com/gin-gonic/gin"
)

type (
	IAdminHandler interface {
		Login(ctx *gin.Context)
	}

	AdminHandler struct {
		adminService service.IAdminService
	}
)

func NewAdminHandler(adminService service.IAdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (ah *AdminHandler) Login(ctx *gin.Context) {
	var payload dto.LoginRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.Login(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN_ADMIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN_ADMIN, result)
	ctx.JSON(http.StatusOK, res)
}
