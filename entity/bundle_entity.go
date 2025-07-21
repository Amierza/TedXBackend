package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bundle struct {
	ID    uuid.UUID  `gorm:"type:uuid;primaryKey" json:"bundle_id"`
	Name  string     `gorm:"not null" json:"bundle_name"`
	Image string     `gorm:"not null" json:"bundle_image"`
	Type  BundleType `gorm:"not null" json:"bundle_type"`
	Price float64    `gorm:"not null;default:0" json:"bundle_price"`
	Quota int        `gorm:"not null;default:0" json:"bundle_quota"`

	BundleItems  []BundleItem  `gorm:"foreignKey:BundleID"`
	Transactions []Transaction `gorm:"foreignKey:BundleID"`

	TimeStamp
}

func (b *Bundle) BeforeCreate(tx *gorm.DB) error {
	if !IsValidBundleType(b.Type) {
		return errors.New("invalid item type")
	}

	return nil
}
