package entity

import "github.com/google/uuid"

type MerchColor struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"merch_color_id"`
	Name string    `gorm:"not null" json:"merch_color_name"`
	Code string    `gorm:"not null" json:"merch_color_code"`

	MerchColorDetails []MerchColorDetail `gorm:"foreignKey:MerchColorID"`

	TimeStamp
}
