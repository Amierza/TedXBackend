package repository

import (
	"context"
	"math"
	"strings"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

type (
	IAdminRepository interface {
		// CREATE / POST
		CreateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error
		CreateSpeaker(ctx context.Context, tx *gorm.DB, speaker entity.Speaker) error

		// READ / GET
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		GetAllSponsorship(ctx context.Context, tx *gorm.DB) ([]entity.Sponsorship, error)
		GetSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) (entity.Sponsorship, bool, error)
		GetSponsorshipByNameAndCategory(ctx context.Context, tx *gorm.DB, name string, category string) (entity.Sponsorship, bool, error)
		GetSpeakerByID(ctx context.Context, tx *gorm.DB, speakerID string) (entity.Speaker, bool, error)
		GetSpeakerByName(ctx context.Context, tx *gorm.DB, name string) (entity.Speaker, bool, error)
		GetAllSpeaker(ctx context.Context, tx *gorm.DB) ([]entity.Speaker, error)
		GetAllSpeakerWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.SpeakerPaginationRepositoryResponse, error)

		// UPDATE / PATCH
		UpdateSponsorship(ctx context.Context, tx *gorm.DB, sponsorship entity.Sponsorship) error
		UpdateSpeaker(ctx context.Context, tx *gorm.DB, speaker entity.Speaker) error

		// DELETE / DELETE
		DeleteSponsorshipByID(ctx context.Context, tx *gorm.DB, sponsorshipID string) error
		DeleteSpeakerByID(ctx context.Context, tx *gorm.DB, speakerID string) error
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
func (ar *AdminRepository) GetSponsorshipByNameAndCategory(ctx context.Context, tx *gorm.DB, name string, category string) (entity.Sponsorship, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var sponsorship entity.Sponsorship
	if err := tx.WithContext(ctx).Where("name = ?", name).Where("category = ?", category).Take(&sponsorship).Error; err != nil {
		return entity.Sponsorship{}, false, err
	}

	return sponsorship, false, nil
}
func (ar *AdminRepository) GetSpeakerByID(ctx context.Context, tx *gorm.DB, speakerID string) (entity.Speaker, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var speaker entity.Speaker
	if err := tx.WithContext(ctx).Where("id = ?", speakerID).Take(&speaker).Error; err != nil {
		return entity.Speaker{}, false, err
	}

	return speaker, false, nil
}
func (ar *AdminRepository) GetSpeakerByName(ctx context.Context, tx *gorm.DB, name string) (entity.Speaker, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var speaker entity.Speaker
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&speaker).Error; err != nil {
		return entity.Speaker{}, false, err
	}

	return speaker, false, nil
}
func (ar *AdminRepository) GetAllSpeaker(ctx context.Context, tx *gorm.DB) ([]entity.Speaker, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		speakers []entity.Speaker
		err      error
	)

	if err := tx.WithContext(ctx).Model(&entity.Sponsorship{}).Find(&speakers).Error; err != nil {
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

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&speakers).Error; err != nil {
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

// UPDATE / PATCH
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

// DELETE / DELETE
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
