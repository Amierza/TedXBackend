package repository

import (
	"context"

	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

type (
	IAdminRepository interface {
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
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
