package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Amierza/TedXBackend/constants"
	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/entity"
	"github.com/Amierza/TedXBackend/helpers"
	"github.com/Amierza/TedXBackend/repository"
	"github.com/Amierza/TedXBackend/utils"
	"github.com/google/uuid"
)

type (
	IAdminService interface {
		// Authentication
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)

		// User
		CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error)
		GetAllUser(ctx context.Context, roleName string) ([]dto.UserResponse, error)
		GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest, roleName string) (dto.UserPaginationResponse, error)
		GetDetailUser(ctx context.Context, userID string) (dto.UserResponse, error)
		UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.UserResponse, error)
		DeleteUser(ctx context.Context, req dto.DeleteUserRequest) (dto.UserResponse, error)

		// Ticket
		CreateTicket(ctx context.Context, req dto.CreateTicketRequest) (dto.TicketResponse, error)
		GetAllTicket(ctx context.Context) ([]dto.TicketResponse, error)
		GetAllTicketWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TicketPaginationResponse, error)
		GetDetailTicket(ctx context.Context, ticketID string) (dto.TicketResponse, error)
		UpdateTicket(ctx context.Context, req dto.UpdateTicketRequest) (dto.TicketResponse, error)
		DeleteTicket(ctx context.Context, req dto.DeleteTicketRequest) (dto.TicketResponse, error)

		// Sponsorship
		CreateSponsorship(ctx context.Context, req dto.CreateSponsorshipRequest) (dto.SponsorshipResponse, error)
		GetAllSponsorship(ctx context.Context) ([]dto.SponsorshipResponse, error)
		GetAllSponsorshipWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.SponsorshipPaginationResponse, error)
		GetDetailSponsorship(ctx context.Context, sponsorshipID string) (dto.SponsorshipResponse, error)
		UpdateSponsorship(ctx context.Context, req dto.UpdateSponsorshipRequest) (dto.SponsorshipResponse, error)
		DeleteSponsorship(ctx context.Context, req dto.DeleteSponsorshipRequest) (dto.SponsorshipResponse, error)

		// Speaker
		CreateSpeaker(ctx context.Context, req dto.CreateSpeakerRequest) (dto.SpeakerResponse, error)
		GetAllSpeaker(ctx context.Context) ([]dto.SpeakerResponse, error)
		GetAllSpeakerWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.SpeakerPaginationResponse, error)
		GetDetailSpeaker(ctx context.Context, speakerID string) (dto.SpeakerResponse, error)
		UpdateSpeaker(ctx context.Context, req dto.UpdateSpeakerRequest) (dto.SpeakerResponse, error)
		DeleteSpeaker(ctx context.Context, req dto.DeleteSpeakerRequest) (dto.SpeakerResponse, error)

		// Merch
		CreateMerch(ctx context.Context, req dto.CreateMerchRequest) (dto.MerchResponse, error)
		GetAllMerch(ctx context.Context) ([]dto.MerchResponse, error)
		GetAllMerchWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MerchPaginationResponse, error)
		GetDetailMerch(ctx context.Context, merchID string) (dto.MerchResponse, error)
		UpdateMerch(ctx context.Context, req dto.UpdateMerchRequest) (dto.MerchResponse, error)
		DeleteMerch(ctx context.Context, req dto.DeleteMerchRequest) (dto.MerchResponse, error)

		// Bundle
		CreateBundle(ctx context.Context, req dto.CreateBundleRequest) (dto.BundleResponse, error)
		GetAllBundle(ctx context.Context) ([]dto.BundleResponse, error)
		GetAllBundleWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BundlePaginationResponse, error)
		GetDetailBundle(ctx context.Context, bundleID string) (dto.BundleResponse, error)
		UpdateBundle(ctx context.Context, req dto.UpdateBundleRequest) (dto.BundleResponse, error)
		DeleteBundle(ctx context.Context, req dto.DeleteBundleRequest) (dto.BundleResponse, error)

		// Student Ambassador
		CreateStudentAmbassador(ctx context.Context, req dto.CreateStudentAmbassadorRequest) (dto.StudentAmbassadorResponse, error)
		GetAllStudentAmbassador(ctx context.Context) ([]dto.StudentAmbassadorResponse, error)
		GetAllStudentAmbassadorWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.StudentAmbassadorPaginationResponse, error)
		GetDetailStudentAmbassador(ctx context.Context, studentAmbassadorID string) (dto.StudentAmbassadorResponse, error)
		UpdateStudentAmbassador(ctx context.Context, req dto.UpdateStudentAmbassadorRequest) (dto.StudentAmbassadorResponse, error)
		DeleteStudentAmbassador(ctx context.Context, req dto.DeleteStudentAmbassadorRequest) (dto.StudentAmbassadorResponse, error)

		// Ticket Form
		CreateTransactionTicket(ctx context.Context, req dto.CreateTransactionTicketRequest) (dto.TransactionResponse, error)
		GetAllTransactionTicket(ctx context.Context, transactionStatus, ticketCategory string) ([]dto.TransactionResponse, error)
		GetAllTransactionTicketWithPagination(ctx context.Context, req dto.PaginationRequest, transactionStatus, ticketCategory string) (dto.TransactionTicketPaginationResponse, error)
		GetDetailTransactionTicket(ctx context.Context, transactionTicketID string) (dto.TransactionResponse, error)

		// Check-in
		GetDetailTicketCheckIn(ctx context.Context, ticketFormIDStr string) (dto.TicketCheckInResponse, error)
		CheckIn(ctx context.Context, ticketFormIDStr string) error
		GetAllTicketCheckIn(ctx context.Context, filter dto.CheckInFilterQuery) ([]dto.TicketCheckInResponse, error)
		GetAllTicketCheckInWithPagination(ctx context.Context, req dto.PaginationRequest, filter dto.CheckInFilterQuery) (dto.TicketFormPaginationResponse, error)
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

