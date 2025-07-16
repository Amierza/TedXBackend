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
	MESSAGE_FAILED_PARSE_UUID           = "failed parse string to uuid"
	MESSAGE_FAILED_PARSE_MULTIPART_FORM = "failed to parse multipart form"
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
	// Merch
	MESSAGE_FAILED_CREATE_MERCH     = "failed create merch"
	MESSAGE_FAILED_GET_LIST_MERCH   = "failed get list merch"
	MESSAGE_FAILED_GET_DETAIL_MERCH = "failed get detail merch"
	MESSAGE_FAILED_UPDATE_MERCH     = "failed update merch"
	MESSAGE_FAILED_DELETE_MERCH     = "failed delete merch"

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
	// Merch
	MESSAGE_SUCCESS_CREATE_MERCH     = "success create merch"
	MESSAGE_SUCCESS_GET_LIST_MERCH   = "success get list merch"
	MESSAGE_SUCCESS_GET_DETAIL_MERCH = "success get detail merch"
	MESSAGE_SUCCESS_UPDATE_MERCH     = "success update merch"
	MESSAGE_SUCCESS_DELETE_MERCH     = "success delete merch"
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
	ErrSponsorshipNameTooShort    = errors.New("failed sponsorship name too short (min 3.)")
	ErrSpeakerNameTooShort        = errors.New("failed speaker name too short (min 3.)")
	ErrSpeakerDescriptionTooShort = errors.New("failed speaker name too short (min 5.)")
	ErrMerchNameTooShort          = errors.New("failed merch name too short (min 3.)")
	ErrMerchDescriptionTooShort   = errors.New("failed merch name too short (min 5.)")
	// Email
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	// Password
	ErrPasswordNotMatch = errors.New("password not match")
	// Sponsorship
	ErrCreateSponsorship          = errors.New("failed create sponsorship")
	ErrGetAllSponsorship          = errors.New("failed get all sponsorship")
	ErrSponsorshipNotFound        = errors.New("failed sponsorship not found")
	ErrUpdateSponsorship          = errors.New("failed update sponsorship")
	ErrDeleteSponsorshipByID      = errors.New("failed delete sponsorship by id")
	ErrSponsorshipAlreadyExists   = errors.New("failed sponsorship already exists")
	ErrInvalidSponsorshipCategory = errors.New("failed invalid sponsorship category")
	// Speaker
	ErrCreateSpeaker               = errors.New("failed create speaker")
	ErrGetAllSpeakerNoPagination   = errors.New("failed get all speaker no pagination")
	ErrGetAllSpeakerWithPagination = errors.New("failed get all speaker with pagination")
	ErrSpeakerNotFound             = errors.New("failed speaker not found")
	ErrUpdateSpeaker               = errors.New("failed update speaker")
	ErrDeleteSpeakerByID           = errors.New("failed delete speaker by id")
	ErrSpeakerAlreadyExists        = errors.New("failed speaker already exists")
	// Merch
	ErrCreateMerch                = errors.New("failed create merch")
	ErrCreateMerchImage           = errors.New("failed create merch image")
	ErrCreateMerchImageDetail     = errors.New("failed create merch image detail")
	ErrGetAllMerchNoPagination    = errors.New("failed get all merch no pagination")
	ErrGetAllMerchWithPagination  = errors.New("failed get all merch with pagination")
	ErrGetMerchImages             = errors.New("failed get merch images")
	ErrMerchNotFound              = errors.New("failed merch not found")
	ErrMerchImageNotFound         = errors.New("failed merch image not found")
	ErrUpdateMerch                = errors.New("failed update merch")
	ErrDeleteMerchByID            = errors.New("failed delete merch by id")
	ErrDeleteMerchImageByID       = errors.New("failed delete merch image detail by id")
	ErrDeleteMerchImagesByMerchID = errors.New("failed delete merch images by merch id")
	ErrMerchAlreadyExists         = errors.New("failed merch already exists")
	ErrInvalidMerchCategory       = errors.New("failed invalid merch category")
	ErrStockOutOfBound            = errors.New("failed stock out of bound")
	ErrPriceOutOfBound            = errors.New("failed price out of bound")
)

// Authentication
type (
	LoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	LoginResponse struct {
		Token string `json:"token"`
	}
)

