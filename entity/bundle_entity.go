package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bundle struct {
	ID    uuid.UUID  `gorm:"type:uuid;primaryKey" json:"bundle_id"`
	Type  BundleType `gorm:"not null" json:"bundle_type"`
	Name  string     `gorm:"not null" json:"bundle_name"`
	Price float64    `gorm:"not null" json:"bundle_price"`

	BundleItems []BundleItem `gorm:"foreignKey:TicketID"`

	TimeStamp
}

func (b *Bundle) BeforeCreate(tx *gorm.DB) error {
	if !IsValidBundleType(b.Type) {
		return errors.New("invalid item type")
	}

	return nil
}
