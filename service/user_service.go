package service

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"time"

	m "github.com/Amierza/TedXBackend/config/midtrans"
	"github.com/Amierza/TedXBackend/constants"
	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/entity"
	"github.com/Amierza/TedXBackend/helpers"
	"github.com/Amierza/TedXBackend/repository"
	"github.com/Amierza/TedXBackend/utils"
	emailtemplate "github.com/Amierza/TedXBackend/utils/email_template"
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
		GetDetailTicket(ctx context.Context, ticketID string) (dto.TicketResponse, error)

		// Sponsorship
		GetAllSponsorship(ctx context.Context) ([]dto.SponsorshipResponse, error)

		// Speaker
		GetAllSpeaker(ctx context.Context) ([]dto.SpeakerResponse, error)

		// Merch
		GetAllMerch(ctx context.Context) ([]dto.MerchResponse, error)
		GetDetailMerch(ctx context.Context, merchID string) (dto.MerchResponse, error)

		// Bundle
		GetAllBundle(ctx context.Context, bundleType string) ([]dto.BundleResponse, error)
		GetDetailBundle(ctx context.Context, bundleID string) (dto.BundleResponse, error)

		// Check Referal Code
		CheckReferalCode(ctx context.Context, req dto.CheckReferalCodeRequest) (dto.StudentAmbassadorResponse, error)

		// Snap for trigger midtrans
		CreateTransactionTicket(ctx context.Context, req dto.CreateTransactionTicketRequest) (dto.TransactionResponse, error)

		// Webhook for Midtrans
		UpdateTransactionTicket(ctx context.Context, req dto.UpdateMidtransTransactionTicketRequest) error

		// Hisoty Transactions
		GetAllTransactions(ctx context.Context) ([]dto.TransactionResponse, error)
		GetDetailTransactions(ctx context.Context, id string) (dto.TransactionResponse, error)
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
		isAvailable := ticket.Quota-ticket.QuotaFilled > 0 && time.Now().After(ticket.EventStartDate) && time.Now().Before(ticket.EventEndDate)
		data := dto.TicketResponse{
			ID:             ticket.ID.String(),
			Name:           ticket.Name,
			Type:           ticket.Type,
			Price:          ticket.Price,
			BundleQuota:    ticket.Bundle_Quota,
			Quota:          ticket.Quota,
			QuotaFilled:    ticket.QuotaFilled,
			QuotaAvailable: ticket.Quota - ticket.QuotaFilled,
			Image:          ticket.Image,
			Description:    ticket.Description,
			EventStartDate: ticket.EventStartDate.Format("2006-01-02 15:04:05"),
			EventEndDate:   ticket.EventEndDate.Format("2006-01-02 15:04:05"),
			IsAvailable:    &isAvailable,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (us *UserService) GetDetailTicket(ctx context.Context, ticketID string) (dto.TicketResponse, error) {
	ticket, _, err := us.userRepo.GetTicketByID(ctx, nil, ticketID)
	if err != nil {
		return dto.TicketResponse{}, dto.ErrTicketNotFound
	}

	isAvailable := ticket.Quota-ticket.QuotaFilled > 0 && time.Now().After(ticket.EventStartDate) && time.Now().Before(ticket.EventEndDate)

	return dto.TicketResponse{
		ID:             ticket.ID.String(),
		Name:           ticket.Name,
		Type:           ticket.Type,
		Price:          ticket.Price,
		Quota:          ticket.Quota,
		BundleQuota:    ticket.Bundle_Quota,
		QuotaFilled:    ticket.QuotaFilled,
		QuotaAvailable: ticket.Quota - ticket.QuotaFilled,
		Image:          ticket.Image,
		Description:    ticket.Description,
		EventStartDate: ticket.EventStartDate.Format("2006-01-02 15:04:05"),
		EventEndDate:   ticket.EventEndDate.Format("2006-01-02 15:04:05"),
		IsAvailable:    &isAvailable,
	}, nil
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
			Size:     string(sponsorship.Size),
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
		var status = true

		if merch.Stock < 1 {
			status = false
		}

		data := dto.MerchResponse{
			ID:          merch.ID,
			Name:        merch.Name,
			Stock:       merch.Stock,
			Price:       merch.Price,
			Status:      status,
			Description: merch.Description,
			Category:    merch.Category,
			Images:      merchImages,
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (us *UserService) GetDetailMerch(ctx context.Context, merchID string) (dto.MerchResponse, error) {
	merch, _, err := us.userRepo.GetMerchByID(ctx, nil, merchID)
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

// Bundle
func (us *UserService) GetAllBundle(ctx context.Context, bundleType string) ([]dto.BundleResponse, error) {
	// bundleType := ""

	bundles, err := us.userRepo.GetAllBundle(ctx, nil, bundleType)
	if err != nil {
		return nil, dto.ErrGetAllBundleNoPagination
	}

	var datas []dto.BundleResponse
	for _, bundle := range bundles {
		isAvailable := bundle.Quota-bundle.QuotaFilled > 0 && time.Now().After(bundle.EventStartDate) && time.Now().Before(bundle.EventEndDate)
		data := dto.BundleResponse{
			ID:             bundle.ID,
			Name:           bundle.Name,
			Image:          bundle.Image,
			Type:           bundle.Type,
			Price:          bundle.Price,
			Quota:          bundle.Quota,
			QuotaAvailable: bundle.Quota - bundle.QuotaFilled,
			QuotaFilled:    bundle.QuotaFilled,
			Description:    bundle.Description,
			IsAvailable:    &isAvailable,
			EventStartDate: bundle.EventStartDate.Format("2006-01-02 15:04:05"),
			EventEndDate:   bundle.EventEndDate.Format("2006-01-02 15:04:05"),
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

			data.BundleItems = append(data.BundleItems, bundleItem)
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (us *UserService) GetDetailBundle(ctx context.Context, bundleID string) (dto.BundleResponse, error) {
	bundle, _, err := us.userRepo.GetBundleByID(ctx, nil, bundleID)
	if err != nil {
		return dto.BundleResponse{}, dto.ErrBundleNotFound
	}

	isAvailable := bundle.Quota-bundle.QuotaFilled > 0 && time.Now().After(bundle.EventStartDate) && time.Now().Before(bundle.EventEndDate)
	b := dto.BundleResponse{
		ID:             bundle.ID,
		Name:           bundle.Name,
		Image:          bundle.Image,
		Type:           bundle.Type,
		Price:          bundle.Price,
		Quota:          bundle.Quota,
		QuotaAvailable: bundle.Quota - bundle.QuotaFilled,
		QuotaFilled:    bundle.QuotaFilled,
		Description:    bundle.Description,
		IsAvailable:    &isAvailable,
		EventStartDate: bundle.EventStartDate.Format("2006-01-02 15:04:05"),
		EventEndDate:   bundle.EventEndDate.Format("2006-01-02 15:04:05"),
	}

	for _, bi := range bundle.BundleItems {
		bundleItem := dto.BundleItemResponse{
			ID:               bi.ID,
			MerchID:          bi.MerchID,
			MerchName:        bi.Merch.Name,
			MerchPrice:       bi.Merch.Price,
			MerchDescription: bi.Merch.Description,
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

// Check Referal Code
func (us *UserService) CheckReferalCode(ctx context.Context, req dto.CheckReferalCodeRequest) (dto.StudentAmbassadorResponse, error) {
	ticket, found, err := us.userRepo.GetTicketByID(ctx, nil, req.TicketID)
	if err != nil || !found {
		return dto.StudentAmbassadorResponse{}, dto.ErrTicketNotFound
	}

	if !strings.EqualFold(ticket.Name, "normal") {
		return dto.StudentAmbassadorResponse{}, dto.ErrInvalidReferalCode
	}

	if ticket.Type != entity.MainEvent {
		return dto.StudentAmbassadorResponse{}, dto.ErrInvalidReferalCode
	}

	sa, found, err := us.userRepo.GetStudentAmbassadorByReferalCode(ctx, nil, req.ReferalCode)
	if err != nil || !found {
		return dto.StudentAmbassadorResponse{}, dto.ErrInvalidReferalCode
	}

	sisaQuota := sa.MaxReferal - sa.QuotaFilled
	if sisaQuota-req.TotalTicketForm < 0 {
		return dto.StudentAmbassadorResponse{}, fmt.Errorf("referal is limit")
	}

	res := dto.StudentAmbassadorResponse{
		ID:          sa.ID,
		Name:        sa.Name,
		ReferalCode: sa.ReferalCode,
		Discount:    sa.Discount,
		MaxReferal:  sa.MaxReferal,
		QuotaFilled: sa.QuotaFilled,
	}

	if sa.MaxReferal-sa.QuotaFilled <= 0 {
		res.IsAvailable = false
	} else {
		res.IsAvailable = true
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

	now := time.Now()

	var transactionResponse dto.TransactionResponse
	err = us.userRepo.RunInTransaction(ctx, func(txRepo repository.IUserRepository) error {
		var (
			ticket            entity.Ticket
			bundle            entity.Bundle
			studentAmbassador entity.StudentAmbassador
		)

		if req.ReferalCode != "" {
			sa, found, err := txRepo.GetStudentAmbassadorByReferalCode(ctx, nil, req.ReferalCode)
			if err != nil || !found {
				return dto.ErrInvalidReferalCode
			}

			if sa.MaxReferal-sa.QuotaFilled <= 0 {
				return dto.ErrReferalCodeSoldOut
			}

			studentAmbassador = sa
		}

		if req.Total <= 0 {
			return dto.ErrTotalOutOfBound
		}

		if !entity.IsValidItemType(req.ItemType) || (req.ItemType != constants.ENUM_TICKET_ITEM_TYPE && req.ItemType != constants.ENUM_BUNDLE_ITEM_TYPE) {
			return dto.ErrItemTypeMustBeTicketOrBundle
		}

		if req.TicketID != nil && *req.TicketID != uuid.Nil {
			t, found, err := txRepo.GetTicketByID(ctx, nil, req.TicketID.String())
			if err != nil || !found {
				return dto.ErrTicketNotFound
			}

			if t.Quota-t.QuotaFilled <= 0 {
				return dto.ErrTicketSoldOut
			}

			if now.Before(t.EventStartDate) || now.After(t.EventEndDate) {
				return fmt.Errorf("failed event ticket not available")
			}

			ticket = t
		}

		if req.BundleID != nil && *req.BundleID != uuid.Nil {
			b, found, err := txRepo.GetBundleByID(ctx, nil, req.BundleID.String())
			if err != nil || !found {
				return dto.ErrTicketNotFound
			}

			if b.Quota-b.QuotaFilled <= 0 {
				return dto.ErrBundleSoldOut
			}

			if now.Before(b.EventStartDate) || now.After(b.EventEndDate) {
				return fmt.Errorf("failed event bundle not available")
			}

			bundle = b
		}

		transactionID := uuid.New()
		orderID := fmt.Sprintf("TEDX-%s%03d", time.Now().Format("060102150405"), rand.Intn(1000))

		transaction := entity.Transaction{
			ID:                transactionID,
			OrderID:           orderID,
			ItemType:          req.ItemType,
			TransactionStatus: "pending",
			ReferalCode:       req.ReferalCode,
			UserID:            &userID,
			TicketID:          req.TicketID,
			BundleID:          req.BundleID,
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

			if req.ReferalCode != "" {
				sisaQuota := studentAmbassador.MaxReferal - studentAmbassador.QuotaFilled
				reqQuota := len(req.TicketForms)

				if reqQuota > sisaQuota {
					return fmt.Errorf("invalid referal quota")
				}

				err = txRepo.UpdateSAQuotaFilled(ctx, nil, studentAmbassador.ID.String(), 1)
				if err != nil {
					return dto.ErrUpdateSAQuotaFilled
				}
			}

			if req.BundleID != nil && *req.BundleID != uuid.Nil {
				sisaQuota := bundle.Quota - bundle.QuotaFilled
				reqQuota := len(req.TicketForms)

				if reqQuota > sisaQuota {
					return fmt.Errorf("invalid bundle quota")
				}

				if err := txRepo.UpdateBundleQuota(ctx, nil, bundle.ID.String(), 1); err != nil {
					return dto.ErrUpdateBundleQuota
				}
			}

			if req.TicketID != nil && *req.TicketID != uuid.Nil {
				sisaQuota := ticket.Quota - ticket.QuotaFilled
				reqQuota := len(req.TicketForms)

				if reqQuota > sisaQuota {
					return fmt.Errorf("invalid ticket quota")
				}

				if err := txRepo.UpdateTicketQuota(ctx, nil, ticket.ID.String(), 1); err != nil {
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
			transactionResponse.TicketType = ticket.Type
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
	HeaderImage  string
	TicketID     string
	TicketName   string
	TicketType   string
	Status       string
	AttendeeName string
	Email        string
	AudienceType string
	BookingDate  string
	Price        string
	QRCode       string
}) (map[string]string, error) {
	tmpl, err := template.New("eticket").Parse(emailtemplate.EticketHTML)
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
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			loc = time.FixedZone("UTC+7", 7*60*60)
		}
		settlementTime, err := time.ParseInLocation("2006-01-02 15:04:05", req.SettlementTime, loc)
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

		err = us.userRepo.UpdateTransactionTicket(ctx, nil, transaction)
		if err != nil {
			return dto.ErrUpdateTransactionTicket
		}

		sentEmails := make(map[string]bool)
		for _, form := range transaction.TicketForms {
			// if sentEmails[form.Email] {
			// 	continue
			// }
			sentEmails[form.Email] = true

			qrURL, err := helpers.GenerateQRCodeFile(form.ID.String(), form.ID.String()+".png")
			if err != nil {
				return dto.ErrGenerateQRCode
			}

			// headerImage := fmt.Sprintf("%s/assets_static/header-e-ticket-mail.png", os.Getenv("BASE_URL"))
			headerImage := "https://tedxuniversitasairlangga.com/images/header-e-ticket-mail.png"
			emailData := struct {
				HeaderImage  string
				TicketID     string
				TicketName   string
				TicketType   string
				Status       string
				AttendeeName string
				Email        string
				AudienceType string
				BookingDate  string
				Price        string
				QRCode       string
			}{
				HeaderImage:  headerImage,
				TicketID:     transaction.OrderID,
				TicketName:   transaction.Ticket.Name,
				TicketType:   string(transaction.Ticket.Type),
				Status:       transaction.TransactionStatus,
				AttendeeName: form.FullName,
				Email:        form.Email,
				AudienceType: string(form.AudienceType),
				BookingDate:  transaction.CreatedAt.Format("02 Jan 2006 15:04"),
				Price:        fmt.Sprintf("Rp %.0f", transaction.GrossAmount),
				QRCode:       qrURL,
			}

			draftEmail, err := makeETicketEmail(emailData)
			if err != nil {
				return dto.ErrMakeETicketEmail
			}

			err = utils.SendEmail(emailData.Email, draftEmail["subject"], draftEmail["body"])
			if err != nil {
				return dto.ErrSendEmail
			}
		}

		return nil

	case "pending":
		transaction.TransactionStatus = "pending"

	case "deny", "failure", "cancel", "expire":
		if transaction.ReferalCode != "" {
			sa, found, _ := us.userRepo.GetStudentAmbassadorByReferalCode(ctx, nil, transaction.ReferalCode)
			if found {
				if err := us.userRepo.UpdateSAQuotaFilled(ctx, nil, sa.ID.String(), -len(transaction.TicketForms)); err != nil {
					return dto.ErrUpdateSAQuotaFilled
				}
			}
		}

		if err := us.userRepo.UpdateTicketQuota(ctx, nil, transaction.Ticket.ID.String(), -len(transaction.TicketForms)); err != nil {
			return dto.ErrUpdateTicketQuota
		}

		transaction.TransactionStatus = req.TransactionStatus

	default:
		return dto.ErrUnknownTransactionStatus
	}

	return us.userRepo.UpdateTransactionTicket(ctx, nil, transaction)
}

// History Transactions
func (us *UserService) GetAllTransactions(ctx context.Context) ([]dto.TransactionResponse, error) {
	token := ctx.Value("Authorization").(string)
	userIDStr, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return nil, err
	}

	transactions, err := us.userRepo.GetAllTransactions(ctx, nil, userIDStr)
	if err != nil {
		return nil, err
	}

	var datas []dto.TransactionResponse
	for _, transaction := range transactions {
		data := dto.TransactionResponse{
			ID:                transaction.ID,
			OrderID:           transaction.OrderID,
			ItemType:          transaction.ItemType,
			TicketName:        transaction.Ticket.Name,
			TicketType:        transaction.Ticket.Type,
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

		for _, tf := range transaction.TicketForms {
			transactionItem := dto.TicketFormResponse{
				ID:           tf.ID,
				AudienceType: tf.AudienceType,
				Instansi:     tf.Instansi,
				Email:        tf.Email,
				FullName:     tf.FullName,
				PhoneNumber:  tf.PhoneNumber,
				LineID:       tf.LineID,
			}

			data.TicketForms = append(data.TicketForms, transactionItem)
		}

		datas = append(datas, data)
	}

	return datas, nil
}
func (us *UserService) GetDetailTransactions(ctx context.Context, id string) (dto.TransactionResponse, error) {
	transaction, _, err := us.userRepo.GetTransactionByID(ctx, nil, id)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	data := dto.TransactionResponse{
		ID:                transaction.ID,
		OrderID:           transaction.OrderID,
		ItemType:          transaction.ItemType,
		TicketName:        transaction.Ticket.Name,
		TicketType:        transaction.Ticket.Type,
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
	for _, tf := range transaction.TicketForms {
		transactionItem := dto.TicketFormResponse{
			ID:           tf.ID,
			AudienceType: tf.AudienceType,
			Instansi:     tf.Instansi,
			Email:        tf.Email,
			FullName:     tf.FullName,
			PhoneNumber:  tf.PhoneNumber,
			LineID:       tf.LineID,
			QRCodeURL:    fmt.Sprintf("assets/qrcodes/%s", tf.ID),
		}
		data.TicketForms = append(data.TicketForms, transactionItem)
	}

	return data, nil
}