// User
func (as *AdminService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
	if req.Email == "" || req.Name == "" || req.Password == "" {
		return dto.UserResponse{}, dto.ErrEmptyFields
	}

	_, flag, err := as.adminRepo.GetUserByEmail(ctx, nil, req.Email)
	if err == nil || flag {
		return dto.UserResponse{}, dto.ErrUserAlreadyExists
	}

	if !helpers.IsValidEmail(req.Email) {
		return dto.UserResponse{}, dto.ErrInvalidEmail
	}

	if len(req.Name) < 3 {
		return dto.UserResponse{}, dto.ErrUserNameTooShort
	}

	if len(req.Password) < 8 {
		return dto.UserResponse{}, dto.ErrPasswordTooShort
	}

	role := "admin"

	user := entity.User{
		ID:            uuid.New(),
		Name:          req.Name,
		Email:         req.Email,
		EmailVerified: req.EmailVerified,
		Password:      req.Password,
		Role:          entity.Role(role),
	}

	err = as.adminRepo.CreateUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
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
func (as *AdminService) GetAllUser(ctx context.Context, roleName string) ([]dto.UserResponse, error) {
	users, err := as.adminRepo.GetAllUser(ctx, nil, roleName)
	if err != nil {
		return nil, dto.ErrGetAllUserNoPagination
	}

	var datas []dto.UserResponse
	for _, user := range users {
		data := dto.UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			EmailVerified: user.EmailVerified,
			Password:      user.Password,
			Role:          user.Role,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest, roleName string) (dto.UserPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllUserWithPagination(ctx, nil, req, roleName)
	if err != nil {
		return dto.UserPaginationResponse{}, dto.ErrGetAllUserWithPagination
	}

	var datas []dto.UserResponse
	for _, user := range dataWithPaginate.Users {
		data := dto.UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			EmailVerified: user.EmailVerified,
			Password:      user.Password,
			Role:          user.Role,
		}

		datas = append(datas, data)
	}

	return dto.UserPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailUser(ctx context.Context, userID string) (dto.UserResponse, error) {
	user, _, err := as.adminRepo.GetUserByID(ctx, nil, userID)
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
func (as *AdminService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	user, flag, err := as.adminRepo.GetUserByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.UserResponse{}, dto.ErrUserNotFound
	}

	if req.Email != "" {
		_, flag, err := as.adminRepo.GetUserByEmail(ctx, nil, req.Email)
		if err == nil || flag {
			return dto.UserResponse{}, dto.ErrEmailAlreadyExists
		}

		if !helpers.IsValidEmail(req.Email) {
			return dto.UserResponse{}, dto.ErrInvalidEmail
		}

		user.Email = req.Email
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.UserResponse{}, dto.ErrUserNameTooShort
		}

		user.Name = req.Name
	}

	if req.Password != "" {
		if len(req.Password) < 8 {
			return dto.UserResponse{}, dto.ErrPasswordTooShort
		}

		hashP, err := helpers.HashPassword(req.Password)
		if err != nil {
			return dto.UserResponse{}, dto.ErrHashPassword
		}

		user.Password = hashP
	}

	err = as.adminRepo.UpdateUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrUpdateUser
	}

	res := dto.UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Password:      user.Password,
		Role:          user.Role,
	}

	return res, nil
}
func (as *AdminService) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) (dto.UserResponse, error) {
	deletedUser, _, err := as.adminRepo.GetUserByID(ctx, nil, req.UserID)
	if err != nil {
		return dto.UserResponse{}, dto.ErrUserNotFound
	}

	err = as.adminRepo.DeleteUserByID(ctx, nil, req.UserID)
	if err != nil {
		return dto.UserResponse{}, dto.ErrDeleteUserByID
	}

	res := dto.UserResponse{
		ID:            deletedUser.ID,
		Name:          deletedUser.Name,
		Email:         deletedUser.Email,
		EmailVerified: deletedUser.EmailVerified,
		Password:      deletedUser.Password,
		Role:          deletedUser.Role,
	}

	return res, nil
}

