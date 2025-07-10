package dto

import (
	"errors"

	"github.com/Amierza/TedXBackend/entity"
	"github.com/google/uuid"
)

const (
	// ====================================== Failed ======================================
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	// File
	MESSAGE_FAILED_READ_PHOTO = "failed read photo"
	MESSAGE_FAILED_OPEN_PHOTO = "failed open photo"
	// PARSE
	MESSAGE_FAILED_PARSE_UUID = "failed parse string to uuid"
	// Authentication
	MESSAGE_FAILED_LOGIN_ADMIN = "failed login admin"
	// Middleware
	MESSAGE_FAILED_PROSES_REQUEST             = "failed proses request"
	MESSAGE_FAILED_ACCESS_DENIED              = "failed access denied"
	MESSAGE_FAILED_TOKEN_NOT_FOUND            = "failed token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID            = "failed token not valid"
	MESSAGE_FAILED_TOKEN_DENIED_ACCESS        = "failed token denied access"
	MESSAGE_FAILED_GET_CUSTOM_CLAIMS          = "failed get custom claims"
	MESSAGE_FAILED_GET_ROLE_USER              = "failed get role user"
	MESSAGE_FAILED_INAVLID_ROUTE_FORMAT_TOKEN = "failed invalid route format in token"
	// Sponsorship
	MESSAGE_FAILED_CREATE_SPONSORSHIP     = "failed create sponsorship"
	MESSAGE_FAILED_GET_LIST_SPONSORSHIP   = "failed get list sponsorship"
	MESSAGE_FAILED_GET_DETAIL_SPONSORSHIP = "failed get detail sponsorship"
	MESSAGE_FAILED_UPDATE_SPONSORSHIP     = "failed update sponsorship"
	MESSAGE_FAILED_DELETE_SPONSORSHIP     = "failed delete sponsorship"

	// ====================================== Success ======================================
	// Authentication
	MESSAGE_SUCCESS_LOGIN_ADMIN = "success login admin"
	// Sponsorship
	MESSAGE_SUCCESS_CREATE_SPONSORSHIP     = "success create sponsorship"
	MESSAGE_SUCCESS_GET_LIST_SPONSORSHIP   = "success get list sponsorship"
	MESSAGE_SUCCESS_GET_DETAIL_SPONSORSHIP = "success get detail sponsorship"
	MESSAGE_SUCCESS_UPDATE_SPONSORSHIP     = "success update sponsorship"
	MESSAGE_SUCCESS_DELETE_SPONSORSHIP     = "success delete sponsorship"
)

var (
	// Middleware
	ErrDeniedAccess = errors.New("denied access")
	// Token
	ErrGenerateToken           = errors.New("failed to generate token")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrDecryptToken            = errors.New("failed to decrypt token")
	ErrTokenInvalid            = errors.New("token invalid")
	ErrValidateToken           = errors.New("failed to validate token")
	// Parse
	ErrParseUUID = errors.New("failed parse uuid")
	// Input Validation
	ErrEmptyFields                = errors.New("failed there are empty fields")
	ErrInvalidEmail               = errors.New("failed invalid email")
	ErrInvalidPassword            = errors.New("failed invalid password")
	ErrInvalidSponsorshipCategory = errors.New("failed invalid sponsroship category")
	ErrSponsorshipNameTooShort    = errors.New("failed sponsorship name too short (min 3.)")
	// Email
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	// Password
	ErrPasswordNotMatch = errors.New("password not match")
	// Sponsorship
	ErrCreateSponsorship     = errors.New("failed create sponsorship")
	ErrGetAllSponsorship     = errors.New("failed get all sponsorship")
	ErrSponsorshipNotFound   = errors.New("failed sponsorship not found")
	ErrUpdateSponsorship     = errors.New("failed update sponsorship")
	ErrDeleteSponsorshipByID = errors.New("failed delete sponsorship by id")
)

type (
	// Authentication
	LoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	LoginResponse struct {
		Token string `json:"token"`
	}

	// Sponsorship
	SponsorshipResponse struct {
		ID                  uuid.UUID `json:"sponsorship_id"`
		SponsorshipCategory string    `json:"sponsorship_cat"`
		Name                string    `json:"sponsorship_name"`
	}
	SponsorshipPaginationResponse struct {
		PaginationResponse
		Data []SponsorshipResponse `json:"data"`
	}
	SponsorshipPaginationRepositoryResponse struct {
		PaginationResponse
		Sponsorships []entity.Sponsorship
	}
	CreateSponsorshipRequest struct {
		SponsorshipCategory string `json:"sponsorship_cat"`
		Name                string `json:"sponsorship_name"`
	}
	UpdateSponsorshipRequest struct {
		ID                  string `json:"-"`
		SponsorshipCategory string `json:"sponsorship_cat"`
		Name                string `json:"sponsorship_name"`
	}
	DeleteSponsorshipRequest struct {
		SponsorshipID string `json:"-"`
	}
)
