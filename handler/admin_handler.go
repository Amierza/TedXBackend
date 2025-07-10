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
		// Authentication
		Login(ctx *gin.Context)

		// Sponsorship
		CreateSponsorship(ctx *gin.Context)
		GetAllSponsorship(ctx *gin.Context)
		GetDetailSponsorship(ctx *gin.Context)
		UpdateSponsorship(ctx *gin.Context)
		DeleteSponsorship(ctx *gin.Context)

		// Sponsorship
		CreateSpeaker(ctx *gin.Context)
		GetAllSpeaker(ctx *gin.Context)
		GetDetailSpeaker(ctx *gin.Context)
		UpdateSpeaker(ctx *gin.Context)
		DeleteSpeaker(ctx *gin.Context)
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

// Authentication
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

// Sponsorship
func (ah *AdminHandler) CreateSponsorship(ctx *gin.Context) {
	var payload dto.CreateSponsorshipRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.CreateSponsorship(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_SPONSORSHIP, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllSponsorship(ctx *gin.Context) {
	result, err := ah.adminService.GetAllSponsorship(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_SPONSORSHIP, result)
	ctx.JSON(http.StatusOK, res)
	return
}
func (ah *AdminHandler) GetDetailSponsorship(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailSponsorship(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_SPONSORSHIP, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) UpdateSponsorship(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.UpdateSponsorshipRequest
	payload.ID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.UpdateSponsorship(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_SPONSORSHIP, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) DeleteSponsorship(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.DeleteSponsorshipRequest
	payload.SponsorshipID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.DeleteSponsorship(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_DELETE_SPONSORSHIP, result)
	ctx.JSON(http.StatusOK, res)
}

// Spekear
func (ah *AdminHandler) CreateSpeaker(ctx *gin.Context) {
	var payload dto.CreateSpeakerRequest
	fileHeader, err := ctx.FormFile("speaker_image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		defer file.Close()

		payload.FileHeader = fileHeader
		payload.FileReader = file
	}

	payload.Name = ctx.PostForm("speaker_name")

	result, err := ah.adminService.CreateSpeaker(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_SPEAKER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_SPEAKER, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllSpeaker(ctx *gin.Context) {
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllSpeakerNoPagination(ctx.Request.Context())
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPEAKER, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_SPEAKER, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllSpeakerWithPagination(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPEAKER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_SPEAKER,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetDetailSpeaker(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailSpeaker(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_SPEAKER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_SPEAKER, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) UpdateSpeaker(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.UpdateSpeakerRequest
	fileHeader, err := ctx.FormFile("speaker_image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		defer file.Close()

		payload.FileHeader = fileHeader
		payload.FileReader = file
	}

	payload.Name = ctx.PostForm("speaker_name")
	payload.ID = idStr

	result, err := ah.adminService.UpdateSpeaker(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_SPEAKER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_SPEAKER, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) DeleteSpeaker(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.DeleteSpeakerRequest
	payload.SpeakerID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.DeleteSpeaker(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_SPEAKER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_DELETE_SPEAKER, result)
	ctx.JSON(http.StatusOK, res)
}
