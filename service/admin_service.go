package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

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

		// Speaker
		CreateSpeaker(ctx context.Context, req dto.CreateSpeakerRequest) (dto.SpeakerResponse, error)
		GetAllSpeakerNoPagination(ctx context.Context) ([]dto.SpeakerResponse, error)
		GetAllSpeakerWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.SpeakerPaginationResponse, error)
		GetDetailSpeaker(ctx context.Context, speakerID string) (dto.SpeakerResponse, error)
		UpdateSpeaker(ctx context.Context, req dto.UpdateSpeakerRequest) (dto.SpeakerResponse, error)
		DeleteSpeaker(ctx context.Context, req dto.DeleteSpeakerRequest) (dto.SpeakerResponse, error)
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
	if req.Category == "" || req.Name == "" {
		return dto.SponsorshipResponse{}, dto.ErrEmptyFields
	}

	_, flag, err := as.adminRepo.GetSponsorshipByNameAndCategory(ctx, nil, req.Category, req.Name)
	if err == nil || flag {
		return dto.SponsorshipResponse{}, dto.ErrSponsorshipAlreadyExists
	}

	sponCat := entity.SponsorshipCategory(req.Category)
	if !entity.IsValidSponsorshipCategory(sponCat) {
		return dto.SponsorshipResponse{}, dto.ErrInvalidSponsorshipCategory
	}

	if len(req.Name) < 3 {
		return dto.SponsorshipResponse{}, dto.ErrSponsorshipNameTooShort
	}

	spon := entity.Sponsorship{
		ID:       uuid.New(),
		Category: sponCat,
		Name:     req.Name,
	}

	err = as.adminRepo.CreateSponsorship(ctx, nil, spon)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrCreateSponsorship
	}

	return dto.SponsorshipResponse{
		ID:       spon.ID,
		Category: string(spon.Category),
		Name:     spon.Name,
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
			ID:       sponsorship.ID,
			Category: string(sponsorship.Category),
			Name:     sponsorship.Name,
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
		ID:       sponsorship.ID,
		Category: string(sponsorship.Category),
		Name:     sponsorship.Name,
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

	if req.Category != "" {
		sponCat := entity.SponsorshipCategory(req.Category)
		if !entity.IsValidSponsorshipCategory(sponCat) {
			return dto.SponsorshipResponse{}, dto.ErrInvalidSponsorshipCategory
		}

		sponsorship.Category = sponCat
	}

	err = as.adminRepo.UpdateSponsorship(ctx, nil, sponsorship)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrUpdateSponsorship
	}

	res := dto.SponsorshipResponse{
		ID:       sponsorship.ID,
		Category: string(sponsorship.Category),
		Name:     sponsorship.Name,
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
		ID:       deletedSponsorship.ID,
		Category: string(deletedSponsorship.Category),
		Name:     deletedSponsorship.Name,
	}

	return res, nil
}

