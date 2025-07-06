package entity

import "github.com/google/uuid"

type MerchImage struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"merch_image_id"`
	Name string    `gorm:"not null" json:"merch_image_name"`

	MerchImageDetails []MerchImageDetail `gorm:"foreignKey:MerchImageID"`

	TimeStamp
}
