package repository

import (
	"context"

	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

type (
	IAdminRepository interface {
		// CREATE / POST
		CreateSponsorship(ctx context.Context, tx *gorm.DB, spon entity.Sponsorship) error

		// READ / GET
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error)
		GetSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) (entity.Sponsorship, bool, error)

		// UPDATE / PATCH
		UpdateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error

		// DELETE / DELETE
		DeleteSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) error
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

// CREATE / POST
func (ar *AdminRepository) CreateSponsorship(ctx context.Context, tx *gorm.DB, spon entity.Sponsorship) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&spon).Error
}

// READ / GET
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
func (ar *AdminRepository) GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		sponsorships []entity.Sponsorship
		err          error
	)

	if err := tx.WithContext(ctx).Model(&entity.Sponsorship{}).Find(&sponsorships).Error; err != nil {
		return []entity.Sponsorship{}, err
	}

	return sponsorships, err
}
func (ar *AdminRepository) GetSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) (entity.Sponsorship, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var sponsorship entity.Sponsorship
	if err := tx.WithContext(ctx).Where("id = ?", sponsorshipID).Take(&sponsorship).Error; err != nil {
		return entity.Sponsorship{}, false, err
	}

	return sponsorship, false, nil
}

// UPDATE / PATCH
func (ar *AdminRepository) UpdateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", sponsorship.ID).Updates(&sponsorship).Error
}

// DELETE / DELETE
func (ar *AdminRepository) DeleteSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", sponsorshipID).Delete(&entity.Sponsorship{}).Error
}
