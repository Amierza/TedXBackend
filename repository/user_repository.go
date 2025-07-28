package repository

import (
	"context"
	"errors"

	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		RunInTransaction(ctx context.Context, fn func(txRepo IUserRepository) error) error

		// CREATE / POST
		CreateTransaction(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) error
		CreateTicketForm(ctx context.Context, tx *gorm.DB, ticketForm entity.TicketForm) error

		// READ / GET
		GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, bool, error)
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		GetAllTicket(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, error)
		GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error)
		GetAllSpeaker(ctx context.Context, tx *gorm.DB) ([]entity.Speaker, error)
		GetAllMerch(ctx context.Context, tx *gorm.DB) ([]entity.Merch, error)
		GetAllBundle(ctx context.Context, tx *gorm.DB, bundleType string) ([]entity.Bundle, error)
		GetTicketByID(ctx context.Context, tx *gorm.DB, ticketID string) (entity.Ticket, bool, error)
		GetBundleByID(ctx context.Context, tx *gorm.DB, bundleID string) (entity.Bundle, bool, error)
		GetTransactionByOrderID(ctx context.Context, tx *gorm.DB, orderID string) (entity.Transaction, bool, error)
		GetStudentAmbassadorByReferalCode(ctx context.Context, tx *gorm.DB, referalCode string) (entity.StudentAmbassador, bool, error)

		// UPDATE / PATCH
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) error
		UpdateBundleQuota(ctx context.Context, tx *gorm.DB, bundleID string, newQuota int) error
		UpdateTicketQuota(ctx context.Context, tx *gorm.DB, ticketID string, newQuota int) error
		UpdateTransactionTicket(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) error
		UpdateMaxReferal(ctx context.Context, tx *gorm.DB, saID string, maxReferal int) error

		// DELETE / DELETE
	}

	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) RunInTransaction(ctx context.Context, fn func(txRepo IUserRepository) error) error {
	return ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &UserRepository{db: tx}
		return fn(txRepo)
	})
}

// CREATE / POST
func (ur *UserRepository) CreateTransaction(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Create(&transaction).Error
}
func (ur *UserRepository) CreateTicketForm(ctx context.Context, tx *gorm.DB, ticketForm entity.TicketForm) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Create(&ticketForm).Error
}

// READ / GET
func (ur *UserRepository) GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ur *UserRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ur *UserRepository) GetAllTicket(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, error) {
	if tx == nil {
		tx = ur.db
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
func (ur *UserRepository) GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error) {
	if tx == nil {
		tx = ur.db
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
func (ur *UserRepository) GetAllSpeaker(ctx context.Context, tx *gorm.DB) ([]entity.Speaker, error) {
	if tx == nil {
		tx = ur.db
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
func (ur *UserRepository) GetAllMerch(ctx context.Context, tx *gorm.DB) ([]entity.Merch, error) {
	if tx == nil {
		tx = ur.db
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
func (ur *UserRepository) GetAllBundle(ctx context.Context, tx *gorm.DB, bundleType string) ([]entity.Bundle, error) {
	if tx == nil {
		tx = ur.db
	}

	var (
		bundles []entity.Bundle
		err     error
	)

	query := tx.WithContext(ctx).Model(&entity.Bundle{}).Preload("BundleItems.Merch")

	if bundleType != "" {
		query = query.Where("type = ?", bundleType)
	}

	if err := query.Order(`"createdAt" DESC`).Find(&bundles).Error; err != nil {
		return []entity.Bundle{}, err
	}

	return bundles, err
}
func (ur *UserRepository) GetTicketByID(ctx context.Context, tx *gorm.DB, ticketID string) (entity.Ticket, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var ticket entity.Ticket
	if err := tx.WithContext(ctx).Where("id = ?", ticketID).Take(&ticket).Error; err != nil {
		return entity.Ticket{}, false, err
	}

	return ticket, true, nil
}
func (ur *UserRepository) GetBundleByID(ctx context.Context, tx *gorm.DB, bundleID string) (entity.Bundle, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var bundle entity.Bundle
	if err := tx.WithContext(ctx).Preload("BundleItems.Merch").Where("id = ?", bundleID).Take(&bundle).Error; err != nil {
		return entity.Bundle{}, false, err
	}

	return bundle, true, nil
}
func (ur *UserRepository) GetTransactionByOrderID(ctx context.Context, tx *gorm.DB, orderID string) (entity.Transaction, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var transaction entity.Transaction
	if err := tx.WithContext(ctx).Where("order_id = ?", orderID).Take(&transaction).Error; err != nil {
		return entity.Transaction{}, false, err
	}

	return transaction, true, nil
}
func (ur *UserRepository) GetStudentAmbassadorByReferalCode(ctx context.Context, tx *gorm.DB, referalCode string) (entity.StudentAmbassador, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var studentAmbassador entity.StudentAmbassador
	if err := tx.WithContext(ctx).Where("referal_code = ?", referalCode).Take(&studentAmbassador).Error; err != nil {
		return entity.StudentAmbassador{}, false, err
	}

	return studentAmbassador, true, nil
}

// UPDATE / PATCH
func (ur *UserRepository) UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Where("id = ?", user.ID).Updates(&user).Error
}
func (ur *UserRepository) UpdateBundleQuota(ctx context.Context, tx *gorm.DB, bundleID string, newQuota int) error {
	if tx == nil {
		tx = ur.db
	}

	result := tx.WithContext(ctx).
		Model(&entity.Bundle{}).
		Where("id = ?", bundleID).
		Update("quota", newQuota)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("bundle not found or no change made")
	}

	return nil
}
func (ur *UserRepository) UpdateTicketQuota(ctx context.Context, tx *gorm.DB, ticketID string, newQuota int) error {
	if tx == nil {
		tx = ur.db
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
func (ur *UserRepository) UpdateTransactionTicket(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Where("id = ?", transaction.ID).Updates(&transaction).Error
}
func (ur *UserRepository) UpdateMaxReferal(ctx context.Context, tx *gorm.DB, saID string, maxReferal int) error {
	if tx == nil {
		tx = ur.db
	}

	result := tx.WithContext(ctx).
		Model(&entity.StudentAmbassador{}).
		Where("id = ?", saID).
		Update("max_referal", maxReferal)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("student ambassador not found or no change made")
	}

	return nil
}

// DELETE / DELETE
