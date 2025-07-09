package service

import (
	"context"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/helpers"
	"github.com/Amierza/TedXBackend/repository"
)

type (
	IAdminService interface {
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	}

	AdminService struct {
		adminRepo  repository.IAdminRepository
		jwtService IJWTService
	}
)

func NewAdminService(adminRepo repository.IAdminRepository, jwtService IJWTService) *AdminService {
	return &AdminService{
		adminRepo:  adminRepo,
		jwtService: jwtService,
	}
}

func (as *AdminService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	if !helpers.IsValidEmail(req.Email) {
		return dto.LoginResponse{}, dto.ErrInvalidEmail
	}

	if len(req.Password) < 8 {
		return dto.LoginResponse{}, dto.ErrInvalidPassword
	}

	user, flag, err := as.adminRepo.GetUserByEmail(ctx, nil, req.Email)
	if !flag || err != nil {
		return dto.LoginResponse{}, dto.ErrEmailNotFound
	}

	if user.Role != "admin" {
		return dto.LoginResponse{}, dto.ErrDeniedAccess
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.LoginResponse{}, dto.ErrPasswordNotMatch
	}

	token, err := as.jwtService.GenerateToken(user.ID.String(), string(user.Role))
	if err != nil {
		return dto.LoginResponse{}, dto.ErrGenerateToken
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}
