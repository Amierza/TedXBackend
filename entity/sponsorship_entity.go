package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sponsorship struct {
	ID       uuid.UUID           `gorm:"type:uuid;primaryKey" json:"id"`
	Category SponsorshipCategory `gorm:"not null" json:"cat"`
	Name     string              `gorm:"not null" json:"name"`
	Image    string              `gorm:"not null" json:"image"`

	TimeStamp
}

func (s *Sponsorship) BeforeCreate(tx *gorm.DB) error {
	if !IsValidSponsorshipCategory(s.Category) {
		return errors.New("invalid sponsorship category")
	}

	return nil
}
