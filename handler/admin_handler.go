package handler

import (
	"net/http"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/service"
	"github.com/Amierza/TedXBackend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		// Speaker
		CreateSpeaker(ctx *gin.Context)
		GetAllSpeaker(ctx *gin.Context)
		GetDetailSpeaker(ctx *gin.Context)
		UpdateSpeaker(ctx *gin.Context)
		DeleteSpeaker(ctx *gin.Context)

		// Merch
		CreateMerch(ctx *gin.Context)
		GetAllMerch(ctx *gin.Context)
		GetDetailMerch(ctx *gin.Context)
		UpdateMerch(ctx *gin.Context)
		DeleteMerch(ctx *gin.Context)
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
	fileHeader, err := ctx.FormFile("sponsorship_image")
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

	payload.Category = ctx.PostForm("sponsorship_cat")
	payload.Name = ctx.PostForm("sponsorship_name")

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
	result, err := ah.adminService.GetAllSponsorship(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_SPONSORSHIP, result)
	ctx.JSON(http.StatusOK, res)
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
	fileHeader, err := ctx.FormFile("sponsorship_image")
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

	payload.Category = ctx.PostForm("sponsorship_cat")
	payload.Name = ctx.PostForm("sponsorship_name")

	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.UpdateSponsorship(ctx, payload)
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

	result, err := ah.adminService.DeleteSponsorship(ctx, payload)
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
	payload.Description = ctx.PostForm("speaker_desc")

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
		result, err := ah.adminService.GetAllSpeakerNoPagination(ctx)
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

	result, err := ah.adminService.GetAllSpeakerWithPagination(ctx, payload)
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
	payload.ID = idStr
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
	payload.Description = ctx.PostForm("speaker_desc")

	result, err := ah.adminService.UpdateSpeaker(ctx, payload)
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

// Merch
func (ah *AdminHandler) CreateMerch(ctx *gin.Context) {
	var payload dto.CreateMerchRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	files := form.File["merch_images"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		defer file.Close()

		payload.Images = append(payload.Images, dto.ImageUpload{
			FileHeader: fileHeader,
			FileReader: file,
		})
	}

	result, err := ah.adminService.CreateMerch(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_MERCH, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_MERCH, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllMerch(ctx *gin.Context) {
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllMerchNoPagination(ctx)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_MERCH, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_MERCH, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllMerchWithPagination(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_MERCH, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_MERCH,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetDetailMerch(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailMerch(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_MERCH, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_MERCH, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) UpdateMerch(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.UpdateMerchRequest
	payload.ID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	files := form.File["merch_images"]
	targetIDs := form.Value["target_image_id"]
	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		var targetImageID uuid.UUID
		if len(targetIDs) > i {
			targetImageID, err = uuid.Parse(targetIDs[i])
			if err != nil {
				file.Close()
				res := utils.BuildResponseFailed("Invalid target image ID", err.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
				return
			}
		}

		payload.ReplaceImages = append(payload.ReplaceImages, dto.ReplaceImageUpload{
			TargetImageID: targetImageID,
			FileHeader:    fileHeader,
			FileReader:    file,
		})
	}

	result, err := ah.adminService.UpdateMerch(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_MERCH, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_MERCH, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) DeleteMerch(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.DeleteMerchRequest
	payload.MerchID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.DeleteMerch(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_MERCH, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_DELETE_MERCH, result)
	ctx.JSON(http.StatusOK, res)
}
