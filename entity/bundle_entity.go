package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bundle struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Image       string     `gorm:"not null" json:"image"`
	Type        BundleType `gorm:"not null" json:"type"`
	Price       float64    `gorm:"not null;default:0" json:"price"`
	Quota       int        `gorm:"not null;default:0" json:"quota"`
	Description string     `json:"description"`
	EventDate   time.Time  `gorm:"not null" json:"event_date"`

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
