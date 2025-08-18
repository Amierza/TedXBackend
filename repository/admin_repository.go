package repository

import (
	"context"
	"errors"
	"math"
	"strings"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

type (
	IAdminRepository interface {
		RunInTransaction(ctx context.Context, fn func(txRepo IAdminRepository) error) error

		// CREATE / POST
		CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) error
		CreateTicket(ctx context.Context, tx *gorm.DB, ticket entity.Ticket) error
		CreateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error
		CreateSpeaker(ctx context.Context, tx *gorm.DB, speaker entity.Speaker) error
		CreateMerch(ctx context.Context, tx *gorm.DB, merch entity.Merch) error
		CreateMerchImage(ctx context.Context, tx *gorm.DB, image entity.MerchImage) error
		CreateBundle(ctx context.Context, tx *gorm.DB, bundle entity.Bundle) error
		CreateBundleItem(ctx context.Context, tx *gorm.DB, bundleItem entity.BundleItem) error
		CreateTransaction(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) error
		CreateTicketForm(ctx context.Context, tx *gorm.DB, ticketForm entity.TicketForm) error
		CreateStudentAmbassador(ctx context.Context, tx *gorm.DB, studentAmbassador entity.StudentAmbassador) error
		CreateGuestAttendance(ctx context.Context, tx *gorm.DB, guestAttendance entity.GuestAttendance) error

		// READ / GET
		GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, bool, error)
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		GetAllUser(ctx context.Context, tx *gorm.DB, roleName string) ([]entity.User, error)
		GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, roleName string) (dto.UserPaginationRepositoryResponse, error)
		GetTicketByID(ctx context.Context, tx *gorm.DB, ticketID string) (entity.Ticket, bool, error)
		GetTicketByName(ctx context.Context, tx *gorm.DB, ticketName string) (entity.Ticket, bool, error)
		GetAllTicket(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, error)
		GetAllTicketWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.TicketPaginationRepositoryResponse, error)
		GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error)
		GetAllSponsorshipWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.SponsorshipPaginationRepositoryResponse, error)
		GetSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) (entity.Sponsorship, bool, error)
		GetSponsorshipByNameAndCategory(ctx context.Context, tx *gorm.DB, name string, category string) (entity.Sponsorship, bool, error)
		GetSpeakerByID(ctx context.Context, tx *gorm.DB, speakerID string) (entity.Speaker, bool, error)
		GetSpeakerByName(ctx context.Context, tx *gorm.DB, name string) (entity.Speaker, bool, error)
		GetAllSpeaker(ctx context.Context, tx *gorm.DB) ([]entity.Speaker, error)
		GetAllSpeakerWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.SpeakerPaginationRepositoryResponse, error)
		GetAllMerch(ctx context.Context, tx *gorm.DB) ([]entity.Merch, error)
		GetAllMerchWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.MerchPaginationRepositoryResponse, error)
		GetMerchByID(ctx context.Context, tx *gorm.DB, merchID string) (entity.Merch, bool, error)
		GetMerchImageByID(ctx context.Context, tx *gorm.DB, merchImageID string) (entity.MerchImage, bool, error)
		GetMerchImagesByMerchID(ctx context.Context, tx *gorm.DB, merchID string) ([]entity.MerchImage, error)
		GetBundleByName(ctx context.Context, tx *gorm.DB, name string) (entity.Bundle, bool, error)
		GetAllBundle(ctx context.Context, tx *gorm.DB) ([]entity.Bundle, error)
		GetAllBundleWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.BundlePaginationRepositoryResponse, error)
		GetBundleByID(ctx context.Context, tx *gorm.DB, bundleID string) (entity.Bundle, bool, error)
		GetBundleItemsByBundleID(ctx context.Context, tx *gorm.DB, bundleID string) ([]entity.BundleItem, error)
		GetAllTransaction(ctx context.Context, tx *gorm.DB, transactionStatus, ticketCategory string) ([]entity.Transaction, error)
		GetAllTransactionWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, transactionStatus, ticketCategory string) (dto.TransactionTicketPaginationRepositoryResponse, error)
		GetTransactionByID(ctx context.Context, tx *gorm.DB, transactionID string) (entity.Transaction, bool, error)
		GetStudentAmbassadorByReferalCode(ctx context.Context, tx *gorm.DB, studentAmbassadorReferalCode string) (entity.StudentAmbassador, bool, error)
		GetAllStudentAmbassador(ctx context.Context, tx *gorm.DB) ([]entity.StudentAmbassador, error)
		GetAllStudentAmbassadorWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.StudentAmbassadorPaginationRepositoryResponse, error)
		GetStudentAmbassadorByID(ctx context.Context, tx *gorm.DB, studentAmbassadorID string) (entity.StudentAmbassador, bool, error)
		GetTicketFormByID(ctx context.Context, tx *gorm.DB, ticketFormID string) (entity.TicketForm, bool, error)
		GetAllTicketForm(ctx context.Context, tx *gorm.DB) ([]entity.TicketForm, error)
		GetAllTicketFormWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.TicketFormPaginationRepositoryResponse, error)

		// UPDATE / PATCH
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) error
		UpdateTicket(ctx context.Context, tx *gorm.DB, ticket entity.Ticket) error
		UpdateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error
		UpdateSpeaker(ctx context.Context, tx *gorm.DB, speaker entity.Speaker) error
		UpdateMerch(ctx context.Context, tx *gorm.DB, merch entity.Merch) error
		UpdateBundle(ctx context.Context, tx *gorm.DB, bundle entity.Bundle) error
		UpdateTicketQuota(ctx context.Context, tx *gorm.DB, ticketID string, newQuota int) error
		UpdateStudentAmbassador(ctx context.Context, tx *gorm.DB, studentAmbassador entity.StudentAmbassador) error

		// DELETE / DELETE
		DeleteUserByID(ctx context.Context, tx *gorm.DB, userID string) error
		DeleteTicketByID(ctx context.Context, tx *gorm.DB, ticketID string) error
		DeleteSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) error
		DeleteSpeakerByID(ctx context.Context, tx *gorm.DB, speakerID string) error
		DeleteMerchImageByID(ctx context.Context, tx *gorm.DB, merchImageID string) error
		DeleteMerchByID(ctx context.Context, tx *gorm.DB, merchID string) error
		DeleteMerchImagesByMerchID(ctx context.Context, tx *gorm.DB, merchID string) error
		DeleteBundleByID(ctx context.Context, tx *gorm.DB, bundleID string) error
		DeleteBundleItemsByBundleID(ctx context.Context, tx *gorm.DB, bundleID string) error
		DeleteStudentAmbassadorByID(ctx context.Context, tx *gorm.DB, studentAmbassadorID string) error
	}

	AdminRepository struct {
		db *gorm.DB
	}
)

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (ar *AdminRepository) RunInTransaction(ctx context.Context, fn func(txRepo IAdminRepository) error) error {
	return ar.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &AdminRepository{db: tx}
		return fn(txRepo)
	})
}

