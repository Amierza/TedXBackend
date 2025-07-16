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

		// Merch
		CreateMerch(ctx context.Context, req dto.CreateMerchRequest) (dto.MerchResponse, error)
		GetAllMerchNoPagination(ctx context.Context) ([]dto.MerchResponse, error)
		GetAllMerchWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MerchPaginationResponse, error)
		GetDetailMerch(ctx context.Context, merchID string) (dto.MerchResponse, error)
		UpdateMerch(ctx context.Context, req dto.UpdateMerchRequest) (dto.MerchResponse, error)
		DeleteMerch(ctx context.Context, req dto.DeleteMerchRequest) (dto.MerchResponse, error)
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
	if req.FileHeader == nil || req.FileReader == nil || req.Category == "" || req.Name == "" {
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

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))
	if ext != "jpg" && ext != "jpeg" && ext != "png" {
		return dto.SponsorshipResponse{}, dto.ErrInvalidExtensionPhoto
	}

	sponsorshipName := strings.ToLower(req.Name)
	sponsorshipName = strings.ReplaceAll(sponsorshipName, " ", "_")

	fileName := fmt.Sprintf("sponsorship_%s_%s.%s", time.Now().Format("060102150405"), sponsorshipName, ext)

	saveDir := "assets/sponsorship"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return dto.SponsorshipResponse{}, dto.ErrCreateFile
	}
	savePath := filepath.Join(saveDir, fileName)

	out, err := os.Create(savePath)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrCreateFile
	}
	defer out.Close()

	if _, err := io.Copy(out, req.FileReader); err != nil {
		return dto.SponsorshipResponse{}, dto.ErrSaveFile
	}
	req.Image = fileName

	spon := entity.Sponsorship{
		ID:       uuid.New(),
		Category: sponCat,
		Name:     req.Name,
		Image:    req.Image,
	}

	err = as.adminRepo.CreateSponsorship(ctx, nil, spon)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrCreateSponsorship
	}

	return dto.SponsorshipResponse{
		ID:       spon.ID,
		Category: string(spon.Category),
		Name:     spon.Name,
		Image:    req.Image,
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
			Image:    sponsorship.Image,
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
		Image:    sponsorship.Image,
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

	if req.FileHeader != nil || req.FileReader != nil {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))
		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.SponsorshipResponse{}, dto.ErrInvalidExtensionPhoto
		}

		if sponsorship.Image != "" {
			oldPath := filepath.Join("assets/sponsorship", sponsorship.Image)
			if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
				return dto.SponsorshipResponse{}, dto.ErrDeleteOldImage
			}
		}

		if req.Name == "" {
			req.Name = sponsorship.Name
		}

		sponsorshipName := strings.ToLower(req.Name)
		sponsorshipName = strings.ReplaceAll(sponsorshipName, " ", "_")

		fileName := fmt.Sprintf("sponsorship_%s_%s.%s", time.Now().Format("060102150405"), sponsorshipName, ext)

		saveDir := "assets/sponsorship"
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return dto.SponsorshipResponse{}, dto.ErrCreateFile
		}
		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.SponsorshipResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, req.FileReader); err != nil {
			return dto.SponsorshipResponse{}, dto.ErrSaveFile
		}
		sponsorship.Image = fileName
	}

	err = as.adminRepo.UpdateSponsorship(ctx, nil, sponsorship)
	if err != nil {
		return dto.SponsorshipResponse{}, dto.ErrUpdateSponsorship
	}

	res := dto.SponsorshipResponse{
		ID:       sponsorship.ID,
		Category: string(sponsorship.Category),
		Name:     sponsorship.Name,
		Image:    sponsorship.Image,
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
		Image:    deletedSponsorship.Image,
	}

	return res, nil
}