// Sponsorship
type (
	SponsorshipResponse struct {
		ID       uuid.UUID `json:"sponsorship_id"`
		Category string    `json:"sponsorship_cat"`
		Name     string    `json:"sponsorship_name"`
		Image    string    `json:"sponsorship_image"`
	}
	CreateSponsorshipRequest struct {
		Category   string                `json:"sponsorship_cat" form:"sponsorship_cat"`
		Name       string                `json:"sponsorship_name" form:"sponsorship_name"`
		Image      string                `json:"sponsorship_image" form:"sponsorship_image"`
		FileHeader *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader multipart.File        `json:"filereader,omitempty"`
	}
	UpdateSponsorshipRequest struct {
		ID         string                `json:"-"`
		Category   string                `json:"sponsorship_cat,omitempty"`
		Name       string                `json:"sponsorship_name,omitempty"`
		Image      string                `json:"sponsorship_image,omitempty"`
		FileHeader *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader multipart.File        `json:"filereader,omitempty"`
	}
	DeleteSponsorshipRequest struct {
		SponsorshipID string `json:"-"`
	}
)

// Speaker
type (
	SpeakerResponse struct {
		ID          uuid.UUID `json:"speaker_id"`
		Name        string    `json:"speaker_name"`
		Image       string    `json:"speaker_image"`
		Description string    `json:"speaker_desc"`
	}
	CreateSpeakerRequest struct {
		Name        string                `json:"speaker_name" form:"speaker_name"`
		Image       string                `json:"speaker_image" form:"speaker_image"`
		Description string                `json:"speaker_desc"`
		FileHeader  *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader  multipart.File        `json:"filereader,omitempty"`
	}
	UpdateSpeakerRequest struct {
		ID          string                `json:"-"`
		Name        string                `json:"speaker_name,omitempty"`
		Image       string                `json:"speaker_image,omitempty"`
		Description string                `json:"speaker_desc,omitempty"`
		FileHeader  *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader  multipart.File        `json:"filereader,omitempty"`
	}
	SpeakerPaginationResponse struct {
		PaginationResponse
		Data []SpeakerResponse `json:"data"`
	}
	SpeakerPaginationRepositoryResponse struct {
		PaginationResponse
		Speakers []entity.Speaker
	}
	DeleteSpeakerRequest struct {
		SpeakerID string `json:"-"`
	}
)

// Merch
type (
	MerchResponse struct {
		ID          uuid.UUID            `json:"merch_id"`
		Name        string               `json:"merch_name"`
		Stock       int                  `json:"merch_stock"`
		Price       float64              `json:"merch_price"`
		Description string               `json:"merch_desc"`
		Category    entity.MerchCategory `json:"merch_cat"`
		Images      []MerchImageResponse `json:"merch_images"`
	}
	MerchImageResponse struct {
		ID   uuid.UUID `json:"merch_image_id"`
		Name string    `json:"merch_image_name"`
	}
	ImageUpload struct {
		FileHeader *multipart.FileHeader
		FileReader multipart.File
	}
	CreateMerchRequest struct {
		Name        string               `json:"merch_name" form:"merch_name"`
		Stock       int                  `json:"merch_stock" form:"merch_stock"`
		Price       float64              `json:"merch_price" form:"merch_price"`
		Description string               `json:"merch_desc" form:"merch_desc"`
		Category    entity.MerchCategory `json:"merch_cat" form:"merch_cat"`
		Images      []ImageUpload        `json:"-" form:"-"`
	}
	ReplaceImageUpload struct {
		TargetImageID uuid.UUID             `form:"target_image_id"`
		FileHeader    *multipart.FileHeader `form:"-"`
		FileReader    multipart.File        `form:"-"`
	}
	UpdateMerchRequest struct {
		ID            string               `json:"-"`
		Name          string               `json:"merch_name,omitempty" form:"merch_name"`
		Stock         *int                 `json:"merch_stock,omitempty" form:"merch_stock"`
		Price         *float64             `json:"merch_price,omitempty" form:"merch_price"`
		Description   string               `json:"merch_desc,omitempty" form:"merch_desc"`
		Category      entity.MerchCategory `json:"merch_cat,omitempty" form:"merch_cat"`
		ReplaceImages []ReplaceImageUpload `json:"-" form:"-"`
	}
	MerchPaginationResponse struct {
		PaginationResponse
		Data []MerchResponse `json:"data"`
	}
	MerchPaginationRepositoryResponse struct {
		PaginationResponse
		Merchs []entity.Merch
	}
	DeleteMerchRequest struct {
		MerchID string `json:"-"`
	}
)
