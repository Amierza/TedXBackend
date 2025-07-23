package repository

import (
	"context"

	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		RunInTransaction(ctx context.Context, fn func(txRepo IUserRepository) error) error

		// CREATE / POST

		// READ / GET
		GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, bool, error)
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		GetAllTicket(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, error)
		GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error)
		GetAllSpeaker(ctx context.Context, tx *gorm.DB) ([]entity.Speaker, error)
		GetAllMerch(ctx context.Context, tx *gorm.DB) ([]entity.Merch, error)
		// GetAllBundle(ctx context.Context, tx *gorm.DB, bundleType string) ([]entity.Bundle, error)

		// UPDATE / PATCH

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

// func (ur *UserRepository) GetAllBundle(ctx context.Context, tx *gorm.DB, bundleType string) ([]entity.Bundle, error) {
// 	if tx == nil {
// 		tx = ur.db
// 	}

// 	var (
// 		bundles []entity.Bundle
// 		err     error
// 	)

// 	query := tx.WithContext(ctx).Model(&entity.Bundle{}).Preload("BundleItems.Merch")

// 	if bundleType != "" {
// 		query = query.Where("type = ?", bundleType)
// 	}

// 	if err := query.Order(`"createdAt" DESC`).Find(&bundles).Error; err != nil {
// 		return []entity.Bundle{}, err
// 	}

// 	return bundles, err
// }

// UPDATE / PATCH

// DELETE / DELETE
