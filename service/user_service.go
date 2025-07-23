package service

import (
	"context"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/helpers"
	"github.com/Amierza/TedXBackend/repository"
)

type (
	IUserService interface {
		// Authentication
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)

		// User
		GetDetailUser(ctx context.Context) (dto.UserResponse, error)

		// Ticket
		GetAllTicket(ctx context.Context) ([]dto.TicketResponse, error)

		// Sponsorship
		GetAllSponsorship(ctx context.Context) ([]dto.SponsorshipResponse, error)

		// Speaker
		GetAllSpeaker(ctx context.Context) ([]dto.SpeakerResponse, error)

		// Merch
		GetAllMerch(ctx context.Context) ([]dto.MerchResponse, error)

		// Bundle
		// GetAllBundle(ctx context.Context) ([]dto.BundleResponse, error)
	}

	UserService struct {
		userRepo   repository.IUserRepository
		jwtService IJWTService
	}
)

func NewUserService(userRepo repository.IUserRepository, jwtService IJWTService) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

// Authentication
func (us *UserService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	if !helpers.IsValidEmail(req.Email) {
		return dto.LoginResponse{}, dto.ErrInvalidEmail
	}

	if len(req.Password) < 8 {
		return dto.LoginResponse{}, dto.ErrInvalidPassword
	}

	user, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.LoginResponse{}, dto.ErrUserNotFound
	}

	if user.Role != "guest" {
		return dto.LoginResponse{}, dto.ErrDeniedAccess
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.LoginResponse{}, dto.ErrPasswordNotMatch
	}

	token, err := us.jwtService.GenerateToken(user.ID.String(), string(user.Role))
	if err != nil {
		return dto.LoginResponse{}, dto.ErrGenerateToken
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}

// User
func (us *UserService) GetDetailUser(ctx context.Context) (dto.UserResponse, error) {
	token := ctx.Value("Authorization").(string)

	userID, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserIDFromToken
	}

	user, _, err := us.userRepo.GetUserByID(ctx, nil, userID)
	if err != nil {
		return dto.UserResponse{}, dto.ErrUserNotFound
	}

	return dto.UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Password:      user.Password,
		Role:          user.Role,
	}, nil
}

// Ticket
func (us *UserService) GetAllTicket(ctx context.Context) ([]dto.TicketResponse, error) {
	tickets, err := us.userRepo.GetAllTicket(ctx, nil)
	if err != nil {
		return nil, dto.ErrGetAllTicketNoPagination
	}

	var datas []dto.TicketResponse
	for _, ticket := range tickets {
		data := dto.TicketResponse{
			ID:    ticket.ID,
			Name:  ticket.Name,
			Price: ticket.Price,
			Quota: ticket.Quota,
			Image: ticket.Image,
		}

		datas = append(datas, data)
	}

	return datas, nil
}

// Sponsorship
func (us *UserService) GetAllSponsorship(ctx context.Context) ([]dto.SponsorshipResponse, error) {
	sponsorships, err := us.userRepo.GetAllSponsorship(ctx, nil)
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

// Speaker
func (us *UserService) GetAllSpeaker(ctx context.Context) ([]dto.SpeakerResponse, error) {
	speakers, err := us.userRepo.GetAllSpeaker(ctx, nil)
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

// Merch
func (us *UserService) GetAllMerch(ctx context.Context) ([]dto.MerchResponse, error) {
	merchs, err := us.userRepo.GetAllMerch(ctx, nil)
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

// Bundle
// func (us *UserService) GetAllBundle(ctx context.Context) ([]dto.BundleResponse, error) {
// 	bundleType := "bundle merch ticket"

// 	bundles, err := us.userRepo.GetAllBundle(ctx, nil, bundleType)
// 	if err != nil {
// 		return nil, dto.ErrGetAllBundleNoPagination
// 	}

// 	var datas []dto.BundleResponse
// 	for _, bundle := range bundles {
// 		data := dto.BundleResponse{
// 			ID:    bundle.ID,
// 			Name:  bundle.Name,
// 			Image: bundle.Image,
// 			Type:  bundle.Type,
// 			Price: bundle.Price,
// 			Quota: bundle.Quota,
// 		}

// 		for _, bi := range bundle.BundleItems {
// 			bundleItem := dto.BundleItemResponse{
// 				ID:        bi.ID,
// 				MerchID:   bi.MerchID,
// 				MerchName: bi.Merch.Name,
// 			}

// 			data.BundleItems = append(data.BundleItems, bundleItem)
// 		}

// 		datas = append(datas, data)
// 	}

// 	return datas, nil
// }
