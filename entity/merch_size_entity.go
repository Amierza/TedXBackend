package entity

import "github.com/google/uuid"

type MerchSize struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"merch_size_id"`
	Name string    `gorm:"not null" json:"merch_size_name"`

	MerchSizeDetails []MerchSizeDetail `gorm:"foreignKey:MerchSizeID"`

	TimeStamp
}