// Ticket
func (as *AdminService) CreateTicket(ctx context.Context, req dto.CreateTicketRequest) (dto.TicketResponse, error) {
	if req.Name == "" || req.FileHeader == nil || req.FileReader == nil || req.Type == "" || req.EventDate == "" {
		return dto.TicketResponse{}, dto.ErrEmptyFields
	}

	_, flag, err := as.adminRepo.GetTicketByName(ctx, nil, req.Name)
	if err == nil || flag {
		return dto.TicketResponse{}, dto.ErrTicketAlreadyExists
	}

	if len(req.Name) < 3 {
		return dto.TicketResponse{}, dto.ErrTicketNameTooShort
	}

	if !entity.IsValidTicketType(req.Type) {
		return dto.TicketResponse{}, dto.ErrInvalidTicketType
	}

	if req.Price < 0 {
		return dto.TicketResponse{}, dto.ErrPriceOutOfBound
	}

	if req.Quota < 0 {
		return dto.TicketResponse{}, dto.ErrQuotaOutOfBound
	}

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))
	if ext != "jpg" && ext != "jpeg" && ext != "png" {
		return dto.TicketResponse{}, dto.ErrInvalidExtensionPhoto
	}

	ticketName := strings.ToLower(req.Name)
	ticketName = strings.ReplaceAll(ticketName, " ", "_")

	fileName := fmt.Sprintf("ticket_%s_%s.%s", time.Now().Format("060102150405"), ticketName, ext)

	saveDir := "assets/ticket"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return dto.TicketResponse{}, dto.ErrCreateFile
	}
	savePath := filepath.Join(saveDir, fileName)

	out, err := os.Create(savePath)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrCreateFile
	}
	defer out.Close()

	if _, err := io.Copy(out, req.FileReader); err != nil {
		return dto.TicketResponse{}, dto.ErrSaveFile
	}
	req.Image = fileName

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.FixedZone("UTC+7", 7*60*60)
	}

	eventDate, err := time.ParseInLocation("2006-01-02", req.EventDate, loc)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrParseTime
	}

	ticket := entity.Ticket{
		ID:          uuid.New(),
		Name:        req.Name,
		Type:        req.Type,
		Price:       req.Price,
		Quota:       req.Quota,
		Image:       req.Image,
		Description: req.Description,
		EventDate:   eventDate,
	}

	err = as.adminRepo.CreateTicket(ctx, nil, ticket)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrCreateTicket
	}

	return dto.TicketResponse{
		ID:          ticket.ID.String(),
		Name:        ticket.Name,
		Type:        ticket.Type,
		Price:       ticket.Price,
		Quota:       ticket.Quota,
		Image:       ticket.Image,
		Description: ticket.Description,
		EventDate:   ticket.EventDate.Format("2006-01-02"),
	}, nil
}
func (as *AdminService) GetAllTicket(ctx context.Context) ([]dto.TicketResponse, error) {
	tickets, err := as.adminRepo.GetAllTicket(ctx, nil)
	if err != nil {
		return nil, dto.ErrGetAllTicketNoPagination
	}

	var datas []dto.TicketResponse
	for _, ticket := range tickets {
		isAvailable := ticket.Quota > 0 && time.Now().Before(ticket.EventDate)

		data := dto.TicketResponse{
			ID:          ticket.ID.String(),
			Name:        ticket.Name,
			Type:        ticket.Type,
			Price:       ticket.Price,
			Quota:       ticket.Quota,
			Image:       ticket.Image,
			Description: ticket.Description,
			EventDate:   ticket.EventDate.Format("2006-01-02"),
			IsAvailable: &isAvailable,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllTicketWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TicketPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllTicketWithPagination(ctx, nil, req)
	if err != nil {
		return dto.TicketPaginationResponse{}, dto.ErrGetAllTicketWithPagination
	}

	var datas []dto.TicketResponse
	for _, ticket := range dataWithPaginate.Tickets {
		isAvailable := ticket.Quota > 0 && time.Now().Before(ticket.EventDate)

		data := dto.TicketResponse{
			ID:          ticket.ID.String(),
			Name:        ticket.Name,
			Type:        ticket.Type,
			Price:       ticket.Price,
			Quota:       ticket.Quota,
			Image:       ticket.Image,
			Description: ticket.Description,
			EventDate:   ticket.EventDate.Format("2006-01-02"),
			IsAvailable: &isAvailable,
		}

		datas = append(datas, data)
	}

	return dto.TicketPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailTicket(ctx context.Context, ticketID string) (dto.TicketResponse, error) {
	ticket, _, err := as.adminRepo.GetTicketByID(ctx, nil, ticketID)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrTicketNotFound
	}

	return dto.TicketResponse{
		ID:          ticket.ID.String(),
		Name:        ticket.Name,
		Type:        ticket.Type,
		Price:       ticket.Price,
		Quota:       ticket.Quota,
		Image:       ticket.Image,
		Description: ticket.Description,
		EventDate:   ticket.EventDate.Format("2006-01-02"),
	}, nil
}
func (as *AdminService) UpdateTicket(ctx context.Context, req dto.UpdateTicketRequest) (dto.TicketResponse, error) {
	ticket, flag, err := as.adminRepo.GetTicketByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.TicketResponse{}, dto.ErrTicketNotFound
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.TicketResponse{}, dto.ErrTicketNameTooShort
		}

		ticket.Name = req.Name
	}

	if req.Type != "" {
		if req.Type == ticket.Type {
			return dto.TicketResponse{}, dto.ErrSameTicketType
		}

		if !entity.IsValidTicketType(req.Type) {
			return dto.TicketResponse{}, dto.ErrInvalidTicketType
		}

		ticket.Type = req.Type
	}

	if req.Price != nil {
		if *req.Price < 0 {
			return dto.TicketResponse{}, dto.ErrPriceOutOfBound
		}

		ticket.Price = *req.Price
	}

	if req.Quota != nil {
		if *req.Quota < 0 {
			return dto.TicketResponse{}, dto.ErrQuotaOutOfBound
		}

		ticket.Quota = *req.Quota
	}

	if req.FileHeader != nil || req.FileReader != nil {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))
		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.TicketResponse{}, dto.ErrInvalidExtensionPhoto
		}

		if ticket.Image != "" {
			oldPath := filepath.Join("assets/ticket", ticket.Image)
			if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
				return dto.TicketResponse{}, dto.ErrDeleteOldImage
			}
		}

		ticketName := strings.ToLower(ticket.Name)
		ticketName = strings.ReplaceAll(ticketName, " ", "_")

		fileName := fmt.Sprintf("ticket_%s_%s.%s", time.Now().Format("060102150405"), ticketName, ext)

		saveDir := "assets/ticket"
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return dto.TicketResponse{}, dto.ErrCreateFile
		}
		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.TicketResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, req.FileReader); err != nil {
			return dto.TicketResponse{}, dto.ErrSaveFile
		}
		ticket.Image = fileName
	}

	if req.Description != "" {
		ticket.Description = req.Description
	}

	if req.EventDate != "" {
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			loc = time.FixedZone("UTC+7", 7*60*60)
		}
		eventDate, err := time.ParseInLocation("2006-01-02", req.EventDate, loc)
		if err != nil {
			return dto.TicketResponse{}, dto.ErrParseTime
		}

		ticket.EventDate = eventDate
	}

	err = as.adminRepo.UpdateTicket(ctx, nil, ticket)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrCreateTicket
	}

	return dto.TicketResponse{
		ID:          ticket.ID.String(),
		Name:        ticket.Name,
		Type:        ticket.Type,
		Price:       ticket.Price,
		Quota:       ticket.Quota,
		Image:       ticket.Image,
		Description: ticket.Description,
		EventDate:   ticket.EventDate.Format("2006-01-02"),
	}, nil
}
func (as *AdminService) DeleteTicket(ctx context.Context, req dto.DeleteTicketRequest) (dto.TicketResponse, error) {
	deletedTicket, _, err := as.adminRepo.GetTicketByID(ctx, nil, req.TicketID)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrTicketNotFound
	}

	err = as.adminRepo.DeleteTicketByID(ctx, nil, req.TicketID)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrDeleteTicketByID
	}

	res := dto.TicketResponse{
		ID:        deletedTicket.ID.String(),
		Name:      deletedTicket.Name,
		Type:      deletedTicket.Type,
		Price:     deletedTicket.Price,
		Quota:     deletedTicket.Quota,
		Image:     deletedTicket.Image,
		EventDate: deletedTicket.EventDate.Format("2006-01-02"),
	}

	return res, nil
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
func (as *AdminService) GetAllSponsorshipWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.SponsorshipPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllSponsorshipWithPagination(ctx, nil, req)
	if err != nil {
		return dto.SponsorshipPaginationResponse{}, dto.ErrGetAllSponsorshipWithPagination
	}

	var datas []dto.SponsorshipResponse
	for _, sponsorship := range dataWithPaginate.Sponsorships {
		data := dto.SponsorshipResponse{
			ID:       sponsorship.ID,
			Category: string(sponsorship.Category),
			Name:     sponsorship.Name,
			Image:    sponsorship.Image,
		}

		datas = append(datas, data)
	}

	return dto.SponsorshipPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
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

		sponsorshipName := strings.ToLower(sponsorship.Name)
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
func (as *AdminService) GetAllSpeaker(ctx context.Context) ([]dto.SpeakerResponse, error) {
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

	saveDir := "assets/merch"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return dto.MerchResponse{}, dto.ErrCreateFile
	}

	var merchImages []entity.MerchImage
	var imageResponses []dto.MerchImageResponse

	for _, img := range req.Images {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(img.FileHeader.Filename), "."))
		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.MerchResponse{}, dto.ErrInvalidExtensionPhoto
		}

		imageID := uuid.New()
		fileName := fmt.Sprintf("merch_%s.%s", imageID.String(), ext)
		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.MerchResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, img.FileReader); err != nil {
			return dto.MerchResponse{}, dto.ErrSaveFile
		}

		merchImages = append(merchImages, entity.MerchImage{
			ID:      imageID,
			Name:    fileName,
			MerchID: nil,
		})

		imageResponses = append(imageResponses, dto.MerchImageResponse{
			ID:   imageID,
			Name: fileName,
		})
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

	err := as.adminRepo.RunInTransaction(ctx, func(txRepo repository.IAdminRepository) error {
		if err := txRepo.CreateMerch(ctx, nil, merch); err != nil {
			return dto.ErrCreateMerch
		}

		for _, img := range merchImages {
			img.MerchID = &merchID
			if err := txRepo.CreateMerchImage(ctx, nil, img); err != nil {
				return dto.ErrCreateMerchImage
			}
		}

		return nil
	})
	if err != nil {
		return dto.MerchResponse{}, err
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
func (as *AdminService) GetAllMerch(ctx context.Context) ([]dto.MerchResponse, error) {
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
	merch, found, err := as.adminRepo.GetMerchByID(ctx, nil, req.ID)
	if err != nil || !found {
		return dto.MerchResponse{}, dto.ErrMerchNotFound
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.MerchResponse{}, dto.ErrMerchNameTooShort
		}
		merch.Name = req.Name
	}

	if req.Description != "" {
		if len(req.Description) < 5 {
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
		category := entity.MerchCategory(req.Category)
		if !entity.IsValidMerchCategory(category) {
			return dto.MerchResponse{}, dto.ErrInvalidMerchCategory
		}
		merch.Category = category
	}

	err = as.adminRepo.RunInTransaction(ctx, func(txRepo repository.IAdminRepository) error {
		if err := txRepo.UpdateMerch(ctx, nil, merch); err != nil {
			return dto.ErrUpdateMerch
		}

		if len(req.Images) > 0 {
			oldImages, err := txRepo.GetMerchImagesByMerchID(ctx, nil, merch.ID.String())
			if err != nil {
				return dto.ErrGetMerchImages
			}

			for _, img := range oldImages {
				path := filepath.Join("assets/merch", img.Name)
				if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
					return dto.ErrDeleteOldImage
				}
				if err := txRepo.DeleteMerchImageByID(ctx, nil, img.ID.String()); err != nil {
					return dto.ErrDeleteMerchImageByID
				}
			}

			for _, img := range req.Images {
				ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(img.FileHeader.Filename), "."))
				if ext != "jpg" && ext != "jpeg" && ext != "png" {
					return dto.ErrInvalidExtensionPhoto
				}

				newImageID := uuid.New()
				newFileName := fmt.Sprintf("merch_%s.%s", newImageID, ext)

				saveDir := "assets/merch"
				if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
					return dto.ErrCreateFile
				}
				savePath := filepath.Join(saveDir, newFileName)

				out, err := os.Create(savePath)
				if err != nil {
					return dto.ErrCreateFile
				}
				defer out.Close()

				if _, err := io.Copy(out, img.FileReader); err != nil {
					return dto.ErrSaveFile
				}

				newImage := entity.MerchImage{
					ID:      newImageID,
					MerchID: &merch.ID,
					Name:    newFileName,
				}
				if err := txRepo.CreateMerchImage(ctx, nil, newImage); err != nil {
					return dto.ErrCreateMerchImage
				}
			}
		}

		return nil
	})
	if err != nil {
		return dto.MerchResponse{}, err
	}

	images, err := as.adminRepo.GetMerchImagesByMerchID(ctx, nil, merch.ID.String())
	if err != nil {
		return dto.MerchResponse{}, dto.ErrGetMerchImages
	}

	var imageResponses []dto.MerchImageResponse
	for _, img := range images {
		imageResponses = append(imageResponses, dto.MerchImageResponse{
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

// Bundle
func (as *AdminService) CreateBundle(ctx context.Context, req dto.CreateBundleRequest) (dto.BundleResponse, error) {
	if req.Name == "" || req.FileHeader == nil || req.FileReader == nil || len(req.BundleItems) == 0 || req.Type == "" || req.EventDate == "" {
		return dto.BundleResponse{}, dto.ErrEmptyFields
	}

	if len(req.Name) < 3 {
		return dto.BundleResponse{}, dto.ErrBundleNameTooShort
	}

	if !entity.IsValidBundleType(req.Type) {
		return dto.BundleResponse{}, dto.ErrInvalidBundleType
	}

	if req.Price < 0 {
		return dto.BundleResponse{}, dto.ErrPriceOutOfBound
	}

	if req.Quota < 0 {
		return dto.BundleResponse{}, dto.ErrQuotaOutOfBound
	}

	var fileName string
	if req.FileHeader != nil && req.FileReader != nil {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))
		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.BundleResponse{}, dto.ErrInvalidExtensionPhoto
		}

		bundleName := strings.ToLower(req.Name)
		bundleName = strings.ReplaceAll(bundleName, " ", "_")

		fileName = fmt.Sprintf("bundle_%s_%s.%s", time.Now().Format("060102150405"), bundleName, ext)

		saveDir := "assets/bundle"
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return dto.BundleResponse{}, dto.ErrCreateFile
		}
		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.BundleResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, req.FileReader); err != nil {
			return dto.BundleResponse{}, dto.ErrSaveFile
		}
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.FixedZone("UTC+7", 7*60*60)
	}
	eventDate, err := time.ParseInLocation("2006-01-02", req.EventDate, loc)
	if err != nil {
		return dto.BundleResponse{}, dto.ErrParseTime
	}

	bundleID := uuid.New()
	bundle := entity.Bundle{
		ID:          bundleID,
		Name:        req.Name,
		Type:        req.Type,
		Price:       req.Price,
		Quota:       req.Quota,
		Image:       fileName,
		Description: req.Description,
		EventDate:   eventDate,
	}

	var bundleItems []entity.BundleItem
	for _, merchID := range req.BundleItems {
		bundleItems = append(bundleItems, entity.BundleItem{
			ID:       uuid.New(),
			BundleID: &bundleID,
			MerchID:  merchID,
		})
	}

	err = as.adminRepo.RunInTransaction(ctx, func(txRepo repository.IAdminRepository) error {
		if err := txRepo.CreateBundle(ctx, nil, bundle); err != nil {
			return err
		}
		for _, item := range bundleItems {
			_, _, err := txRepo.GetMerchByID(ctx, nil, item.MerchID.String())
			if err != nil {
				return dto.ErrMerchNotFound
			}

			if err := txRepo.CreateBundleItem(ctx, nil, item); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return dto.BundleResponse{}, err
	}

	var itemsResp []dto.BundleItemResponse
	for _, item := range bundleItems {
		merch, found, err := as.adminRepo.GetMerchByID(ctx, nil, item.MerchID.String())
		if err == nil && found {
			itemsResp = append(itemsResp, dto.BundleItemResponse{
				ID:        item.ID,
				MerchID:   item.MerchID,
				MerchName: merch.Name,
			})
		}
	}

	return dto.BundleResponse{
		ID:          bundle.ID,
		Name:        bundle.Name,
		Image:       bundle.Image,
		Type:        bundle.Type,
		Price:       bundle.Price,
		Quota:       bundle.Quota,
		Description: bundle.Description,
		EventDate:   bundle.EventDate.Format("2006-01-02"),
		BundleItems: itemsResp,
	}, nil
}
func (as *AdminService) GetAllBundle(ctx context.Context) ([]dto.BundleResponse, error) {
	bundles, err := as.adminRepo.GetAllBundle(ctx, nil)
	if err != nil {
		return nil, dto.ErrGetAllBundleNoPagination
	}

	var datas []dto.BundleResponse
	for _, bundle := range bundles {
		isAvailable := bundle.Quota > 0 && time.Now().Before(bundle.EventDate)

		data := dto.BundleResponse{
			ID:          bundle.ID,
			Name:        bundle.Name,
			Image:       bundle.Image,
			Type:        bundle.Type,
			Price:       bundle.Price,
			Quota:       bundle.Quota,
			Description: bundle.Description,
			EventDate:   bundle.EventDate.Format("2006-01-02"),
			IsAvailable: &isAvailable,
		}

		for _, bi := range bundle.BundleItems {
			bundleItem := dto.BundleItemResponse{
				ID:        bi.ID,
				MerchID:   bi.MerchID,
				MerchName: bi.Merch.Name,
			}

			data.BundleItems = append(data.BundleItems, bundleItem)
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllBundleWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.BundlePaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllBundleWithPagination(ctx, nil, req)
	if err != nil {
		return dto.BundlePaginationResponse{}, dto.ErrGetAllBundleWithPagination
	}

	var datas []dto.BundleResponse
	for _, bundle := range dataWithPaginate.Bundles {
		isAvailable := bundle.Quota > 0 && time.Now().Before(bundle.EventDate)

		data := dto.BundleResponse{
			ID:          bundle.ID,
			Name:        bundle.Name,
			Image:       bundle.Image,
			Type:        bundle.Type,
			Price:       bundle.Price,
			Quota:       bundle.Quota,
			Description: bundle.Description,
			EventDate:   bundle.EventDate.Format("2006-01-02"),
			IsAvailable: &isAvailable,
		}

		for _, bi := range bundle.BundleItems {
			bundleItem := dto.BundleItemResponse{
				ID:        bi.ID,
				MerchID:   bi.MerchID,
				MerchName: bi.Merch.Name,
			}

			data.BundleItems = append(data.BundleItems, bundleItem)
		}

		datas = append(datas, data)
	}

	return dto.BundlePaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailBundle(ctx context.Context, bundleID string) (dto.BundleResponse, error) {
	bundle, _, err := as.adminRepo.GetBundleByID(ctx, nil, bundleID)
	if err != nil {
		return dto.BundleResponse{}, dto.ErrBundleNotFound
	}

	b := dto.BundleResponse{
		ID:          bundle.ID,
		Name:        bundle.Name,
		Image:       bundle.Image,
		Type:        bundle.Type,
		Price:       bundle.Price,
		Quota:       bundle.Quota,
		Description: bundle.Description,
		EventDate:   bundle.EventDate.Format("2006-01-02"),
	}

	for _, bi := range bundle.BundleItems {
		bundleItem := dto.BundleItemResponse{
			ID:        bi.ID,
			MerchID:   bi.MerchID,
			MerchName: bi.Merch.Name,
		}

		for _, mi := range bi.Merch.MerchImages {
			bundleItem.MerchImages = append(bundleItem.MerchImages, dto.MerchImageResponse{
				ID:   mi.ID,
				Name: mi.Name,
			})
		}

		b.BundleItems = append(b.BundleItems, bundleItem)
	}

	return b, nil
}
func (as *AdminService) UpdateBundle(ctx context.Context, req dto.UpdateBundleRequest) (dto.BundleResponse, error) {
	bundle, flag, err := as.adminRepo.GetBundleByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.BundleResponse{}, dto.ErrBundleNotFound
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.BundleResponse{}, dto.ErrBundleNameTooShort
		}

		bundle.Name = req.Name
	}

	if req.Type != "" {
		if !entity.IsValidBundleType(req.Type) {
			return dto.BundleResponse{}, dto.ErrInvalidBundleType
		}

		bundle.Type = req.Type
	}

	if req.Price != nil {
		if *req.Price < 0 {
			return dto.BundleResponse{}, dto.ErrPriceOutOfBound
		}

		bundle.Price = *req.Price
	}

	if req.Quota != nil {
		if *req.Quota < 0 {
			return dto.BundleResponse{}, dto.ErrQuotaOutOfBound
		}

		bundle.Quota = *req.Quota
	}

	if req.FileHeader != nil && req.FileReader != nil {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(req.FileHeader.Filename), "."))
		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			return dto.BundleResponse{}, dto.ErrInvalidExtensionPhoto
		}

		if bundle.Image != "" {
			oldPath := filepath.Join("assets/bundle", bundle.Image)
			if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
				return dto.BundleResponse{}, dto.ErrDeleteOldImage
			}
		}

		bundleName := strings.ToLower(bundle.Name)
		bundleName = strings.ReplaceAll(bundleName, " ", "_")

		fileName := fmt.Sprintf("bundle_%s_%s.%s", time.Now().Format("060102150405"), bundleName, ext)

		saveDir := "assets/bundle"
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return dto.BundleResponse{}, dto.ErrCreateFile
		}
		savePath := filepath.Join(saveDir, fileName)

		out, err := os.Create(savePath)
		if err != nil {
			return dto.BundleResponse{}, dto.ErrCreateFile
		}
		defer out.Close()

		if _, err := io.Copy(out, req.FileReader); err != nil {
			return dto.BundleResponse{}, dto.ErrSaveFile
		}
		bundle.Image = fileName
	}

	if req.Description != "" {
		bundle.Description = req.Description
	}

	if req.EventDate != "" {
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			loc = time.FixedZone("UTC+7", 7*60*60)
		}
		eventDate, err := time.ParseInLocation("2006-01-02", req.EventDate, loc)
		if err != nil {
			return dto.BundleResponse{}, dto.ErrParseTime
		}

		bundle.EventDate = eventDate
	}

	updateItems := req.BundleItems != nil

	var newItems []dto.BundleItemResponse
	if updateItems {
		for _, biID := range req.BundleItems {
			var item dto.BundleItemResponse
			item.ID = uuid.New()

			merch, flag, err := as.adminRepo.GetMerchByID(ctx, nil, biID.String())
			if flag || err == nil {
				item.MerchID = &merch.ID
				item.MerchName = merch.Name
			}

			newItems = append(newItems, item)
		}
	}

	err = as.adminRepo.RunInTransaction(ctx, func(txRepo repository.IAdminRepository) error {
		if err := txRepo.UpdateBundle(ctx, nil, bundle); err != nil {
			return dto.ErrUpdateBundle
		}

		if updateItems {
			if err := txRepo.DeleteBundleItemsByBundleID(ctx, nil, bundle.ID.String()); err != nil {
				return dto.ErrDeleteBundleItemsByBundleID
			}

			for _, bi := range newItems {
				bundleItem := entity.BundleItem{
					ID:       bi.ID,
					BundleID: &bundle.ID,
				}

				if bi.MerchID != nil {
					bundleItem.MerchID = bi.MerchID
				}

				if err := txRepo.CreateBundleItem(ctx, nil, bundleItem); err != nil {
					return dto.ErrCreateBundleItem
				}
			}
		}

		return nil
	})
	if err != nil {
		return dto.BundleResponse{}, err
	}

	var respItems []dto.BundleItemResponse
	if updateItems {
		respItems = newItems
	} else {
		itemsEntity, err := as.adminRepo.GetBundleItemsByBundleID(ctx, nil, bundle.ID.String())
		if err != nil || !flag {
			return dto.BundleResponse{}, dto.ErrGetBundleItems
		}
		for _, e := range itemsEntity {
			var biResp dto.BundleItemResponse
			biResp.ID = e.ID
			if e.MerchID != nil {
				biResp.MerchID = e.MerchID
				biResp.MerchName = e.Merch.Name
			}

			respItems = append(respItems, biResp)
		}
	}

	return dto.BundleResponse{
		ID:          bundle.ID,
		Name:        bundle.Name,
		Image:       bundle.Image,
		Type:        bundle.Type,
		Price:       bundle.Price,
		Quota:       bundle.Quota,
		Description: bundle.Description,
		EventDate:   bundle.EventDate.Format("2006-01-02"),
		BundleItems: respItems,
	}, nil
}
func (as *AdminService) DeleteBundle(ctx context.Context, req dto.DeleteBundleRequest) (dto.BundleResponse, error) {
	deletedBundle, _, err := as.adminRepo.GetBundleByID(ctx, nil, req.BundleID)
	if err != nil {
		return dto.BundleResponse{}, dto.ErrBundleNotFound
	}

	err = as.adminRepo.RunInTransaction(ctx, func(txRepo repository.IAdminRepository) error {
		err = txRepo.DeleteBundleByID(ctx, nil, req.BundleID)
		if err != nil {
			return dto.ErrDeleteBundleByID
		}

		if err = txRepo.DeleteBundleItemsByBundleID(ctx, nil, req.BundleID); err != nil {
			return dto.ErrDeleteBundleItemsByBundleID
		}

		return nil
	})

	b := dto.BundleResponse{
		ID:          deletedBundle.ID,
		Name:        deletedBundle.Name,
		Image:       deletedBundle.Image,
		Type:        deletedBundle.Type,
		Price:       deletedBundle.Price,
		Quota:       deletedBundle.Quota,
		Description: deletedBundle.Description,
		EventDate:   deletedBundle.EventDate.Format("2006-01-02"),
	}

	for _, bi := range deletedBundle.BundleItems {
		bundleItem := dto.BundleItemResponse{
			ID:        bi.ID,
			MerchID:   bi.MerchID,
			MerchName: bi.Merch.Name,
		}

		b.BundleItems = append(b.BundleItems, bundleItem)
	}

	return b, nil
}

