package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Merch struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey" json:"merch_id"`
	Name        string        `gorm:"not null" json:"merch_name"`
	Stock       int           `gorm:"not null;default:0" json:"merch_stock"`
	Price       float64       `gorm:"not null;default:0" json:"merch_price"`
	Description string        `gorm:"not null" json:"merch_desc"`
	Category    MerchCategory `gorm:"not null;default:'t-shirt'" json:"merch_cat"`

	BundleItems []BundleItem `gorm:"foreignKey:MerchID"`
	MerchImages []MerchImage `gorm:"foreignKey:MerchID"`

	TimeStamp
}

func (m *Merch) BeforeCreate(tx *gorm.DB) error {
	if !IsValidMerchCategory(m.Category) {
		return errors.New("invalid item type")
	}

	return nil
}
