package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sponsorship struct {
	ID                  uuid.UUID           `gorm:"type:uuid;primaryKey" json:"sponsorship_id"`
	SponsorshipCategory SponsorshipCategory `gorm:"not null" json:"sponsorship_cat"`
	Name                string              `gorm:"not null" json:"sponsorship_name"`

	TimeStamp
}

func (s *Sponsorship) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
		}
	}()

	if !IsValidSponsorshipCategory(s.SponsorshipCategory) {
		return errors.New("invalid sponsorship category")
	}

	return nil
}
