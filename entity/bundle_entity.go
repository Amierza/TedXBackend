package entity

import "github.com/google/uuid"

type Bundle struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"bundle_id"`
	Name  string    `gorm:"not null" json:"bundle_name"`
	Price float64   `gorm:"not null" json:"bundle_price"`

	BundleItems []BundleItem `gorm:"foreignKey:TicketID"`

	TimeStamp
}
