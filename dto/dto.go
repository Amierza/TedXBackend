package dto

import (
	"errors"
	"mime/multipart"
	"time"

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
	MESSAGE_FAILED_PARSE_PRICE          = "failed to parse price"
	MESSAGE_FAILED_PARSE_QUOTA          = "failed to parse quota"
	MESSAGE_FAILED_PARSE_STOCK          = "failed to parse stock"
	// Authentication
	MESSAGE_FAILED_LOGIN_ADMIN = "failed login admin"
	MESSAGE_FAILED_LOGIN_USER  = "failed login user"
	// Middleware
	MESSAGE_FAILED_PROSES_REQUEST             = "failed proses request"
	MESSAGE_FAILED_ACCESS_DENIED              = "failed access denied"
	MESSAGE_FAILED_TOKEN_NOT_FOUND            = "failed token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID            = "failed token not valid"
	MESSAGE_FAILED_TOKEN_DENIED_ACCESS        = "failed token denied access"
	MESSAGE_FAILED_GET_CUSTOM_CLAIMS          = "failed get custom claims"
	MESSAGE_FAILED_GET_ROLE_USER              = "failed get role user"
	MESSAGE_FAILED_INAVLID_ROUTE_FORMAT_TOKEN = "failed invalid route format in token"
	// User
	MESSAGE_FAILED_CREATE_USER     = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER   = "failed get list user"
	MESSAGE_FAILED_GET_DETAIL_USER = "failed get detail user"
	MESSAGE_FAILED_UPDATE_USER     = "failed update user"
	MESSAGE_FAILED_DELETE_USER     = "failed delete user"
	// Ticket
	MESSAGE_FAILED_CREATE_TICKET     = "failed create ticket"
	MESSAGE_FAILED_GET_LIST_TICKET   = "failed get list ticket"
	MESSAGE_FAILED_GET_DETAIL_TICKET = "failed get detail ticket"
	MESSAGE_FAILED_UPDATE_TICKET     = "failed update ticket"
	MESSAGE_FAILED_DELETE_TICKET     = "failed delete ticket"
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
	// Bundle
	MESSAGE_FAILED_INVALID_BUNDLE_ITEM_ID = "failed invalid bundle item id"
	MESSAGE_FAILED_CREATE_BUNDLE          = "failed create bundle"
	MESSAGE_FAILED_GET_LIST_BUNDLE        = "failed get list bundle"
	MESSAGE_FAILED_GET_DETAIL_BUNDLE      = "failed get detail bundle"
	MESSAGE_FAILED_UPDATE_BUNDLE          = "failed update bundle"
	MESSAGE_FAILED_DELETE_BUNDLE          = "failed delete bundle"
	// Transaction & Ticket Form
	MESSAGE_FAILED_CREATE_TICKET_FORM     = "failed create ticket form"
	MESSAGE_FAILED_GET_LIST_TICKET_FORM   = "failed get list ticket form"
	MESSAGE_FAILED_GET_DETAIL_TICKET_FORM = "failed get detail ticket form"
	MESSAGE_FAILED_UPDATE_TICKET_FORM     = "failed update ticket form"
	MESSAGE_FAILED_DELETE_TICKET_FORM     = "failed delete ticket form"

	// ====================================== Success ======================================
	// Authentication
	MESSAGE_SUCCESS_LOGIN_ADMIN = "success login admin"
	MESSAGE_SUCCESS_LOGIN_USER  = "success login user"
	// User
	MESSAGE_SUCCESS_CREATE_USER     = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER   = "success get list user"
	MESSAGE_SUCCESS_GET_DETAIL_USER = "success get detail user"
	MESSAGE_SUCCESS_UPDATE_USER     = "success update user"
	MESSAGE_SUCCESS_DELETE_USER     = "success delete user"
	// Ticket
	MESSAGE_SUCCESS_CREATE_TICKET     = "success create ticket"
	MESSAGE_SUCCESS_GET_LIST_TICKET   = "success get list ticket"
	MESSAGE_SUCCESS_GET_DETAIL_TICKET = "success get detail ticket"
	MESSAGE_SUCCESS_UPDATE_TICKET     = "success update ticket"
	MESSAGE_SUCCESS_DELETE_TICKET     = "success delete ticket"
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
	// Bundle
	MESSAGE_SUCCESS_CREATE_BUNDLE     = "success create bundle"
	MESSAGE_SUCCESS_GET_LIST_BUNDLE   = "success get list bundle"
	MESSAGE_SUCCESS_GET_DETAIL_BUNDLE = "success get detail bundle"
	MESSAGE_SUCCESS_UPDATE_BUNDLE     = "success update bundle"
	MESSAGE_SUCCESS_DELETE_BUNDLE     = "success delete bundle"
	// Transaction & Ticket Form
	MESSAGE_SUCCESS_CREATE_TRANSACTION_TICKET     = "success create transaction ticket"
	MESSAGE_SUCCESS_GET_LIST_TRANSACTION_TICKET   = "success get list transaction ticket"
	MESSAGE_SUCCESS_GET_DETAIL_TRANSACTION_TICKET = "success get detail transaction ticket"
	MESSAGE_SUCCESS_UPDATE_TRANSACTION_TICKET     = "success update transaction ticket"
	MESSAGE_SUCCESS_DELETE_TRANSACTION_TICKET     = "success delete transaction ticket"
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
	ErrGetUserIDFromToken      = errors.New("failed to get user id from token")
	// Parse
	ErrParseUUID = errors.New("failed parse uuid")
	// Input Validation
	ErrEmptyFields                = errors.New("failed there are empty fields")
	ErrInvalidEmail               = errors.New("failed invalid email")
	ErrInvalidPassword            = errors.New("failed invalid password")
	ErrInvalidPhoneNumber         = errors.New("failed invalid phone number")
	ErrInvalidUserRole            = errors.New("failed invalid user role")
	ErrUserNameTooShort           = errors.New("failed user name too short (min 3.)")
	ErrUserFullNameTooShort       = errors.New("failed user name too short (min 5.)")
	ErrPasswordTooShort           = errors.New("failed password too short (min 8.)")
	ErrTicketNameTooShort         = errors.New("failed ticket name too short (min 3.)")
	ErrBundleNameTooShort         = errors.New("failed bundle name too short (min 3.)")
	ErrSponsorshipNameTooShort    = errors.New("failed sponsorship name too short (min 3.)")
	ErrSpeakerNameTooShort        = errors.New("failed speaker name too short (min 3.)")
	ErrSpeakerDescriptionTooShort = errors.New("failed speaker description too short (min 5.)")
	ErrMerchNameTooShort          = errors.New("failed merch name too short (min 3.)")
	ErrMerchDescriptionTooShort   = errors.New("failed merch description too short (min 5.)")
	// Email
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	// Password
	ErrPasswordNotMatch = errors.New("password not match")
	// User
	ErrCreateUser               = errors.New("failed create user")
	ErrGetAllUserNoPagination   = errors.New("failed get all user no pagination")
	ErrGetAllUserWithPagination = errors.New("failed get all user with pagination")
	ErrUserNotFound             = errors.New("failed user not found")
	ErrUpdateUser               = errors.New("failed update user")
	ErrDeleteUserByID           = errors.New("failed delete user by id")
	ErrUserAlreadyExists        = errors.New("failed user already exists")
	// Ticket
	ErrCreateTicket               = errors.New("failed create ticket")
	ErrGetAllTicketNoPagination   = errors.New("failed get all ticket no pagination")
	ErrGetAllTicketWithPagination = errors.New("failed get all ticket with pagination")
	ErrTicketNotFound             = errors.New("failed ticket not found")
	ErrUpdateTicket               = errors.New("failed update ticket")
	ErrDeleteTicketByID           = errors.New("failed delete ticket by id")
	ErrTicketAlreadyExists        = errors.New("failed ticket already exists")
	ErrQuotaOutOfBound            = errors.New("failed quota out of bound")
	ErrTicketSoldOut              = errors.New("failed ticket sold out")
	// Sponsorship
	ErrCreateSponsorship               = errors.New("failed create sponsorship")
	ErrGetAllSponsorship               = errors.New("failed get all sponsorship")
	ErrGetAllSponsorshipWithPagination = errors.New("failed get all sponsorship with pagination")
	ErrSponsorshipNotFound             = errors.New("failed sponsorship not found")
	ErrUpdateSponsorship               = errors.New("failed update sponsorship")
	ErrDeleteSponsorshipByID           = errors.New("failed delete sponsorship by id")
	ErrSponsorshipAlreadyExists        = errors.New("failed sponsorship already exists")
	ErrInvalidSponsorshipCategory      = errors.New("failed invalid sponsorship category")
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
	// Bundle
	ErrCreateBundle                 = errors.New("failed create bundle")
	ErrGetAllBundleNoPagination     = errors.New("failed get all bundle no pagination")
	ErrGetAllBundleWithPagination   = errors.New("failed get all bundle with pagination")
	ErrBundleNotFound               = errors.New("failed bundle not found")
	ErrInvalidBundleItemID          = errors.New("failed bundle item id")
	ErrInvalidBundleType            = errors.New("failed invalid bundle type")
	ErrUpdateBundle                 = errors.New("failed update bundle")
	ErrDeleteBundleByID             = errors.New("failed delete bundle by id")
	ErrBundleAlreadyExists          = errors.New("failed bundle already exists")
	ErrCreateBundleItem             = errors.New("failed create bundle item")
	ErrGetBundleItems               = errors.New("failed get bundle items")
	ErrInvalidTicketIDInBundleMerch = errors.New("bundle type 'bundle merch' cannot contain ticket")
	ErrDeleteBundleItemsByBundleID  = errors.New("failed delete bundle items by bundle id")
	// Transaction & Ticket Form
	ErrEmptyTicketForms              = errors.New("failed empty ticket forms")
	ErrInvalidAudienceType           = errors.New("failed invalid audience type")
	ErrInvalidInstansi               = errors.New("failed invalid instansi")
	ErrInvalidItemType               = errors.New("failed invalid item type")
	ErrCreateTicketForm              = errors.New("failed create ticket form")
	ErrCreateTransaction             = errors.New("failed create transaction")
	ErrMustBeInvitedGuest            = errors.New("failed audience must be invited guest")
	ErrItemTypeMustBeTicket          = errors.New("failed item type must be ticket")
	ErrGetAllTransactionNoPagination = errors.New("failed get all transaction no pagination")
	ErrTicketFormNotFound            = errors.New("failed ticket form not found")
	ErrUpdateTicketForm              = errors.New("failed update ticket form")
	ErrDeleteTicketFormByID          = errors.New("failed delete ticket form by id")
)

