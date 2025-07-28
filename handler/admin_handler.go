package handler

import (
	"net/http"
	"strconv"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/entity"
	"github.com/Amierza/TedXBackend/service"
	"github.com/Amierza/TedXBackend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	IAdminHandler interface {
		// Authentication
		Login(ctx *gin.Context)

		// User
		CreateUser(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
		GetDetailUser(ctx *gin.Context)
		UpdateUser(ctx *gin.Context)
		DeleteUser(ctx *gin.Context)

		// Ticket
		CreateTicket(ctx *gin.Context)
		GetAllTicket(ctx *gin.Context)
		GetDetailTicket(ctx *gin.Context)
		UpdateTicket(ctx *gin.Context)
		DeleteTicket(ctx *gin.Context)

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

		// Bundle
		CreateBundle(ctx *gin.Context)
		GetAllBundle(ctx *gin.Context)
		GetDetailBundle(ctx *gin.Context)
		UpdateBundle(ctx *gin.Context)
		DeleteBundle(ctx *gin.Context)

		// Student Ambassador
		CreateStudentAmbassador(ctx *gin.Context)
		GetAllStudentAmbassador(ctx *gin.Context)
		GetDetailStudentAmbassador(ctx *gin.Context)
		UpdateStudentAmbassador(ctx *gin.Context)
		DeleteStudentAmbassador(ctx *gin.Context)

		// Transaction & Ticket Form
		CreateTransactionTicket(ctx *gin.Context)
		GetAllTransactionTicket(ctx *gin.Context)
		GetDetailTransactionTicket(ctx *gin.Context)
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

// User
func (ah *AdminHandler) CreateUser(ctx *gin.Context) {
	var payload dto.CreateUserRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.CreateUser(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_USER, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllUser(ctx *gin.Context) {
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"
	roleName := ctx.Query("role")

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllUser(ctx, roleName)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_USER, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllUserWithPagination(ctx, payload, roleName)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetDetailUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailUser(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_USER, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.UpdateUserRequest
	payload.ID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.UpdateUser(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_USER, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.DeleteUserRequest
	payload.UserID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.DeleteUser(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_DELETE_USER, result)
	ctx.JSON(http.StatusOK, res)
}

// Ticket
func (ah *AdminHandler) CreateTicket(ctx *gin.Context) {
	var payload dto.CreateTicketRequest
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	payload.Name = ctx.PostForm("ticket_name")

	if priceStr := ctx.PostForm("ticket_price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			payload.Price = price
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_PRICE, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	if quotaStr := ctx.PostForm("ticket_quota"); quotaStr != "" {
		if quota, err := strconv.Atoi(quotaStr); err == nil {
			payload.Quota = quota
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_QUOTA, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	fileHeader, err := ctx.FormFile("ticket_image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		defer file.Close()

		payload.ImageUpload.FileHeader = fileHeader
		payload.ImageUpload.FileReader = file
	}

	result, err := ah.adminService.CreateTicket(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllTicket(ctx *gin.Context) {
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllTicket(ctx)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TICKET, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_TICKET, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllTicketWithPagination(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_TICKET,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetDetailTicket(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailTicket(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) UpdateTicket(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.UpdateTicketRequest
	payload.ID = idStr
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	payload.Name = ctx.PostForm("ticket_name")

	if priceStr := ctx.PostForm("ticket_price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			payload.Price = &price
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_PRICE, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	if quotaStr := ctx.PostForm("ticket_quota"); quotaStr != "" {
		if quota, err := strconv.Atoi(quotaStr); err == nil {
			payload.Quota = &quota
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_QUOTA, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	fileHeader, err := ctx.FormFile("ticket_image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		defer file.Close()

		payload.ImageUpload.FileHeader = fileHeader
		payload.ImageUpload.FileReader = file
	}

	result, err := ah.adminService.UpdateTicket(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) DeleteTicket(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.DeleteTicketRequest
	payload.TicketID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.DeleteTicket(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_DELETE_TICKET, result)
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
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllSponsorship(ctx)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPONSORSHIP, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_SPONSORSHIP, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllSponsorshipWithPagination(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_SPONSORSHIP,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

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
		result, err := ah.adminService.GetAllSpeaker(ctx)
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
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	payload.Name = ctx.PostForm("merch_name")
	payload.Description = ctx.PostForm("merch_desc")
	payload.Category = entity.MerchCategory(ctx.PostForm("merch_cat"))

	if priceStr := ctx.PostForm("merch_price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			payload.Price = price
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_PRICE, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	if stockStr := ctx.PostForm("merch_stock"); stockStr != "" {
		if stock, err := strconv.Atoi(stockStr); err == nil {
			payload.Stock = stock
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_STOCK, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
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
		result, err := ah.adminService.GetAllMerch(ctx)
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
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	payload.Name = ctx.PostForm("merch_name")
	payload.Description = ctx.PostForm("merch_desc")
	payload.Category = entity.MerchCategory(ctx.PostForm("merch_cat"))

	if priceStr := ctx.PostForm("merch_price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			payload.Price = &price
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_PRICE, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	if stockStr := ctx.PostForm("merch_stock"); stockStr != "" {
		if stock, err := strconv.Atoi(stockStr); err == nil {
			payload.Stock = &stock
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_STOCK, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
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

// Bundle
func (ah *AdminHandler) CreateBundle(ctx *gin.Context) {
	var payload dto.CreateBundleRequest
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	payload.Name = ctx.PostForm("bundle_name")
	payload.Type = entity.BundleType(ctx.PostForm("bundle_type"))

	if priceStr := ctx.PostForm("bundle_price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			payload.Price = price
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_PRICE, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	if quotaStr := ctx.PostForm("bundle_quota"); quotaStr != "" {
		if quota, err := strconv.Atoi(quotaStr); err == nil {
			payload.Quota = quota
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_QUOTA, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	itemIDs := ctx.Request.PostForm["bundle_items"]
	for _, idStr := range itemIDs {
		if id, err := uuid.Parse(idStr); err == nil {
			payload.BundleItems = append(payload.BundleItems, &id)
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_BUNDLE_ITEM_ID, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	fileHeader, err := ctx.FormFile("bundle_image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		defer file.Close()

		payload.ImageUpload.FileHeader = fileHeader
		payload.ImageUpload.FileReader = file
	}

	result, err := ah.adminService.CreateBundle(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BUNDLE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_BUNDLE, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllBundle(ctx *gin.Context) {
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllBundle(ctx)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_BUNDLE, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_BUNDLE, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllBundleWithPagination(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_BUNDLE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_BUNDLE,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetDetailBundle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailBundle(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_BUNDLE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_BUNDLE, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) UpdateBundle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.UpdateBundleRequest
	payload.ID = idStr
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_MULTIPART_FORM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	payload.Name = ctx.PostForm("bundle_name")
	payload.Type = entity.BundleType(ctx.PostForm("bundle_type"))

	if priceStr := ctx.PostForm("bundle_price"); priceStr != "" {
		if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
			payload.Price = &price
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_PRICE, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	if quotaStr := ctx.PostForm("bundle_quota"); quotaStr != "" {
		if quota, err := strconv.Atoi(quotaStr); err == nil {
			payload.Quota = &quota
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PARSE_QUOTA, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	itemIDs := ctx.Request.PostForm["bundle_items"]
	for _, idStr := range itemIDs {
		if id, err := uuid.Parse(idStr); err == nil {
			payload.BundleItems = append(payload.BundleItems, &id)
		} else {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_BUNDLE_ITEM_ID, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	fileHeader, err := ctx.FormFile("bundle_image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_PHOTO, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
		defer file.Close()

		payload.ImageUpload.FileHeader = fileHeader
		payload.ImageUpload.FileReader = file
	}

	result, err := ah.adminService.UpdateBundle(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_BUNDLE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_BUNDLE, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) DeleteBundle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.DeleteBundleRequest
	payload.BundleID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.DeleteBundle(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_BUNDLE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_DELETE_BUNDLE, result)
	ctx.JSON(http.StatusOK, res)
}

// Student Ambassador
func (ah *AdminHandler) CreateStudentAmbassador(ctx *gin.Context) {
	var payload dto.CreateStudentAmbassadorRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.CreateStudentAmbassador(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_STUDENT_AMBASSADOR, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_STUDENT_AMBASSADOR, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllStudentAmbassador(ctx *gin.Context) {
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllStudentAmbassador(ctx)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_STUDENT_AMBASSADOR, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_STUDENT_AMBASSADOR, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllStudentAmbassadorWithPagination(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_STUDENT_AMBASSADOR, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_STUDENT_AMBASSADOR,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetDetailStudentAmbassador(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailStudentAmbassador(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_STUDENT_AMBASSADOR, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_STUDENT_AMBASSADOR, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) UpdateStudentAmbassador(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.UpdateStudentAmbassadorRequest
	payload.ID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.UpdateStudentAmbassador(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_STUDENT_AMBASSADOR, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_STUDENT_AMBASSADOR, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) DeleteStudentAmbassador(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var payload dto.DeleteStudentAmbassadorRequest
	payload.StudentAmbassadorID = idStr
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.DeleteStudentAmbassador(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_STUDENT_AMBASSADOR, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_DELETE_STUDENT_AMBASSADOR, result)
	ctx.JSON(http.StatusOK, res)
}

// Transaction & Ticket Form
func (ah *AdminHandler) CreateTransactionTicket(ctx *gin.Context) {
	var payload dto.CreateTransactionTicketRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.CreateTransactionTicket(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TRANSACTION_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TRANSACTION_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetAllTransactionTicket(ctx *gin.Context) {
	paginationParam := ctx.DefaultQuery("pagination", "true")
	usePagination := paginationParam != "false"
	transactionStatus := ctx.Query("transaction_status")
	ticketCategory := ctx.Query("ticket_category")

	if !usePagination {
		// Tanpa pagination
		result, err := ah.adminService.GetAllTransactionTicket(ctx, transactionStatus, ticketCategory)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TRANSACTION_TICKET, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_TRANSACTION_TICKET, result)
		ctx.JSON(http.StatusOK, res)
		return
	}

	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ah.adminService.GetAllTransactionTicketWithPagination(ctx, payload, transactionStatus, ticketCategory)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TRANSACTION_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_TRANSACTION_TICKET,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
func (ah *AdminHandler) GetDetailTransactionTicket(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := ah.adminService.GetDetailTransactionTicket(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_TRANSACTION_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_TRANSACTION_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}