// CREATE / POST
func (ar *AdminRepository) CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&user).Error
}
func (ar *AdminRepository) CreateTicket(ctx context.Context, tx *gorm.DB, ticket entity.Ticket) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&ticket).Error
}
func (ar *AdminRepository) CreateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&sponsorship).Error
}
func (ar *AdminRepository) CreateSpeaker(ctx context.Context, tx *gorm.DB, speaker entity.Speaker) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&speaker).Error
}
func (ar *AdminRepository) CreateMerch(ctx context.Context, tx *gorm.DB, merch entity.Merch) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&merch).Error
}
func (ar *AdminRepository) CreateMerchImage(ctx context.Context, tx *gorm.DB, image entity.MerchImage) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&image).Error
}
func (ar *AdminRepository) CreateBundle(ctx context.Context, tx *gorm.DB, bundle entity.Bundle) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&bundle).Error
}
func (ar *AdminRepository) CreateBundleItem(ctx context.Context, tx *gorm.DB, bundleItem entity.BundleItem) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&bundleItem).Error
}
func (ar *AdminRepository) CreateTransaction(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&transaction).Error
}
func (ar *AdminRepository) CreateTicketForm(ctx context.Context, tx *gorm.DB, ticketForm entity.TicketForm) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&ticketForm).Error
}
func (ar *AdminRepository) CreateStudentAmbassador(ctx context.Context, tx *gorm.DB, studentAmbassador entity.StudentAmbassador) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&studentAmbassador).Error
}
func (ar *AdminRepository) CreateGuestAttendance(ctx context.Context, tx *gorm.DB, guestAttendance entity.GuestAttendance) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&guestAttendance).Error
}