// Speaker
func (as *AdminService) CreateSpeaker(ctx context.Context, req dto.CreateSpeakerRequest) (dto.SpeakerResponse, error) {
	if req.FileHeader == nil || req.FileReader == nil || req.Name == "" || req.Description == "" {
		return dto.SpeakerResponse{}, dto.ErrEmptyFields
	}

	if len(req.Name) < 3 {
		return dto.SpeakerResponse{}, dto.ErrSpeakerNameTooShort
	}

	if len(req.Description) < 5 {
		return dto.SpeakerResponse{}, dto.ErrSpeakerDescriptionTooShort
	}

	_, flag, err := as.adminRepo.GetSpeakerByName(ctx, nil, req.Name)
	if err == nil || flag {
		return dto.SpeakerResponse{}, dto.ErrSpeakerAlreadyExists
	}

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))
	if ext != "jpg" && ext != "jpeg" && ext != "png" {
		return dto.SpeakerResponse{}, dto.ErrInvalidExtensionPhoto
	}

	speakerName := strings.ToLower(req.Name)
	speakerName = strings.ReplaceAll(speakerName, " ", "_")

	fileName := fmt.Sprintf("speaker_%s_%s.%s", time.Now().Format("060102150405"), speakerName, ext)
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

	speaker := entity.Speaker{
		ID:          uuid.New(),
		Name:        req.Name,
		Image:       req.Image,
		Description: req.Description,
	}

	err = as.adminRepo.CreateSpeaker(ctx, nil, speaker)
	if err != nil {
		return dto.SpeakerResponse{}, dto.ErrCreateSpeaker
	}

	return dto.SpeakerResponse{
		ID:          speaker.ID,
		Name:        speaker.Name,
		Image:       speaker.Image,
		Description: speaker.Description,
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
			ID:          speaker.ID,
			Name:        speaker.Name,
			Image:       speaker.Image,
			Description: speaker.Description,
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
			ID:          speaker.ID,
			Name:        speaker.Name,
			Image:       speaker.Image,
			Description: speaker.Description,
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
		ID:          speaker.ID,
		Name:        speaker.Name,
		Image:       speaker.Image,
		Description: speaker.Description,
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

	if req.Description != "" {
		if len(req.Description) < 3 {
			return dto.SpeakerResponse{}, dto.ErrSpeakerDescriptionTooShort
		}

		speaker.Description = req.Description
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

		fileName := fmt.Sprintf("speaker_%s_%s.%s", time.Now().Format("060102150405"), speakerName, ext)

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
		ID:          speaker.ID,
		Name:        speaker.Name,
		Image:       speaker.Image,
		Description: speaker.Description,
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
		ID:          deletedSpeaker.ID,
		Name:        deletedSpeaker.Name,
		Image:       deletedSpeaker.Image,
		Description: deletedSpeaker.Description,
	}

	return res, nil
}

