package entity

import "github.com/google/uuid"

type MerchImage struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name"`

	MerchID *uuid.UUID `gorm:"type:uuid" json:"merch_id"`
	Merch   Merch      `gorm:"foreignKey:MerchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