// READ / GET
func (ar *AdminRepository) GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ar *AdminRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ar *AdminRepository) GetAllUser(ctx context.Context, tx *gorm.DB, roleName string) ([]entity.User, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		users []entity.User
		err   error
	)

	query := tx.WithContext(ctx).Model(&entity.User{})

	if roleName != "" {
		query = query.Where("role = ?", roleName)
	}

	if err := query.Order(`"createdAt" DESC`).Find(&users).Error; err != nil {
		return []entity.User{}, err
	}

	return users, err
}
func (ar *AdminRepository) GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, roleName string) (dto.UserPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var users []entity.User
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.User{})

	if roleName != "" {
		query = query.Where("role = ?", roleName)
	}

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.UserPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&users).Error; err != nil {
		return dto.UserPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.UserPaginationRepositoryResponse{
		Users: users,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetTicketByID(ctx context.Context, tx *gorm.DB, ticketID string) (entity.Ticket, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var ticket entity.Ticket
	if err := tx.WithContext(ctx).Where("id = ?", ticketID).Take(&ticket).Error; err != nil {
		return entity.Ticket{}, false, err
	}

	return ticket, true, nil
}
func (ar *AdminRepository) GetTicketByName(ctx context.Context, tx *gorm.DB, ticketName string) (entity.Ticket, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var ticket entity.Ticket
	if err := tx.WithContext(ctx).Where("name = ?", ticketName).Take(&ticket).Error; err != nil {
		return entity.Ticket{}, false, err
	}

	return ticket, true, nil
}
func (ar *AdminRepository) GetAllTicket(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		tickets []entity.Ticket
		err     error
	)

	query := tx.WithContext(ctx).Model(&entity.Ticket{})

	if err := query.Order(`"createdAt" DESC`).Find(&tickets).Error; err != nil {
		return []entity.Ticket{}, err
	}

	return tickets, err
}
func (ar *AdminRepository) GetAllTicketWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.TicketPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var tickets []entity.Ticket
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Ticket{})

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.TicketPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&tickets).Error; err != nil {
		return dto.TicketPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.TicketPaginationRepositoryResponse{
		Tickets: tickets,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		sponsorships []entity.Sponsorship
		err          error
	)

	query := tx.WithContext(ctx).Model(&entity.Sponsorship{})

	if err := query.Order(`"createdAt" DESC`).Find(&sponsorships).Error; err != nil {
		return []entity.Sponsorship{}, err
	}

	return sponsorships, err
}
func (ar *AdminRepository) GetAllSponsorshipWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.SponsorshipPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var sponsorships []entity.Sponsorship
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Sponsorship{})

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.SponsorshipPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&sponsorships).Error; err != nil {
		return dto.SponsorshipPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.SponsorshipPaginationRepositoryResponse{
		Sponsorships: sponsorships,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) (entity.Sponsorship, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var sponsorship entity.Sponsorship
	if err := tx.WithContext(ctx).Where("id = ?", sponsorshipID).Take(&sponsorship).Error; err != nil {
		return entity.Sponsorship{}, false, err
	}

	return sponsorship, true, nil
}
func (ar *AdminRepository) GetSponsorshipByNameAndCategory(ctx context.Context, tx *gorm.DB, name string, category string) (entity.Sponsorship, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var sponsorship entity.Sponsorship
	if err := tx.WithContext(ctx).Where("name = ?", name).Where("category = ?", category).Take(&sponsorship).Error; err != nil {
		return entity.Sponsorship{}, false, err
	}

	return sponsorship, true, nil
}
func (ar *AdminRepository) GetSpeakerByID(ctx context.Context, tx *gorm.DB, speakerID string) (entity.Speaker, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var speaker entity.Speaker
	if err := tx.WithContext(ctx).Where("id = ?", speakerID).Take(&speaker).Error; err != nil {
		return entity.Speaker{}, false, err
	}

	return speaker, true, nil
}
func (ar *AdminRepository) GetSpeakerByName(ctx context.Context, tx *gorm.DB, name string) (entity.Speaker, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var speaker entity.Speaker
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&speaker).Error; err != nil {
		return entity.Speaker{}, false, err
	}

	return speaker, true, nil
}
func (ar *AdminRepository) GetAllSpeaker(ctx context.Context, tx *gorm.DB) ([]entity.Speaker, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		speakers []entity.Speaker
		err      error
	)

	query := tx.WithContext(ctx).Model(&entity.Speaker{})

	if err := query.Order(`"createdAt" DESC`).Find(&speakers).Error; err != nil {
		return []entity.Speaker{}, err
	}

	return speakers, err
}
func (ar *AdminRepository) GetAllSpeakerWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.SpeakerPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var speakers []entity.Speaker
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Speaker{})

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.SpeakerPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&speakers).Error; err != nil {
		return dto.SpeakerPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.SpeakerPaginationRepositoryResponse{
		Speakers: speakers,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetAllMerch(ctx context.Context, tx *gorm.DB) ([]entity.Merch, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		merchs []entity.Merch
		err    error
	)

	query := tx.WithContext(ctx).Model(&entity.Merch{}).Preload("MerchImages")

	if err := query.Order(`"createdAt" DESC`).Find(&merchs).Error; err != nil {
		return []entity.Merch{}, err
	}

	return merchs, err
}
func (ar *AdminRepository) GetAllMerchWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.MerchPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var merchs []entity.Merch
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Merch{}).Preload("MerchImages")

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.MerchPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&merchs).Error; err != nil {
		return dto.MerchPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.MerchPaginationRepositoryResponse{
		Merchs: merchs,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetMerchByID(ctx context.Context, tx *gorm.DB, merchID string) (entity.Merch, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var merch entity.Merch
	if err := tx.WithContext(ctx).Preload("MerchImages").Where("id = ?", merchID).Take(&merch).Error; err != nil {
		return entity.Merch{}, false, err
	}

	return merch, true, nil
}
func (ar *AdminRepository) GetMerchImageByID(ctx context.Context, tx *gorm.DB, merchImageID string) (entity.MerchImage, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var merchImage entity.MerchImage
	if err := tx.WithContext(ctx).Preload("Merch").Where("id = ?", merchImageID).Take(&merchImage).Error; err != nil {
		return entity.MerchImage{}, false, err
	}

	return merchImage, true, nil
}
func (ar *AdminRepository) GetMerchImagesByMerchID(ctx context.Context, tx *gorm.DB, merchID string) ([]entity.MerchImage, error) {
	if tx == nil {
		tx = ar.db
	}

	var images []entity.MerchImage
	if err := tx.WithContext(ctx).Preload("Merch").Where("merch_id = ?", merchID).Find(&images).Error; err != nil {
		return []entity.MerchImage{}, err
	}

	return images, nil
}
func (ar *AdminRepository) GetBundleByName(ctx context.Context, tx *gorm.DB, name string) (entity.Bundle, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var bundle entity.Bundle
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&bundle).Error; err != nil {
		return entity.Bundle{}, false, err
	}

	return bundle, true, nil
}
func (ar *AdminRepository) GetAllBundle(ctx context.Context, tx *gorm.DB) ([]entity.Bundle, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		bundles []entity.Bundle
		err     error
	)

	query := tx.WithContext(ctx).Model(&entity.Bundle{}).Preload("BundleItems.Merch")

	if err := query.Order(`"createdAt" DESC`).Find(&bundles).Error; err != nil {
		return []entity.Bundle{}, err
	}

	return bundles, err
}
func (ar *AdminRepository) GetAllBundleWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.BundlePaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var bundles []entity.Bundle
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Bundle{}).Preload("BundleItems.Merch")

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.BundlePaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&bundles).Error; err != nil {
		return dto.BundlePaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.BundlePaginationRepositoryResponse{
		Bundles: bundles,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetBundleByID(ctx context.Context, tx *gorm.DB, bundleID string) (entity.Bundle, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var bundle entity.Bundle
	if err := tx.WithContext(ctx).Preload("BundleItems.Merch.MerchImages").Where("id = ?", bundleID).Take(&bundle).Error; err != nil {
		return entity.Bundle{}, false, err
	}

	return bundle, true, nil
}
func (ar *AdminRepository) GetBundleItemsByBundleID(ctx context.Context, tx *gorm.DB, bundleID string) ([]entity.BundleItem, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		items []entity.BundleItem
		err   error
	)

	if err := tx.WithContext(ctx).Preload("Bundle").Preload("Merch").Model(&entity.BundleItem{}).Where("bundle_id = ?", bundleID).Find(&items).Error; err != nil {
		return []entity.BundleItem{}, err
	}

	return items, err
}
func (ar *AdminRepository) GetAllTransaction(ctx context.Context, tx *gorm.DB, transactionStatus, ticketCategory string) ([]entity.Transaction, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		transactions []entity.Transaction
		err          error
	)

	query := tx.WithContext(ctx).Model(&entity.Transaction{}).Joins("JOIN ticket_forms ON transactions.id = ticket_forms.transaction_id").Group("transactions.id").Preload("TicketForms").Preload("Ticket").Preload("Bundle")

	if transactionStatus != "" {
		query = query.Where("transactions.transaction_status = ?", transactionStatus)
	}

	if ticketCategory != "" {
		query = query.Where("ticket_forms.audience_type = ?", ticketCategory)
	}

	if err := query.Order(`"createdAt" DESC`).Find(&transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}

	return transactions, err
}
func (ar *AdminRepository) GetAllTransactionWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, transactionStatus, ticketCategory string) (dto.TransactionTicketPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		transactions []entity.Transaction
		err          error
		count        int64
	)

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Transaction{}).Joins("JOIN ticket_forms ON transactions.id = ticket_forms.transaction_id").Group("transactions.id").Preload("TicketForms").Preload("Ticket").Preload("Bundle")

	if transactionStatus != "" {
		query = query.Where("transactions.transaction_status = ?", transactionStatus)
	}

	if ticketCategory != "" {
		query = query.Where("ticket_forms.audience_type = ?", ticketCategory)
	}

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(ticket_forms.full_name) LIKE ? OR LOWER(ticket_forms.email) LIKE ? OR LOWER(ticket_forms.phone_number) LIKE ?", searchValue, searchValue, searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.TransactionTicketPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&transactions).Error; err != nil {
		return dto.TransactionTicketPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.TransactionTicketPaginationRepositoryResponse{
		Transactions: transactions,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetTransactionByID(ctx context.Context, tx *gorm.DB, transactionID string) (entity.Transaction, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var transaction entity.Transaction
	if err := tx.WithContext(ctx).Preload("TicketForms").Preload("Ticket").Preload("Bundle").Where("id = ?", transactionID).Take(&transaction).Error; err != nil {
		return entity.Transaction{}, false, err
	}

	return transaction, true, nil
}
func (ar *AdminRepository) GetStudentAmbassadorByReferalCode(ctx context.Context, tx *gorm.DB, studentAmbassadorReferalCode string) (entity.StudentAmbassador, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var studentAmbassador entity.StudentAmbassador
	if err := tx.WithContext(ctx).Where("referal_code = ?", studentAmbassadorReferalCode).Take(&studentAmbassador).Error; err != nil {
		return entity.StudentAmbassador{}, false, err
	}

	return studentAmbassador, true, nil
}
func (ar *AdminRepository) GetAllStudentAmbassador(ctx context.Context, tx *gorm.DB) ([]entity.StudentAmbassador, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		studentAmbassadors []entity.StudentAmbassador
		err                error
	)

	query := tx.WithContext(ctx).Model(&entity.StudentAmbassador{})

	if err := query.Order(`"createdAt" DESC`).Find(&studentAmbassadors).Error; err != nil {
		return []entity.StudentAmbassador{}, err
	}

	return studentAmbassadors, err
}
func (ar *AdminRepository) GetAllStudentAmbassadorWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.StudentAmbassadorPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var studentAmbassadors []entity.StudentAmbassador
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.StudentAmbassador{})

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(referal_code) LIKE ?", searchValue, searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.StudentAmbassadorPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`"createdAt" DESC`).Scopes(Paginate(req.Page, req.PerPage)).Find(&studentAmbassadors).Error; err != nil {
		return dto.StudentAmbassadorPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.StudentAmbassadorPaginationRepositoryResponse{
		StudentAmbassadors: studentAmbassadors,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetStudentAmbassadorByID(ctx context.Context, tx *gorm.DB, studentAmbassadorID string) (entity.StudentAmbassador, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var studentAmbassador entity.StudentAmbassador
	if err := tx.WithContext(ctx).Where("id = ?", studentAmbassadorID).Take(&studentAmbassador).Error; err != nil {
		return entity.StudentAmbassador{}, false, err
	}

	return studentAmbassador, true, nil
}
func (ar *AdminRepository) GetTicketFormByID(ctx context.Context, tx *gorm.DB, ticketFormID string) (entity.TicketForm, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var ticketForm entity.TicketForm
	if err := tx.WithContext(ctx).Preload("GuestAttendances").Preload("Transaction.Ticket").Where("id = ?", ticketFormID).Take(&ticketForm).Error; err != nil {
		return entity.TicketForm{}, false, err
	}

	return ticketForm, true, nil
}
func (ar *AdminRepository) GetAllTicketForm(ctx context.Context, tx *gorm.DB) ([]entity.TicketForm, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		ticketForms []entity.TicketForm
		err         error
	)

	if err := tx.WithContext(ctx).
		Joins("JOIN guest_attendances ON guest_attendances.ticket_form_id = ticket_forms.id").
		Preload("GuestAttendances").
		Preload("Transaction.Ticket").
		Order(`ticket_forms."createdAt" DESC`).
		Find(&ticketForms).Error; err != nil {
		return []entity.TicketForm{}, err
	}

	return ticketForms, err
}
func (ar *AdminRepository) GetAllTicketFormWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.TicketFormPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		ticketForms []entity.TicketForm
		err         error
		count       int64
	)

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.TicketForm{}).
		Joins("JOIN guest_attendances ON guest_attendances.ticket_form_id = ticket_forms.id").
		Joins("JOIN transactions ON transactions.id = ticket_forms.transaction_id").
		Joins("JOIN tickets ON tickets.id = transactions.ticket_id").
		Preload("GuestAttendances").
		Preload("Transaction.Ticket")

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where(`
			LOWER(ticket_forms.full_name) LIKE ? OR 
			LOWER(guest_attendances.email_checker) LIKE ? OR 
			LOWER(tickets.name) LIKE ?`,
			searchValue, searchValue, searchValue,
		)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.TicketFormPaginationRepositoryResponse{}, err
	}

	if err := query.Order(`ticket_forms."createdAt" DESC`).
		Scopes(Paginate(req.Page, req.PerPage)).
		Find(&ticketForms).Error; err != nil {
		return dto.TicketFormPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.TicketFormPaginationRepositoryResponse{
		TicketForms: ticketForms,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}

// UPDATE / PATCH
func (ar *AdminRepository) UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", user.ID).Updates(&user).Error
}
func (ar *AdminRepository) UpdateTicket(ctx context.Context, tx *gorm.DB, ticket entity.Ticket) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", ticket.ID).Updates(&ticket).Error
}
func (ar *AdminRepository) UpdateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", sponsorship.ID).Updates(&sponsorship).Error
}
func (ar *AdminRepository) UpdateSpeaker(ctx context.Context, tx *gorm.DB, speaker entity.Speaker) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", speaker.ID).Updates(&speaker).Error
}
func (ar *AdminRepository) UpdateMerch(ctx context.Context, tx *gorm.DB, merch entity.Merch) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", merch.ID).Save(&merch).Error
}
func (ar *AdminRepository) UpdateBundle(ctx context.Context, tx *gorm.DB, bundle entity.Bundle) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", bundle.ID).Save(&bundle).Error
}
func (ar *AdminRepository) UpdateTicketQuota(ctx context.Context, tx *gorm.DB, ticketID string, newQuota int) error {
	if tx == nil {
		tx = ar.db
	}

	result := tx.WithContext(ctx).
		Model(&entity.Ticket{}).
		Where("id = ?", ticketID).
		Update("quota", newQuota)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("ticket not found or no change made")
	}

	return nil
}
func (ar *AdminRepository) UpdateStudentAmbassador(ctx context.Context, tx *gorm.DB, studentAmbassador entity.StudentAmbassador) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", studentAmbassador.ID).Save(&studentAmbassador).Error
}