// All About Image Request
type ImageUpload struct {
	FileHeader *multipart.FileHeader
	FileReader multipart.File
}

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

// User
type (
	UserResponse struct {
		ID            uuid.UUID   `json:"user_id"`
		Name          string      `json:"user_name"`
		Email         string      `json:"user_email"`
		EmailVerified *time.Time  `json:"email_verified"`
		Password      string      `json:"user_password"`
		Role          entity.Role `json:"user_role"`
	}
	CreateUserRequest struct {
		Name          string     `json:"user_name" form:"user_name"`
		Email         string     `json:"user_email" form:"user_email"`
		EmailVerified *time.Time `json:"email_verified" form:"email_verified"`
		Password      string     `json:"user_password" form:"user_password"`
	}
	UpdateUserRequest struct {
		ID            string     `json:"-"`
		Name          string     `json:"user_name,omitempty" form:"user_name"`
		Email         string     `json:"user_email,omitempty" form:"user_email"`
		EmailVerified *time.Time `json:"email_verified" form:"email_verified"`
		Password      string     `json:"user_password,omitempty" form:"user_password"`
	}
	UserPaginationResponse struct {
		PaginationResponse
		Data []UserResponse `json:"data"`
	}
	UserPaginationRepositoryResponse struct {
		PaginationResponse
		Users []entity.User
	}
	DeleteUserRequest struct {
		UserID string `json:"-"`
	}
)