// Speaker
func (as *AdminService) CreateSpeaker(ctx context.Context, req dto.CreateSpeakerRequest) (dto.SpeakerResponse, error) {
	if req.FileHeader == nil || req.FileReader == nil || req.Name == "" {
		return dto.SpeakerResponse{}, dto.ErrEmptyFields
	}

	if len(req.Name) < 3 {
		return dto.SpeakerResponse{}, dto.ErrSpeakerNameTooShort
	}

	_, flag, err := as.adminRepo.GetSpeakerByName(ctx, nil, req.Name)
	if err == nil || flag {
		return dto.SpeakerResponse{}, dto.ErrSpeakerAlreadyExists
	}

	if req.FileReader != nil && req.FileHeader != nil {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))

		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.SpeakerResponse{}, dto.ErrInvalidExtensionPhoto
		}

		speakerName := strings.ToLower(req.Name)
		speakerName = strings.ReplaceAll(speakerName, " ", "_")

		fileName := fmt.Sprintf("speaker_%d_%s.%s", time.Now().Unix(), speakerName, ext)

		saveDir := "assets/speaker"
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return dto.SpeakerResponse{}, dto.ErrCreateFile
		}

		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.SpeakerResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, req.FileReader); err != nil {
			return dto.SpeakerResponse{}, dto.ErrSaveFile
		}

		req.Image = fileName
	}

	speaker := entity.Speaker{
		ID:    uuid.New(),
		Name:  req.Name,
		Image: req.Image,
	}

	err = as.adminRepo.CreateSpeaker(ctx, nil, speaker)
	if err != nil {
		return dto.SpeakerResponse{}, dto.ErrCreateSpeaker
	}

	return dto.SpeakerResponse{
		ID:    speaker.ID,
		Name:  speaker.Name,
		Image: speaker.Image,
	}, nil
}
func (as *AdminService) GetAllSpeakerNoPagination(ctx context.Context) ([]dto.SpeakerResponse, error) {
	speakers, err := as.adminRepo.GetAllSpeaker(ctx, nil)
	if err != nil {
		return nil, dto.ErrGetAllSpeakerNoPagination
	}

	var datas []dto.SpeakerResponse
	for _, speaker := range speakers {
		data := dto.SpeakerResponse{
			ID:    speaker.ID,
			Name:  speaker.Name,
			Image: speaker.Image,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllSpeakerWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.SpeakerPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllSpeakerWithPagination(ctx, nil, req)
	if err != nil {
		return dto.SpeakerPaginationResponse{}, dto.ErrGetAllSpeakerWithPagination
	}

	var datas []dto.SpeakerResponse
	for _, speaker := range dataWithPaginate.Speakers {
		data := dto.SpeakerResponse{
			ID:    speaker.ID,
			Name:  speaker.Name,
			Image: speaker.Image,
		}

		datas = append(datas, data)
	}

	return dto.SpeakerPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailSpeaker(ctx context.Context, speakerID string) (dto.SpeakerResponse, error) {
	speaker, _, err := as.adminRepo.GetSpeakerByID(ctx, nil, speakerID)
	if err != nil {
		return dto.SpeakerResponse{}, dto.ErrSpeakerNotFound
	}

	return dto.SpeakerResponse{
		ID:    speaker.ID,
		Name:  speaker.Name,
		Image: speaker.Image,
	}, nil
}
func (as *AdminService) UpdateSpeaker(ctx context.Context, req dto.UpdateSpeakerRequest) (dto.SpeakerResponse, error) {
	speaker, _, err := as.adminRepo.GetSpeakerByID(ctx, nil, req.ID)
	if err != nil {
		return dto.SpeakerResponse{}, dto.ErrSpeakerNotFound
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.SpeakerResponse{}, dto.ErrSpeakerNameTooShort
		}

		speaker.Name = req.Name
	}

	if req.FileReader != nil && req.FileHeader != nil {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))

		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.SpeakerResponse{}, dto.ErrInvalidExtensionPhoto
		}

		if speaker.Image != "" {
			oldPath := filepath.Join("assets/speaker", speaker.Image)
			if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
				return dto.SpeakerResponse{}, dto.ErrDeleteOldImage
			}
		}

		if req.Name == "" {
			req.Name = speaker.Name
		}

		speakerName := strings.ToLower(req.Name)
		speakerName = strings.ReplaceAll(speakerName, " ", "_")

		fileName := fmt.Sprintf("speaker_%d_%s.%s", time.Now().Unix(), speakerName, ext)

		saveDir := "assets/speaker"
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return dto.SpeakerResponse{}, dto.ErrCreateFile
		}

		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.SpeakerResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, req.FileReader); err != nil {
			return dto.SpeakerResponse{}, dto.ErrSaveFile
		}

		speaker.Image = fileName
	}

	err = as.adminRepo.UpdateSpeaker(ctx, nil, speaker)
	if err != nil {
		return dto.SpeakerResponse{}, dto.ErrUpdateSpeaker
	}

	res := dto.SpeakerResponse{
		ID:    speaker.ID,
		Name:  speaker.Name,
		Image: speaker.Image,
	}

	return res, nil
}
func (as *AdminService) DeleteSpeaker(ctx context.Context, req dto.DeleteSpeakerRequest) (dto.SpeakerResponse, error) {
	deletedSpeaker, _, err := as.adminRepo.GetSpeakerByID(ctx, nil, req.SpeakerID)
	if err != nil {
		return dto.SpeakerResponse{}, dto.ErrSpeakerNotFound
	}

	err = as.adminRepo.DeleteSpeakerByID(ctx, nil, req.SpeakerID)
	if err != nil {
		return dto.SpeakerResponse{}, dto.ErrDeleteSpeakerByID
	}

	res := dto.SpeakerResponse{
		ID:    deletedSpeaker.ID,
		Name:  deletedSpeaker.Name,
		Image: deletedSpeaker.Image,
	}

	return res, nil
}
