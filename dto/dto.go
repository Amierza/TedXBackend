package dto

import (
	"errors"
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
	MESSAGE_FAILED_INAVLID_ENPOINTS_TOKEN     = "failed invalid endpoints in token"
	MESSAGE_FAILED_INAVLID_ROUTE_FORMAT_TOKEN = "failed invalid route format in token"
	// User

	// ====================================== Success ======================================
	// Authentication
	MESSAGE_SUCCESS_LOGIN_ADMIN = "success login admin"

	// User

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
	ErrInvalidEmail    = errors.New("failed invalid email")
	ErrInvalidPassword = errors.New("failed invalid password")
	// Email
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	// Password
	ErrPasswordNotMatch = errors.New("password not match")
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
)
