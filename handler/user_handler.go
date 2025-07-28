package handler

import (
	"net/http"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/service"
	"github.com/Amierza/TedXBackend/utils"
	"github.com/gin-gonic/gin"
)

type (
	IUserHandler interface {
		// Authentication
		Login(ctx *gin.Context)

		// User
		GetDetailUser(ctx *gin.Context)
		UpdateUser(ctx *gin.Context)

		// Ticket
		GetAllTicket(ctx *gin.Context)

		// Sponsorship
		GetAllSponsorship(ctx *gin.Context)

		// Speaker
		GetAllSpeaker(ctx *gin.Context)

		// Merch
		GetAllMerch(ctx *gin.Context)

		// Bundle
		GetAllBundle(ctx *gin.Context)

		// Check Referal Code
		CheckReferalCode(ctx *gin.Context)

		// Snap for trigger midtrans
		CreateTransactionTicket(ctx *gin.Context)

		// Webhook for Midtrans
		UpdateTransactionTicket(ctx *gin.Context)
	}

	UserHandler struct {
		userService service.IUserService
	}
)

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Authentication
func (uh *UserHandler) Login(ctx *gin.Context) {
	var payload dto.LoginRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.Login(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN_USER, result)
	ctx.JSON(http.StatusOK, res)
}

// User
func (uh *UserHandler) GetDetailUser(ctx *gin.Context) {
	result, err := uh.userService.GetDetailUser(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_USER, result)
	ctx.JSON(http.StatusOK, res)
}
func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var payload dto.UpdateUserRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.UpdateUser(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_UPDATE_USER, result)
	ctx.JSON(http.StatusOK, res)
}

// Ticket
func (uh *UserHandler) GetAllTicket(ctx *gin.Context) {
	result, err := uh.userService.GetAllTicket(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}

// Sponsorship
func (uh *UserHandler) GetAllSponsorship(ctx *gin.Context) {
	result, err := uh.userService.GetAllSponsorship(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPONSORSHIP, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_SPONSORSHIP, result)
	ctx.JSON(http.StatusOK, res)
}

// Spekear
func (uh *UserHandler) GetAllSpeaker(ctx *gin.Context) {
	result, err := uh.userService.GetAllSpeaker(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_SPEAKER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_SPEAKER, result)
	ctx.JSON(http.StatusOK, res)
}

// Merch
func (uh *UserHandler) GetAllMerch(ctx *gin.Context) {
	result, err := uh.userService.GetAllMerch(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_MERCH, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_MERCH, result)
	ctx.JSON(http.StatusOK, res)
}

// Bundle
func (uh *UserHandler) GetAllBundle(ctx *gin.Context) {
	result, err := uh.userService.GetAllBundle(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_BUNDLE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_BUNDLE, result)
	ctx.JSON(http.StatusOK, res)
}

// Check Referal Code
func (uh *UserHandler) CheckReferalCode(ctx *gin.Context) {
	var payload dto.CheckReferalCodeRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.CheckReferalCode(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_REFERAL_CODE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_VALID_REFERAL_CODE, result)
	ctx.JSON(http.StatusOK, res)
}

// Snap for trigger midtrans
func (uh *UserHandler) CreateTransactionTicket(ctx *gin.Context) {
	var payload dto.CreateTransactionTicketRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.CreateTransactionTicket(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TRANSACTION_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TRANSACTION_TICKET, result)
	ctx.JSON(http.StatusOK, res)
}

// Webhook for Midtrans
func (uh *UserHandler) UpdateTransactionTicket(ctx *gin.Context) {
	var payload dto.UpdateMidtransTransactionTicketRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := uh.userService.UpdateTransactionTicket(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TRANSACTION_TICKET, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TRANSACTION_TICKET, "")
	ctx.JSON(http.StatusOK, res)
}