// User
func (as *AdminService) CreateStudentAmbassador(ctx context.Context, req dto.CreateStudentAmbassadorRequest) (dto.StudentAmbassadorResponse, error) {
	if req.Name == "" || req.ReferalCode == "" {
		return dto.StudentAmbassadorResponse{}, dto.ErrEmptyFields
	}

	_, flag, err := as.adminRepo.GetStudentAmbassadorByReferalCode(ctx, nil, req.ReferalCode)
	if err == nil || flag {
		return dto.StudentAmbassadorResponse{}, dto.ErrStudentAmbassadorAlreadyExists
	}

	if len(req.Name) < 3 {
		return dto.StudentAmbassadorResponse{}, dto.ErrStudentAmbassadorNameTooShort
	}

	studentAmbassador := entity.StudentAmbassador{
		ID:          uuid.New(),
		Name:        req.Name,
		ReferalCode: req.ReferalCode,
		Discount:    req.Discount,
		MaxReferal:  req.MaxReferal,
	}

	err = as.adminRepo.CreateStudentAmbassador(ctx, nil, studentAmbassador)
	if err != nil {
		return dto.StudentAmbassadorResponse{}, dto.ErrCreateStudentAmbassador
	}

	return dto.StudentAmbassadorResponse{
		ID:          studentAmbassador.ID,
		Name:        studentAmbassador.Name,
		ReferalCode: studentAmbassador.ReferalCode,
		Discount:    studentAmbassador.Discount,
		MaxReferal:  studentAmbassador.MaxReferal,
	}, nil
}
func (as *AdminService) GetAllStudentAmbassador(ctx context.Context) ([]dto.StudentAmbassadorResponse, error) {
	studentAmbassadors, err := as.adminRepo.GetAllStudentAmbassador(ctx, nil)
	if err != nil {
		return nil, dto.ErrGetAllStudentAmbassadorNoPagination
	}

	var datas []dto.StudentAmbassadorResponse
	for _, studentAmbassador := range studentAmbassadors {
		data := dto.StudentAmbassadorResponse{
			ID:          studentAmbassador.ID,
			Name:        studentAmbassador.Name,
			ReferalCode: studentAmbassador.ReferalCode,
			Discount:    studentAmbassador.Discount,
			MaxReferal:  studentAmbassador.MaxReferal,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllStudentAmbassadorWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.StudentAmbassadorPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllStudentAmbassadorWithPagination(ctx, nil, req)
	if err != nil {
		return dto.StudentAmbassadorPaginationResponse{}, dto.ErrGetAllStudentAmbassadorWithPagination
	}

	var datas []dto.StudentAmbassadorResponse
	for _, studentAmbassador := range dataWithPaginate.StudentAmbassadors {
		data := dto.StudentAmbassadorResponse{
			ID:          studentAmbassador.ID,
			Name:        studentAmbassador.Name,
			ReferalCode: studentAmbassador.ReferalCode,
			Discount:    studentAmbassador.Discount,
			MaxReferal:  studentAmbassador.MaxReferal,
		}

		datas = append(datas, data)
	}

	return dto.StudentAmbassadorPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailStudentAmbassador(ctx context.Context, studentAmbassadorID string) (dto.StudentAmbassadorResponse, error) {
	studentAmbassador, _, err := as.adminRepo.GetStudentAmbassadorByID(ctx, nil, studentAmbassadorID)
	if err != nil {
		return dto.StudentAmbassadorResponse{}, dto.ErrStudentAmbassadorNotFound
	}

	return dto.StudentAmbassadorResponse{
		ID:          studentAmbassador.ID,
		Name:        studentAmbassador.Name,
		ReferalCode: studentAmbassador.ReferalCode,
		Discount:    studentAmbassador.Discount,
		MaxReferal:  studentAmbassador.MaxReferal,
	}, nil
}
func (as *AdminService) UpdateStudentAmbassador(ctx context.Context, req dto.UpdateStudentAmbassadorRequest) (dto.StudentAmbassadorResponse, error) {
	studentAmbassador, flag, err := as.adminRepo.GetStudentAmbassadorByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.StudentAmbassadorResponse{}, dto.ErrStudentAmbassadorNotFound
	}

	if req.Name != "" {
		if len(req.Name) < 3 {
			return dto.StudentAmbassadorResponse{}, dto.ErrStudentAmbassadorNameTooShort
		}

		studentAmbassador.Name = req.Name
	}

	if req.ReferalCode != "" {
		_, flag, err := as.adminRepo.GetStudentAmbassadorByReferalCode(ctx, nil, req.ReferalCode)
		if err == nil || flag {
			return dto.StudentAmbassadorResponse{}, dto.ErrReferalCodeAlreadyExists
		}

		studentAmbassador.ReferalCode = req.ReferalCode
	}

	if req.Discount != nil {
		studentAmbassador.Discount = *req.Discount
	}

	if req.MaxReferal != nil {
		studentAmbassador.MaxReferal = *req.MaxReferal
	}

	err = as.adminRepo.UpdateStudentAmbassador(ctx, nil, studentAmbassador)
	if err != nil {
		return dto.StudentAmbassadorResponse{}, dto.ErrUpdateStudentAmbassador
	}

	res := dto.StudentAmbassadorResponse{
		ID:          studentAmbassador.ID,
		Name:        studentAmbassador.Name,
		ReferalCode: studentAmbassador.ReferalCode,
		Discount:    studentAmbassador.Discount,
		MaxReferal:  studentAmbassador.MaxReferal,
	}

	return res, nil
}
func (as *AdminService) DeleteStudentAmbassador(ctx context.Context, req dto.DeleteStudentAmbassadorRequest) (dto.StudentAmbassadorResponse, error) {
	deletedStudentAmbassador, _, err := as.adminRepo.GetStudentAmbassadorByID(ctx, nil, req.StudentAmbassadorID)
	if err != nil {
		return dto.StudentAmbassadorResponse{}, dto.ErrStudentAmbassadorNotFound
	}

	err = as.adminRepo.DeleteStudentAmbassadorByID(ctx, nil, req.StudentAmbassadorID)
	if err != nil {
		return dto.StudentAmbassadorResponse{}, dto.ErrDeleteStudentAmbassadorByID
	}

	res := dto.StudentAmbassadorResponse{
		ID:          deletedStudentAmbassador.ID,
		Name:        deletedStudentAmbassador.Name,
		ReferalCode: deletedStudentAmbassador.ReferalCode,
		Discount:    deletedStudentAmbassador.Discount,
		MaxReferal:  deletedStudentAmbassador.MaxReferal,
	}

	return res, nil
}

// Transaction & Ticket Form
func (as *AdminService) CreateTransactionTicket(ctx context.Context, req dto.CreateTransactionTicketRequest) (dto.TransactionResponse, error) {
	if len(req.TicketForms) == 0 {
		return dto.TransactionResponse{}, dto.ErrEmptyTicketForms
	}

	token := ctx.Value("Authorization").(string)

	userIDStr, err := as.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrGetUserIDFromToken
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrParseUUID
	}

	var transactionResponse dto.TransactionResponse
	err = as.adminRepo.RunInTransaction(ctx, func(txRepo repository.IAdminRepository) error {
		if !entity.IsValidItemType(req.ItemType) && req.ItemType != constants.ENUM_TICKET_ITEM_TYPE {
			return dto.ErrItemTypeMustBeTicket
		}

		ticket, found, err := txRepo.GetTicketByID(ctx, nil, req.TicketID.String())
		if err != nil || !found {
			return dto.ErrTicketNotFound
		}

		if ticket.Type != "main-event" {
			return dto.ErrTicketTypeMustBeMainEvent
		}

		if ticket.Quota <= 0 {
			return dto.ErrTicketSoldOut
		}

		transactionID := uuid.New()
		orderID := fmt.Sprintf("TEDX-%s", time.Now().Format("060102150405"))

		now := time.Now()

		transaction := entity.Transaction{
			ID:                transactionID,
			OrderID:           orderID,
			ItemType:          req.ItemType,
			TransactionStatus: "settlement",
			PaymentType:       "invitation",
			SettlementTime:    &now,
			UserID:            &userID,
			TicketID:          req.TicketID,
		}

		if err := txRepo.CreateTransaction(ctx, nil, transaction); err != nil {
			return dto.ErrCreateTransaction
		}

		for _, form := range req.TicketForms {
			if form.AudienceType == "" || form.Instansi == "" || form.Email == "" || form.FullName == "" || form.PhoneNumber == "" {
				return dto.ErrEmptyFields
			}

			if !entity.IsValidAudienceType(form.AudienceType) || form.AudienceType != "invited" {
				return dto.ErrMustBeInvitedGuest
			}

			if !entity.IsValidInstansi(form.Instansi) {
				return dto.ErrInvalidInstansi
			}

			if !helpers.IsValidEmail(form.Email) {
				return dto.ErrInvalidEmail
			}

			if len(form.FullName) < 5 {
				return dto.ErrUserFullNameTooShort
			}

			formattedPhone, err := helpers.StandardizePhoneNumber(form.PhoneNumber)
			if err != nil {
				return dto.ErrInvalidPhoneNumber
			}

			ticketFormID := uuid.New()
			ticketForm := entity.TicketForm{
				ID:            ticketFormID,
				AudienceType:  form.AudienceType,
				Instansi:      form.Instansi,
				Email:         form.Email,
				FullName:      form.FullName,
				PhoneNumber:   formattedPhone,
				LineID:        form.LineID,
				TransactionID: &transactionID,
			}

			if err := txRepo.UpdateTicketQuota(ctx, nil, ticket.ID.String(), ticket.Quota-1); err != nil {
				return dto.ErrUpdateTicket
			}
			if err := txRepo.CreateTicketForm(ctx, nil, ticketForm); err != nil {
				return dto.ErrCreateTicketForm
			}

			qrURL, err := helpers.GenerateQRCodeFile(ticketForm.ID.String(), ticketForm.ID.String()+".png")
			if err != nil {
				return dto.ErrGenerateQRCode
			}

			headerImage := fmt.Sprintf("%s/assets/header-e-ticket-mail.png", os.Getenv("BASE_URL"))
			emailData := struct {
				HeaderImage  string
				TicketID     string
				Status       string
				AttendeeName string
				Email        string
				AudienceType string
				BookingDate  string
				Price        string
				QRCode       string
			}{
				HeaderImage:  headerImage,
				TicketID:     transaction.ID.String(),
				Status:       "settlement",
				AttendeeName: ticketForm.FullName,
				Email:        ticketForm.Email,
				AudienceType: string(ticketForm.AudienceType),
				BookingDate:  time.Now().Format("02 Jan 2006 15:04"),
				Price:        "Rp 0",
				QRCode:       qrURL,
			}

			draftEmail, err := makeETicketEmail(emailData)
			if err != nil {
				return fmt.Errorf("error : %s", err)
			}

			err = utils.SendEmail(emailData.Email, draftEmail["subject"], draftEmail["body"])
			if err != nil {
				return fmt.Errorf("error : %s", err)
			}

			transactionResponse.TicketForms = append(transactionResponse.TicketForms, dto.TicketFormResponse{
				ID:           ticketFormID,
				AudienceType: ticketForm.AudienceType,
				Instansi:     ticketForm.Instansi,
				Email:        ticketForm.Email,
				FullName:     ticketForm.FullName,
				PhoneNumber:  ticketForm.PhoneNumber,
				LineID:       ticketForm.LineID,
			})
		}
		transactionResponse.ID = transactionID
		transactionResponse.OrderID = transaction.OrderID
		transactionResponse.ItemType = transaction.ItemType
		transactionResponse.TicketType = ticket.Type
		transactionResponse.TransactionStatus = transaction.TransactionStatus
		transactionResponse.PaymentType = transaction.PaymentType
		transactionResponse.SettlementTime = transaction.SettlementTime
		transactionResponse.UserID = transaction.UserID
		transactionResponse.TicketID = transaction.TicketID

		return nil
	})
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	return transactionResponse, nil
}
func (as *AdminService) GetAllTransactionTicket(ctx context.Context, transactionStatus, ticketCategory string) ([]dto.TransactionResponse, error) {
	transactions, err := as.adminRepo.GetAllTransaction(ctx, nil, transactionStatus, ticketCategory)
	if err != nil {
		return nil, dto.ErrGetAllTransactionNoPagination
	}

	var datas []dto.TransactionResponse
	for _, transaction := range transactions {
		data := dto.TransactionResponse{
			ID:                transaction.ID,
			OrderID:           transaction.OrderID,
			ItemType:          transaction.ItemType,
			TicketType:        entity.TicketType(transaction.Ticket.Type),
			ReferalCode:       transaction.ReferalCode,
			TransactionStatus: transaction.TransactionStatus,
			PaymentType:       transaction.PaymentType,
			SignatureKey:      transaction.SignatureKey,
			Acquire:           transaction.Acquire,
			SettlementTime:    transaction.SettlementTime,
			GrossAmount:       transaction.GrossAmount,
			UserID:            transaction.UserID,
			TicketID:          transaction.TicketID,
			BundleID:          transaction.BundleID,
		}

		for _, ticketForm := range transaction.TicketForms {
			data.TicketForms = append(data.TicketForms, dto.TicketFormResponse{
				ID:           ticketForm.ID,
				AudienceType: ticketForm.AudienceType,
				Instansi:     ticketForm.Instansi,
				Email:        ticketForm.Email,
				FullName:     ticketForm.FullName,
				PhoneNumber:  ticketForm.PhoneNumber,
				LineID:       ticketForm.LineID,
			})
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllTransactionTicketWithPagination(ctx context.Context, req dto.PaginationRequest, transactionStatus, ticketCategory string) (dto.TransactionTicketPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllTransactionWithPagination(ctx, nil, req, transactionStatus, ticketCategory)
	if err != nil {
		return dto.TransactionTicketPaginationResponse{}, dto.ErrGetAllUserWithPagination
	}

	var datas []dto.TransactionResponse
	for _, transaction := range dataWithPaginate.Transactions {
		data := dto.TransactionResponse{
			ID:                transaction.ID,
			OrderID:           transaction.OrderID,
			ItemType:          transaction.ItemType,
			TicketType:        entity.TicketType(transaction.Ticket.Type),
			ReferalCode:       transaction.ReferalCode,
			TransactionStatus: transaction.TransactionStatus,
			PaymentType:       transaction.PaymentType,
			SignatureKey:      transaction.SignatureKey,
			Acquire:           transaction.Acquire,
			SettlementTime:    transaction.SettlementTime,
			GrossAmount:       transaction.GrossAmount,
			UserID:            transaction.UserID,
			TicketID:          transaction.TicketID,
			BundleID:          transaction.BundleID,
		}

		for _, ticketForm := range transaction.TicketForms {
			data.TicketForms = append(data.TicketForms, dto.TicketFormResponse{
				ID:           ticketForm.ID,
				AudienceType: ticketForm.AudienceType,
				Instansi:     ticketForm.Instansi,
				Email:        ticketForm.Email,
				FullName:     ticketForm.FullName,
				PhoneNumber:  ticketForm.PhoneNumber,
				LineID:       ticketForm.LineID,
			})
		}

		datas = append(datas, data)
	}

	return dto.TransactionTicketPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailTransactionTicket(ctx context.Context, transactionTicketID string) (dto.TransactionResponse, error) {
	transaction, _, err := as.adminRepo.GetTransactionByID(ctx, nil, transactionTicketID)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrTicketFormNotFound
	}

	res := dto.TransactionResponse{
		ID:                transaction.ID,
		OrderID:           transaction.OrderID,
		ItemType:          transaction.ItemType,
		TicketType:        entity.TicketType(transaction.Ticket.Type),
		ReferalCode:       transaction.ReferalCode,
		TransactionStatus: transaction.TransactionStatus,
		PaymentType:       transaction.PaymentType,
		SignatureKey:      transaction.SignatureKey,
		Acquire:           transaction.Acquire,
		SettlementTime:    transaction.SettlementTime,
		GrossAmount:       transaction.GrossAmount,
		UserID:            transaction.UserID,
		TicketID:          transaction.TicketID,
		BundleID:          transaction.BundleID,
	}

	for _, ticketForm := range transaction.TicketForms {
		res.TicketForms = append(res.TicketForms, dto.TicketFormResponse{
			ID:           ticketForm.ID,
			AudienceType: ticketForm.AudienceType,
			Instansi:     ticketForm.Instansi,
			Email:        ticketForm.Email,
			FullName:     ticketForm.FullName,
			PhoneNumber:  ticketForm.PhoneNumber,
			LineID:       ticketForm.LineID,
		})
	}

	return res, nil
}

// Check-in
func (as *AdminService) GetDetailTicketCheckIn(ctx context.Context, ticketFormIDStr string) (dto.TicketCheckInResponse, error) {
	ticketForm, found, err := as.adminRepo.GetTicketFormByID(ctx, nil, ticketFormIDStr)
	if err != nil || !found {
		return dto.TicketCheckInResponse{}, dto.ErrTicketFormNotFound
	}

	if ticketForm.TransactionID == nil || ticketForm.Transaction.TicketID == nil {
		return dto.TicketCheckInResponse{}, dto.ErrTransactionNotFound
	}

	status := false
	if len(ticketForm.GuestAttendances) != 0 {
		status = true
	}

	var emailChecker string
	if status {
		token := ctx.Value("Authorization").(string)

		adminIDStr, err := as.jwtService.GetUserIDByToken(token)
		if err != nil {
			return dto.TicketCheckInResponse{}, dto.ErrGetUserIDFromToken
		}

		admin, found, err := as.adminRepo.GetUserByID(ctx, nil, adminIDStr)
		if err != nil || !found {
			return dto.TicketCheckInResponse{}, dto.ErrUserNotFound
		}

		emailChecker = admin.Email
	}

	res := dto.TicketCheckInResponse{
		TicketID:      *ticketForm.Transaction.TicketID,
		TransactionID: *ticketForm.TransactionID,
		TicketName:    ticketForm.Transaction.Ticket.Name,
		TicketType:    ticketForm.Transaction.Ticket.Type,
		AudienceType:  ticketForm.AudienceType,
		Email:         ticketForm.Email,
		FullName:      ticketForm.FullName,
		PhoneNumber:   ticketForm.PhoneNumber,
		LineID:        ticketForm.LineID,
		Status:        status,
		EmailChecker:  emailChecker,
	}

	return res, nil
}
func (as *AdminService) CheckIn(ctx context.Context, ticketFormIDStr string) error {
	ticketForm, found, err := as.adminRepo.GetTicketFormByID(ctx, nil, ticketFormIDStr)
	if err != nil || !found {
		return dto.ErrTicketFormNotFound
	}

	if len(ticketForm.GuestAttendances) > 0 {
		return dto.ErrAlreadyCheckedIn
	}

	token := ctx.Value("Authorization").(string)

	adminIDStr, err := as.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.ErrGetUserIDFromToken
	}

	adminID, err := uuid.Parse(adminIDStr)
	if err != nil {
		return dto.ErrParseUUID
	}

	guestAttendance := entity.GuestAttendance{
		ID:           uuid.New(),
		TicketFormID: &ticketForm.ID,
		CheckedBy:    &adminID,
	}

	err = as.adminRepo.CreateGuestAttendance(ctx, nil, guestAttendance)
	if err != nil {
		return dto.ErrCreateGuestAttendance
	}

	return nil
}
func (as *AdminService) GetAllTicketCheckIn(ctx context.Context, filter dto.CheckInFilterQuery) ([]dto.TicketCheckInResponse, error) {
	ticketForms, err := as.adminRepo.GetAllTicketForm(ctx, nil, filter)
	if err != nil {
		return nil, dto.ErrGetAllTicketCheckInNoPagination
	}

	var datas []dto.TicketCheckInResponse
	for _, ticketForm := range ticketForms {
		var emailChecker string
		if len(ticketForm.GuestAttendances) > 0 {
			emailChecker = ticketForm.GuestAttendances[0].CheckedByUser.Email
		}

		data := dto.TicketCheckInResponse{
			TicketID:      *ticketForm.Transaction.TicketID,
			TransactionID: *ticketForm.TransactionID,
			TicketName:    ticketForm.Transaction.Ticket.Name,
			TicketType:    ticketForm.Transaction.Ticket.Type,
			AudienceType:  ticketForm.AudienceType,
			Email:         ticketForm.Email,
			FullName:      ticketForm.FullName,
			PhoneNumber:   ticketForm.PhoneNumber,
			LineID:        ticketForm.LineID,
			Status:        true,
			EmailChecker:  emailChecker,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (as *AdminService) GetAllTicketCheckInWithPagination(ctx context.Context, req dto.PaginationRequest, filter dto.CheckInFilterQuery) (dto.TicketFormPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllTicketFormWithPagination(ctx, nil, req, filter)
	if err != nil {
		return dto.TicketFormPaginationResponse{}, dto.ErrGetAllTicketCheckInWithPagination
	}

	var datas []dto.TicketCheckInResponse
	for _, ticketForm := range dataWithPaginate.TicketForms {
		var emailChecker string
		if len(ticketForm.GuestAttendances) > 0 {
			emailChecker = ticketForm.GuestAttendances[0].CheckedByUser.Email
		}

		data := dto.TicketCheckInResponse{
			TicketID:      *ticketForm.Transaction.TicketID,
			TransactionID: *ticketForm.TransactionID,
			TicketName:    ticketForm.Transaction.Ticket.Name,
			TicketType:    ticketForm.Transaction.Ticket.Type,
			AudienceType:  ticketForm.AudienceType,
			Email:         ticketForm.Email,
			FullName:      ticketForm.FullName,
			PhoneNumber:   ticketForm.PhoneNumber,
			LineID:        ticketForm.LineID,
			Status:        true,
			EmailChecker:  emailChecker,
		}

		datas = append(datas, data)
	}

	return dto.TicketFormPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
