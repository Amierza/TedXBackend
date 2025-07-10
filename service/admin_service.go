package service

import (
	"context"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/entity"
	"github.com/Amierza/TedXBackend/helpers"
	"github.com/Amierza/TedXBackend/repository"
	"github.com/google/uuid"
)

type (
	IAdminService interface {
		// Authentication
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)

		// Sponsorship
		CreateSponsorship(ctx context.Context, req dto.CreateSponsorshipRequest) (dto.SponsorshipResponse, error)
		GetAllSponsorship(ctx context.Context) ([]dto.SponsorshipResponse, error)
		GetDetailSponsorship(ctx context.Context, sponsorshipID string) (dto.SponsorshipResponse, error)
		UpdateSponsorship(ctx context.Context, req dto.UpdateSponsorshipRequest) (dto.SponsorshipResponse, error)
		DeleteSponsorship(ctx context.Context, req dto.DeleteSponsorshipRequest) (dto.SponsorshipResponse, error)
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

// Authentication
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

// Sponsorship
func (as *AdminService) CreateSponsorship(ctx context.Context, req dto.CreateSponsorshipRequest) (dto.SponsorshipResponse, error) {
	if req.SponsorshipCategory == "" || req.Name == "" {
		return dto.SponsorshipResponse{}, dto.ErrEmptyFields
	}

	sponCat := entity.SponsorshipCategory(req.SponsorshipCategory)
	if !entity.IsValidSponsorshipCategory(sponCat) {
		return dto.SponsorshipResponse{}, dto.ErrInvalidSponsorshipCategory
	}

	if len(req.Name) < 3 {
		return dto.SponsorshipResponse{}, dto.ErrSponsorshipNameTooShort
	}

	spon := entity.Sponsorship{
		ID:                  uuid.New(),
		SponsorshipCategory: sponCat,
		Name:                req.Name,
	}

	err := as.adminRepo.CreateSponsorship(ctx, nil, spon)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrCreateSponsorship
	}

	return dto.SponsorshipResponse{
		ID:                  spon.ID,
		SponsorshipCategory: string(spon.SponsorshipCategory),
		Name:                spon.Name,
	}, nil
}
func (as *AdminService) GetAllSponsorship(ctx context.Context) ([]dto.SponsorshipResponse, error) {
	sponsorships, err := as.adminRepo.GetAllSponsorship(ctx, nil)
	if err != nil {
		return nil, dto.ErrGetAllSponsorship
	}

	var datas []dto.SponsorshipResponse
	for _, sponsorship := range sponsorships {
		data := dto.SponsorshipResponse{
			ID:                  sponsorship.ID,
			SponsorshipCategory: string(sponsorship.SponsorshipCategory),
			Name:                sponsorship.Name,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetDetailSponsorship(ctx context.Context, sponsorshipID string) (dto.SponsorshipResponse, error) {
	sponsorship, _, err := as.adminRepo.GetSponsorshipByID(ctx, nil, sponsorshipID)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrSponsorshipNotFound
	}

	return dto.SponsorshipResponse{
		ID:                  sponsorship.ID,
		SponsorshipCategory: string(sponsorship.SponsorshipCategory),
		Name:                sponsorship.Name,
	}, nil
}
func (as *AdminService) UpdateSponsorship(ctx context.Context, req dto.UpdateSponsorshipRequest) (dto.SponsorshipResponse, error) {
	sponsorship, _, err := as.adminRepo.GetSponsorshipByID(ctx, nil, req.ID)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrSponsorshipNotFound
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.SponsorshipResponse{}, dto.ErrSponsorshipNameTooShort
		}

		sponsorship.Name = req.Name
	}

	if req.SponsorshipCategory != "" {
		sponCat := entity.SponsorshipCategory(req.SponsorshipCategory)
		if !entity.IsValidSponsorshipCategory(sponCat) {
			return dto.SponsorshipResponse{}, dto.ErrInvalidSponsorshipCategory
		}

		sponsorship.SponsorshipCategory = sponCat
	}

	err = as.adminRepo.UpdateSponsorship(ctx, nil, sponsorship)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrUpdateSponsorship
	}

	res := dto.SponsorshipResponse{
		ID:                  sponsorship.ID,
		SponsorshipCategory: string(sponsorship.SponsorshipCategory),
		Name:                sponsorship.Name,
	}

	return res, nil
}
func (as *AdminService) DeleteSponsorship(ctx context.Context, req dto.DeleteSponsorshipRequest) (dto.SponsorshipResponse, error) {
	deletedSponsorship, _, err := as.adminRepo.GetSponsorshipByID(ctx, nil, req.SponsorshipID)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrSponsorshipNotFound
	}

	err = as.adminRepo.DeleteSponsorshipByID(ctx, nil, req.SponsorshipID)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrDeleteSponsorshipByID
	}

	res := dto.SponsorshipResponse{
		ID:                  deletedSponsorship.ID,
		SponsorshipCategory: string(deletedSponsorship.SponsorshipCategory),
		Name:                deletedSponsorship.Name,
	}

	return res, nil
}