// Ticket
type (
	TicketResponse struct {
		ID    uuid.UUID `json:"ticket_id"`
		Name  string    `json:"ticket_name"`
		Price float64   `json:"ticket_price"`
		Image string    `json:"ticket_image"`
		Quota int       `json:"ticket_quota"`
	}
	CreateTicketRequest struct {
		Name  string  `json:"ticket_name" form:"ticket_name"`
		Price float64 `json:"ticket_price" form:"ticket_price"`
		Image string  `json:"ticket_image" form:"ticket_image"`
		Quota int     `json:"ticket_quota" form:"ticket_quota"`
		ImageUpload
	}
	UpdateTicketRequest struct {
		ID    string   `json:"-"`
		Name  string   `json:"ticket_name,omitempty" form:"ticket_name"`
		Price *float64 `json:"ticket_price,omitempty" form:"ticket_price"`
		Image string   `json:"ticket_image,omitempty" form:"ticket_image"`
		Quota *int     `json:"ticket_quota,omitempty" form:"ticket_quota"`
		ImageUpload
	}
	TicketPaginationResponse struct {
		PaginationResponse
		Data []TicketResponse `json:"data"`
	}
	TicketPaginationRepositoryResponse struct {
		PaginationResponse
		Tickets []entity.Ticket
	}
	DeleteTicketRequest struct {
		TicketID string `json:"-"`
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
	SponsorshipPaginationResponse struct {
		PaginationResponse
		Data []SponsorshipResponse `json:"data"`
	}
	SponsorshipPaginationRepositoryResponse struct {
		PaginationResponse
		Sponsorships []entity.Sponsorship
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
	CreateMerchRequest struct {
		Name        string               `json:"merch_name" form:"merch_name"`
		Stock       int                  `json:"merch_stock" form:"merch_stock"`
		Price       float64              `json:"merch_price" form:"merch_price"`
		Description string               `json:"merch_desc" form:"merch_desc"`
		Category    entity.MerchCategory `json:"merch_cat" form:"merch_cat"`
		Images      []ImageUpload        `json:"-" form:"-"`
	}
	UpdateMerchRequest struct {
		ID          string               `json:"-"`
		Name        string               `json:"merch_name,omitempty" form:"merch_name"`
		Stock       *int                 `json:"merch_stock,omitempty" form:"merch_stock"`
		Price       *float64             `json:"merch_price,omitempty" form:"merch_price"`
		Description string               `json:"merch_desc,omitempty" form:"merch_desc"`
		Category    entity.MerchCategory `json:"merch_cat,omitempty" form:"merch_cat"`
		Images      []ImageUpload        `json:"merch_images,omitempty" form:"merch_images"`
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

// Bundle
type (
	BundleResponse struct {
		ID          uuid.UUID            `json:"bundle_id"`
		Name        string               `json:"bundle_name"`
		Image       string               `json:"bundle_image"`
		Type        entity.BundleType    `json:"bundle_type"`
		Price       float64              `json:"bundle_price"`
		Quota       int                  `json:"bundle_quota"`
		BundleItems []BundleItemResponse `json:"bundle_items"`
	}
	BundleItemResponse struct {
		ID        uuid.UUID  `json:"bundle_item_id"`
		MerchID   *uuid.UUID `json:"merch_id,omitempty"`
		MerchName string     `json:"merch_name,omitempty"`
	}
	CreateBundleRequest struct {
		Name        string            `json:"bundle_name" form:"bundle_name"`
		Image       string            `json:"bundle_image" form:"bundle_image"`
		Type        entity.BundleType `json:"bundle_type" form:"bundle_type"`
		Price       float64           `json:"bundle_price" form:"bundle_price"`
		Quota       int               `json:"bundle_quota" form:"bundle_quota"`
		BundleItems []*uuid.UUID      `json:"bundle_items"`
		ImageUpload
	}
	UpdateBundleRequest struct {
		ID          string            `json:"-"`
		Name        string            `json:"bundle_name,omitempty" form:"bundle_name"`
		Image       string            `json:"bundle_image,omitempty" form:"bundle_image"`
		Type        entity.BundleType `json:"bundle_type,omitempty" form:"bundle_type"`
		Price       *float64          `json:"bundle_price,omitempty" form:"bundle_price"`
		Quota       *int              `json:"bundle_quota,omitempty" form:"bundle_quota"`
		BundleItems []*uuid.UUID      `json:"bundle_items,omitempty" form:"bundle_items"`
		ImageUpload
	}
	BundlePaginationResponse struct {
		PaginationResponse
		Data []BundleResponse `json:"data"`
	}
	BundlePaginationRepositoryResponse struct {
		PaginationResponse
		Bundles []entity.Bundle
	}
	DeleteBundleRequest struct {
		BundleID string `json:"-"`
	}
)

// Transaction & Ticket Form
type (
	TransactionResponse struct {
		ID                uuid.UUID            `json:"transaction_id"`
		OrderID           string               `json:"order_id"`
		ItemType          entity.ItemType      `json:"item_type"`
		ReferalCode       string               `json:"referal_code"`
		TransactionStatus string               `json:"transaction_status"`
		PaymentType       string               `json:"payment_type"`
		SignatureKey      string               `json:"signature_key"`
		Acquire           string               `json:"acquire"`
		SettlementTime    *time.Time           `json:"settlement_time"`
		GrossAmount       float64              `json:"gross_amount"`
		UserID            *uuid.UUID           `json:"user_id"`
		TicketID          *uuid.UUID           `json:"ticket_id"`
		BundleID          *uuid.UUID           `json:"bundle_id"`
		TicketForms       []TicketFormResponse `json:"ticket_forms"`
	}
	TicketFormResponse struct {
		ID           uuid.UUID           `json:"ticket_form_id"`
		AudienceType entity.AudienceType `json:"audience_type"`
		Instansi     entity.Instansi     `json:"instansi"`
		Email        string              `json:"email"`
		FullName     string              `json:"full_name"`
		PhoneNumber  string              `json:"phone_number"`
		LineID       string              `json:"line_id"`
	}
	TicketFormRequest struct {
		AudienceType entity.AudienceType `json:"audience_type" form:"audience_type"`
		Instansi     entity.Instansi     `json:"instansi" form:"instansi"`
		Email        string              `json:"email" form:"email"`
		FullName     string              `json:"full_name" form:"full_name"`
		PhoneNumber  string              `json:"phone_number" form:"phone_number"`
		LineID       string              `json:"line_id" form:"line_id"`
	}
	CreateTransactionTicketRequest struct {
		ItemType    entity.ItemType     `json:"item_type" form:"item_type"`
		TicketID    *uuid.UUID          `json:"ticket_id" form:"ticket_id"`
		TicketForms []TicketFormRequest `json:"ticket_forms"`
	}
	TransactionTicketPaginationResponse struct {
		PaginationResponse
		Data []TransactionResponse `json:"data"`
	}
	TransactionTicketPaginationRepositoryResponse struct {
		PaginationResponse
		Transactions []entity.Transaction
	}
)
