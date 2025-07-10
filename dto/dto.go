package dto

import (
	"errors"
	"mime/multipart"

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
	// Speaker
	MESSAGE_FAILED_CREATE_SPEAKER     = "failed create speaker"
	MESSAGE_FAILED_GET_LIST_SPEAKER   = "failed get list speaker"
	MESSAGE_FAILED_GET_DETAIL_SPEAKER = "failed get detail speaker"
	MESSAGE_FAILED_UPDATE_SPEAKER     = "failed update speaker"
	MESSAGE_FAILED_DELETE_SPEAKER     = "failed delete speaker"

	// ====================================== Success ======================================
	// Authentication
	MESSAGE_SUCCESS_LOGIN_ADMIN = "success login admin"
	// Sponsorship
	MESSAGE_SUCCESS_CREATE_SPONSORSHIP     = "success create sponsorship"
	MESSAGE_SUCCESS_GET_LIST_SPONSORSHIP   = "success get list sponsorship"
	MESSAGE_SUCCESS_GET_DETAIL_SPONSORSHIP = "success get detail sponsorship"
	MESSAGE_SUCCESS_UPDATE_SPONSORSHIP     = "success update sponsorship"
	MESSAGE_SUCCESS_DELETE_SPONSORSHIP     = "success delete sponsorship"
	// Speaker
	MESSAGE_SUCCESS_CREATE_SPEAKER     = "success create speaker"
	MESSAGE_SUCCESS_GET_LIST_SPEAKER   = "success get list speaker"
	MESSAGE_SUCCESS_GET_DETAIL_SPEAKER = "success get detail speaker"
	MESSAGE_SUCCESS_UPDATE_SPEAKER     = "success update speaker"
	MESSAGE_SUCCESS_DELETE_SPEAKER     = "success delete speaker"
)

var (
	// Middleware
	ErrDeniedAccess = errors.New("denied access")
	// File
	ErrInvalidExtensionPhoto = errors.New("only jpg/jpeg/png allowed")
	ErrCreateFile            = errors.New("failed create file")
	ErrSaveFile              = errors.New("failed save file")
	ErrDeleteOldImage        = errors.New("failed delete old image")
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
	ErrSpeakerNameTooShort        = errors.New("failed sponsorship name too short (min 3.)")
	// Email
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	// Password
	ErrPasswordNotMatch = errors.New("password not match")
	// Sponsorship
	ErrCreateSponsorship        = errors.New("failed create sponsorship")
	ErrGetAllSponsorship        = errors.New("failed get all sponsorship")
	ErrSponsorshipNotFound      = errors.New("failed sponsorship not found")
	ErrUpdateSponsorship        = errors.New("failed update sponsorship")
	ErrDeleteSponsorshipByID    = errors.New("failed delete sponsorship by id")
	ErrSponsorshipAlreadyExists = errors.New("failed sponsorship already exists")
	// Speaker
	ErrCreateSpeaker               = errors.New("failed create speaker")
	ErrGetAllSpeakerNoPagination   = errors.New("failed get all speaker no pagination")
	ErrGetAllSpeakerWithPagination = errors.New("failed get all speaker with pagination")
	ErrSpeakerNotFound             = errors.New("failed speaker not found")
	ErrUpdateSpeaker               = errors.New("failed update speaker")
	ErrDeleteSpeakerByID           = errors.New("failed delete speaker by id")
	ErrSpeakerAlreadyExists        = errors.New("failed speaker already exists")
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
		ID       uuid.UUID `json:"sponsorship_id"`
		Category string    `json:"sponsorship_cat"`
		Name     string    `json:"sponsorship_name"`
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
		Category string `json:"sponsorship_cat"`
		Name     string `json:"sponsorship_name"`
	}
	UpdateSponsorshipRequest struct {
		ID       string `json:"-"`
		Category string `json:"sponsorship_cat,omitempty"`
		Name     string `json:"sponsorship_name,omitempty"`
	}
	DeleteSponsorshipRequest struct {
		SponsorshipID string `json:"-"`
	}

	// Speaker
	SpeakerResponse struct {
		ID    uuid.UUID `json:"speaker_id"`
		Name  string    `json:"speaker_name"`
		Image string    `json:"speaker_image"`
	}
	SpeakerPaginationResponse struct {
		PaginationResponse
		Data []SpeakerResponse `json:"data"`
	}
	SpeakerPaginationRepositoryResponse struct {
		PaginationResponse
		Speakers []entity.Speaker
	}
	CreateSpeakerRequest struct {
		Name       string                `json:"speaker_name" form:"speaker_name"`
		Image      string                `json:"speaker_image,omitempty" form:"speaker_image"`
		FileHeader *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader multipart.File        `json:"filereader,omitempty"`
	}
	UpdateSpeakerRequest struct {
		ID         string                `json:"-"`
		Name       string                `json:"speaker_name,omitempty"`
		Image      string                `json:"speaker_image,omitempty"`
		FileHeader *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader multipart.File        `json:"filereader,omitempty"`
	}
	DeleteSpeakerRequest struct {
		SpeakerID string `json:"-"`
	}
)