// Merch
func (as *AdminService) CreateMerch(ctx context.Context, req dto.CreateMerchRequest) (dto.MerchResponse, error) {
	if req.Name == "" || req.Description == "" || req.Category == "" || len(req.Images) == 0 {
		return dto.MerchResponse{}, dto.ErrEmptyFields
	}

	if len(req.Name) < 3 {
		return dto.MerchResponse{}, dto.ErrMerchNameTooShort
	}

	if len(req.Description) < 5 {
		return dto.MerchResponse{}, dto.ErrMerchDescriptionTooShort
	}

	if !entity.IsValidMerchCategory(req.Category) {
		return dto.MerchResponse{}, dto.ErrInvalidMerchCategory
	}

	if req.Stock < 0 {
		return dto.MerchResponse{}, dto.ErrStockOutOfBound
	}

	if req.Price < 0 {
		return dto.MerchResponse{}, dto.ErrPriceOutOfBound
	}

	merchID := uuid.New()
	merch := entity.Merch{
		ID:          merchID,
		Name:        req.Name,
		Stock:       req.Stock,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
	}

	err := as.adminRepo.CreateMerch(ctx, nil, merch)
	if err != nil {
		return dto.MerchResponse{}, dto.ErrCreateMerch
	}

	saveDir := "assets/merch"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return dto.MerchResponse{}, dto.ErrCreateFile
	}

	var imageResponses []dto.MerchImageResponse
	for _, img := range req.Images {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(img.FileHeader.Filename), "."))
		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.MerchResponse{}, dto.ErrInvalidExtensionPhoto
		}

		imageID := uuid.New()
		fileName := fmt.Sprintf("merch_%s.%s", imageID, ext)
		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.MerchResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, img.FileReader); err != nil {
			return dto.MerchResponse{}, dto.ErrSaveFile
		}

		image := entity.MerchImage{
			ID:      imageID,
			MerchID: &merchID,
			Name:    fileName,
		}

		if err := as.adminRepo.CreateMerchImage(ctx, nil, image); err != nil {
			return dto.MerchResponse{}, dto.ErrCreateMerchImage
		}

		imageResponses = append(imageResponses, dto.MerchImageResponse{
			ID:   image.ID,
			Name: image.Name,
		})
	}

	return dto.MerchResponse{
		ID:          merch.ID,
		Name:        merch.Name,
		Stock:       merch.Stock,
		Price:       merch.Price,
		Description: merch.Description,
		Category:    merch.Category,
		Images:      imageResponses,
	}, nil
}
func (as *AdminService) GetAllMerchNoPagination(ctx context.Context) ([]dto.MerchResponse, error) {
	merchs, err := as.adminRepo.GetAllMerch(ctx, nil)
	if err != nil {
		return nil, dto.ErrGetAllMerchNoPagination
	}

	var datas []dto.MerchResponse
	for _, merch := range merchs {
		var merchImages []dto.MerchImageResponse
		for _, img := range merch.MerchImages {
			merchImages = append(merchImages, dto.MerchImageResponse{
				ID:   img.ID,
				Name: img.Name,
			})
		}

		data := dto.MerchResponse{
			ID:          merch.ID,
			Name:        merch.Name,
			Stock:       merch.Stock,
			Price:       merch.Price,
			Description: merch.Description,
			Category:    merch.Category,
			Images:      merchImages,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllMerchWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MerchPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllMerchWithPagination(ctx, nil, req)
	if err != nil {
		return dto.MerchPaginationResponse{}, dto.ErrGetAllMerchWithPagination
	}

	var datas []dto.MerchResponse
	for _, merch := range dataWithPaginate.Merchs {
		var merchImages []dto.MerchImageResponse
		for _, img := range merch.MerchImages {
			merchImages = append(merchImages, dto.MerchImageResponse{
				ID:   img.ID,
				Name: img.Name,
			})
		}

		data := dto.MerchResponse{
			ID:          merch.ID,
			Name:        merch.Name,
			Stock:       merch.Stock,
			Price:       merch.Price,
			Description: merch.Description,
			Category:    merch.Category,
			Images:      merchImages,
		}

		datas = append(datas, data)
	}

	return dto.MerchPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailMerch(ctx context.Context, merchID string) (dto.MerchResponse, error) {
	merch, _, err := as.adminRepo.GetMerchByID(ctx, nil, merchID)
	if err != nil {
		return dto.MerchResponse{}, dto.ErrMerchNotFound
	}

	var merchImages []dto.MerchImageResponse
	for _, img := range merch.MerchImages {
		merchImages = append(merchImages, dto.MerchImageResponse{
			ID:   img.ID,
			Name: img.Name,
		})
	}

	return dto.MerchResponse{
		ID:          merch.ID,
		Name:        merch.Name,
		Stock:       merch.Stock,
		Price:       merch.Price,
		Description: merch.Description,
		Category:    merch.Category,
		Images:      merchImages,
	}, nil
}
func (as *AdminService) UpdateMerch(ctx context.Context, req dto.UpdateMerchRequest) (dto.MerchResponse, error) {
	merch, flag, err := as.adminRepo.GetMerchByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.MerchResponse{}, dto.ErrMerchNotFound
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.MerchResponse{}, dto.ErrMerchNameTooShort
		}

		merch.Name = req.Name
	}

	if req.Description != "" {
		if len(req.Description) < 3 {
			return dto.MerchResponse{}, dto.ErrMerchDescriptionTooShort
		}

		merch.Description = req.Description
	}

	if req.Stock != nil {
		if *req.Stock < 0 {
			return dto.MerchResponse{}, dto.ErrStockOutOfBound
		}

		merch.Stock = *req.Stock
	}

	if req.Price != nil {
		if *req.Price < 0 {
			return dto.MerchResponse{}, dto.ErrPriceOutOfBound
		}

		merch.Price = *req.Price
	}

	if req.Category != "" {
		merchCat := entity.MerchCategory(req.Category)
		if !entity.IsValidMerchCategory(merchCat) {
			return dto.MerchResponse{}, dto.ErrInvalidMerchCategory
		}

		merch.Category = merchCat
	}

	if err = as.adminRepo.UpdateMerch(ctx, nil, merch); err != nil {
		return dto.MerchResponse{}, dto.ErrUpdateMerch
	}

	saveDir := "assets/merch"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return dto.MerchResponse{}, dto.ErrCreateFile
	}

	var imageResponses []dto.MerchImageResponse
	if len(req.ReplaceImages) != 0 {
		for _, img := range req.ReplaceImages {
			oldImage, flag, err := as.adminRepo.GetMerchImageByID(ctx, nil, img.TargetImageID.String())
			if err != nil || !flag {
				return dto.MerchResponse{}, dto.ErrMerchImageNotFound
			}

			oldImagePath := filepath.Join("assets/merch", oldImage.Name)
			if err := os.Remove(oldImagePath); err != nil {
				return dto.MerchResponse{}, dto.ErrDeleteOldImage
			}

			if err := as.adminRepo.DeleteMerchImageByID(ctx, nil, img.TargetImageID.String()); err != nil {
				return dto.MerchResponse{}, dto.ErrDeleteMerchImageByID
			}

			ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(img.FileHeader.Filename), "."))
			if ext != "jpg" && ext != "jpeg" && ext != "png" {
				return dto.MerchResponse{}, dto.ErrInvalidExtensionPhoto
			}

			newImageID := uuid.New()
			newFileName := fmt.Sprintf("merch_%s.%s", newImageID, ext)
			savePath := filepath.Join("assets/merch", newFileName)

			out, err := os.Create(savePath)
			if err != nil {
				return dto.MerchResponse{}, dto.ErrCreateFile
			}
			defer out.Close()
			if _, err := io.Copy(out, img.FileReader); err != nil {
				return dto.MerchResponse{}, dto.ErrSaveFile
			}

			newImage := entity.MerchImage{
				ID:      newImageID,
				MerchID: &merch.ID,
				Name:    newFileName,
			}

			if err := as.adminRepo.CreateMerchImage(ctx, nil, newImage); err != nil {
				return dto.MerchResponse{}, dto.ErrCreateMerchImage
			}
		}
	}

	if len(imageResponses) == 0 {
		images, err := as.adminRepo.GetMerchImagesByMerchID(ctx, nil, merch.ID.String())
		if err != nil {
			return dto.MerchResponse{}, dto.ErrGetMerchImages
		}

		for _, img := range images {
			imageResponses = append(imageResponses, dto.MerchImageResponse{
				ID:   img.ID,
				Name: img.Name,
			})
		}
	}

	return dto.MerchResponse{
		ID:          merch.ID,
		Name:        merch.Name,
		Stock:       merch.Stock,
		Price:       merch.Price,
		Description: merch.Description,
		Category:    merch.Category,
		Images:      imageResponses,
	}, nil
}
func (as *AdminService) DeleteMerch(ctx context.Context, req dto.DeleteMerchRequest) (dto.MerchResponse, error) {
	deletedMerch, flag, err := as.adminRepo.GetMerchByID(ctx, nil, req.MerchID)
	if err != nil || !flag {
		return dto.MerchResponse{}, dto.ErrMerchNotFound
	}

	var merchImages []dto.MerchImageResponse
	for _, img := range deletedMerch.MerchImages {
		merchImages = append(merchImages, dto.MerchImageResponse{
			ID:   img.ID,
			Name: img.Name,
		})
	}

	if err = as.adminRepo.DeleteMerchImagesByMerchID(ctx, nil, req.MerchID); err != nil {
		return dto.MerchResponse{}, dto.ErrDeleteMerchImagesByMerchID
	}

	if err = as.adminRepo.DeleteMerchByID(ctx, nil, req.MerchID); err != nil {
		return dto.MerchResponse{}, dto.ErrDeleteMerchByID
	}

	res := dto.MerchResponse{
		ID:          deletedMerch.ID,
		Name:        deletedMerch.Name,
		Stock:       deletedMerch.Stock,
		Price:       deletedMerch.Price,
		Description: deletedMerch.Description,
		Category:    deletedMerch.Category,
		Images:      merchImages,
	}

	return res, nil
}