// DELETE / DELETE
func (ar *AdminRepository) DeleteUserByID(ctx context.Context, tx *gorm.DB, userID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", userID).Delete(&entity.User{}).Error
}
func (ar *AdminRepository) DeleteTicketByID(ctx context.Context, tx *gorm.DB, ticketID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", ticketID).Delete(&entity.Ticket{}).Error
}
func (ar *AdminRepository) DeleteSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", sponsorshipID).Delete(&entity.Sponsorship{}).Error
}
func (ar *AdminRepository) DeleteSpeakerByID(ctx context.Context, tx *gorm.DB, speakerID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", speakerID).Delete(&entity.Speaker{}).Error
}
func (ar *AdminRepository) DeleteMerchImageByID(ctx context.Context, tx *gorm.DB, merchImageID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", merchImageID).Delete(&entity.MerchImage{}).Error
}
func (ar *AdminRepository) DeleteMerchByID(ctx context.Context, tx *gorm.DB, merchID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", merchID).Delete(&entity.Merch{}).Error
}
func (ar *AdminRepository) DeleteMerchImagesByMerchID(ctx context.Context, tx *gorm.DB, merchID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("merch_id = ?", merchID).Delete(&entity.MerchImage{}).Error
}
func (ar *AdminRepository) DeleteBundleByID(ctx context.Context, tx *gorm.DB, bundleID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", bundleID).Delete(&entity.Bundle{}).Error
}
func (ar *AdminRepository) DeleteBundleItemsByBundleID(ctx context.Context, tx *gorm.DB, bundleID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("bundle_id = ?", bundleID).Delete(&entity.BundleItem{}).Error
}
func (ar *AdminRepository) DeleteStudentAmbassadorByID(ctx context.Context, tx *gorm.DB, studentAmbassadorID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", studentAmbassadorID).Delete(&entity.StudentAmbassador{}).Error
}
