package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"

	m "github.com/Amierza/TedXBackend/config/midtrans"
	"github.com/Amierza/TedXBackend/constants"
	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/entity"
	"github.com/Amierza/TedXBackend/helpers"
	"github.com/Amierza/TedXBackend/repository"
	"github.com/Amierza/TedXBackend/utils"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type (
	IUserService interface {
		// Authentication
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)

		// User
		GetDetailUser(ctx context.Context) (dto.UserResponse, error)
		UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.UserResponse, error)

		// Ticket
		GetAllTicket(ctx context.Context) ([]dto.TicketResponse, error)

		// Sponsorship
		GetAllSponsorship(ctx context.Context) ([]dto.SponsorshipResponse, error)

		// Speaker
		GetAllSpeaker(ctx context.Context) ([]dto.SpeakerResponse, error)

		// Merch
		GetAllMerch(ctx context.Context) ([]dto.MerchResponse, error)

		// Bundle
		GetAllBundle(ctx context.Context) ([]dto.BundleResponse, error)

		// Check Referal Code
		CheckReferalCode(ctx context.Context, req dto.CheckReferalCodeRequest) (dto.StudentAmbassadorResponse, error)

		// Snap for trigger midtrans
		CreateTransactionTicket(ctx context.Context, req dto.CreateTransactionTicketRequest) (dto.TransactionResponse, error)

		// Webhook for Midtrans
		UpdateTransactionTicket(ctx context.Context, req dto.UpdateMidtransTransactionTicketRequest) error
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
func (us *UserService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	token := ctx.Value("Authorization").(string)

	userIDStr, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserIDFromToken
	}

	req.ID = userIDStr

	user, flag, err := us.userRepo.GetUserByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.UserResponse{}, dto.ErrUserNotFound
	}

	if req.Email != "" {
		_, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
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

	err = us.userRepo.UpdateUser(ctx, nil, user)
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
func (us *UserService) GetAllBundle(ctx context.Context) ([]dto.BundleResponse, error) {
	bundleType := "bundle merch ticket"

	bundles, err := us.userRepo.GetAllBundle(ctx, nil, bundleType)
	if err != nil {
		return nil, dto.ErrGetAllBundleNoPagination
	}

	var datas []dto.BundleResponse
	for _, bundle := range bundles {
		data := dto.BundleResponse{
			ID:    bundle.ID,
			Name:  bundle.Name,
			Image: bundle.Image,
			Type:  bundle.Type,
			Price: bundle.Price,
			Quota: bundle.Quota,
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

// Check Referal Code
func (us *UserService) CheckReferalCode(ctx context.Context, req dto.CheckReferalCodeRequest) (dto.StudentAmbassadorResponse, error) {
	sa, found, err := us.userRepo.GetStudentAmbassadorByReferalCode(ctx, nil, req.ReferalCode)
	if err != nil || !found {
		return dto.StudentAmbassadorResponse{}, dto.ErrInvalidReferalCode
	}

	res := dto.StudentAmbassadorResponse{
		ID:          sa.ID,
		Name:        sa.Name,
		ReferalCode: sa.ReferalCode,
		Discount:    sa.Discount,
		MaxReferal:  sa.MaxReferal,
	}

	return res, nil
}

// Snap for trigger midtrans
func (us *UserService) CreateTransactionTicket(ctx context.Context, req dto.CreateTransactionTicketRequest) (dto.TransactionResponse, error) {
	if len(req.TicketForms) == 0 {
		return dto.TransactionResponse{}, dto.ErrEmptyTicketForms
	}

	token := ctx.Value("Authorization").(string)

	userIDStr, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrGetUserIDFromToken
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return dto.TransactionResponse{}, dto.ErrParseUUID
	}

	user, found, err := us.userRepo.GetUserByID(ctx, nil, userIDStr)
	if err != nil || !found {
		return dto.TransactionResponse{}, dto.ErrUserNotFound
	}

	var transactionResponse dto.TransactionResponse
	err = us.userRepo.RunInTransaction(ctx, func(txRepo repository.IUserRepository) error {
		if req.ReferalCode != "" {
			sa, found, err := txRepo.GetStudentAmbassadorByReferalCode(ctx, nil, req.ReferalCode)
			if err != nil || !found {
				return dto.ErrInvalidReferalCode
			}

			err = txRepo.UpdateMaxReferal(ctx, nil, sa.ID.String(), sa.MaxReferal-1)
			if err != nil {
				return dto.ErrUpdateMaxReferal
			}
		}

		if req.Total <= 0 {
			return dto.ErrTotalOutOfBound
		}

		if !entity.IsValidItemType(req.ItemType) || (req.ItemType != constants.ENUM_TICKET_ITEM_TYPE && req.ItemType != constants.ENUM_BUNDLE_ITEM_TYPE) {
			return dto.ErrItemTypeMustBeTicketOrBundle
		}

		var (
			ticket entity.Ticket
			bundle entity.Bundle
		)

		if req.TicketID != nil && *req.TicketID != uuid.Nil {
			t, found, err := txRepo.GetTicketByID(ctx, nil, req.TicketID.String())
			if err != nil || !found {
				return dto.ErrTicketNotFound
			}

			if t.Quota <= 0 {
				return dto.ErrTicketSoldOut
			}

			ticket = t
		}

		if req.BundleID != nil && *req.BundleID != uuid.Nil {
			b, found, err := txRepo.GetBundleByID(ctx, nil, req.BundleID.String())
			if err != nil || !found {
				return dto.ErrTicketNotFound
			}

			if b.Quota <= 0 {
				return dto.ErrBundleSoldOut
			}

			bundle = b
		}

		transactionID := uuid.New()
		orderID := fmt.Sprintf("TEDX-%s", time.Now().Format("060102150405"))

		transaction := entity.Transaction{
			ID:          transactionID,
			OrderID:     orderID,
			ItemType:    req.ItemType,
			ReferalCode: req.ReferalCode,
			UserID:      &userID,
			TicketID:    req.TicketID,
			BundleID:    req.BundleID,
		}

		if err := txRepo.CreateTransaction(ctx, nil, transaction); err != nil {
			return dto.ErrCreateTransaction
		}

		for _, form := range req.TicketForms {
			if form.AudienceType == "" || form.Instansi == "" || form.Email == "" || form.FullName == "" || form.PhoneNumber == "" {
				return dto.ErrEmptyFields
			}

			if !entity.IsValidAudienceType(form.AudienceType) || form.AudienceType != "regular" {
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

			if req.BundleID != nil && *req.BundleID != uuid.Nil {
				if err := txRepo.UpdateBundleQuota(ctx, nil, bundle.ID.String(), bundle.Quota-1); err != nil {
					return dto.ErrUpdateBundleQuota
				}
			}

			if req.TicketID != nil && *req.TicketID != uuid.Nil {
				if err := txRepo.UpdateTicketQuota(ctx, nil, ticket.ID.String(), ticket.Quota-1); err != nil {
					return dto.ErrUpdateTicketQuota
				}
			}

			if err := txRepo.CreateTicketForm(ctx, nil, ticketForm); err != nil {
				return dto.ErrCreateTicketForm
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

			r := &snap.Request{
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  orderID,
					GrossAmt: int64(req.Total),
				},
				CustomerDetail: &midtrans.CustomerDetails{
					FName: user.Name,
					LName: user.Name,
					Email: user.Email,
					Phone: "",
				},
			}

			snapResp, midtransErr := m.SnapClient.CreateTransaction(r)
			if midtransErr != nil {
				return midtransErr
			}

			transactionResponse.ID = transactionID
			transactionResponse.OrderID = transaction.OrderID
			transactionResponse.ItemType = transaction.ItemType
			transactionResponse.UserID = transaction.UserID
			transactionResponse.TicketID = transaction.TicketID
			transactionResponse.BundleID = transaction.BundleID
			transactionResponse.Token = snapResp.Token
			transactionResponse.RedirectURL = snapResp.RedirectURL
		}

		return nil
	})
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	return transactionResponse, nil
}

// Webhook for Midtrans
func makeETicketEmail(data struct {
	TicketID     string
	Status       string
	AttendeeName string
	Email        string
	AudienceType string
	BookingDate  string
	Price        string
}) (map[string]string, error) {
	readHTML, err := os.ReadFile("utils/email_template/e-ticket-mail.html")
	if err != nil {
		return nil, fmt.Errorf("failed to read HTML template: %w", err)
	}

	tmpl, err := template.New("eticket").Parse(string(readHTML))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML template: %w", err)
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, fmt.Errorf("failed to execute HTML template: %w", err)
	}

	draftEmail := map[string]string{
		"subject": "tedxuniversitasairlangga",
		"body":    strMail.String(),
	}
	return draftEmail, nil
}
func (us *UserService) UpdateTransactionTicket(ctx context.Context, req dto.UpdateMidtransTransactionTicketRequest) error {
	transaction, found, err := us.userRepo.GetTransactionByOrderID(ctx, nil, req.OrderID)
	if err != nil || !found {
		return dto.ErrTransactionNotFound
	}

	switch req.TransactionStatus {
	case "settlement":
		transaction.TransactionStatus = "settlement"
		settlementTime, err := time.Parse("2006-01-02 15:04:05", req.SettlementTime)
		if err != nil {
			return dto.ErrParseTime
		}
		transaction.SettlementTime = &settlementTime
		transaction.PaymentType = req.PaymentType
		transaction.SignatureKey = req.SignatureKey
		transaction.Acquire = req.Aquirer
		grossAmount, err := strconv.ParseFloat(req.GrossAmount, 64)
		if err != nil {
			return fmt.Errorf("invalid gross amount: %w", err)
		}
		transaction.GrossAmount = grossAmount

	case "pending":
		transaction.TransactionStatus = "pending"

	case "deny", "failure":
		transaction.TransactionStatus = "failed"

	case "cancel":
		transaction.TransactionStatus = "cancelled"

	case "expire":
		transaction.TransactionStatus = "expired"

	default:
		return dto.ErrUnknownTransactionStatus
	}

	err = us.userRepo.UpdateTransactionTicket(ctx, nil, transaction)
	if err != nil {
		return dto.ErrUpdateTransactionTicket
	}

	for _, form := range transaction.TicketForms {
		emailData := struct {
			TicketID     string
			Status       string
			AttendeeName string
			Email        string
			AudienceType string
			BookingDate  string
			Price        string
		}{
			TicketID:     transaction.ID.String(),
			Status:       transaction.TransactionStatus,
			AttendeeName: form.FullName,
			Email:        form.Email,
			AudienceType: string(form.AudienceType),
			BookingDate:  transaction.CreatedAt.Format("02 Jan 2006 15:04"),
			Price:        fmt.Sprintf("Rp %.0f", transaction.GrossAmount),
		}

		draftEmail, err := makeETicketEmail(emailData)
		if err != nil {
			return dto.ErrMakeETicketEmail
		}

		if err := utils.SendEmail(emailData.Email, draftEmail["subject"], draftEmail["body"]); err != nil {
			log.Printf("gagal kirim email ke %s: %v", emailData.Email, err)
			continue
		}
	}

	return nil
}
