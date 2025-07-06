package entity

import "github.com/google/uuid"

type MerchColorDetail struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"merch_color_detail_id"`

	MerchID      *uuid.UUID `gorm:"type:uuid" json:"merch_id"`
	Merch        Merch      `gorm:"foreignKey:MerchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MerchColorID *uuid.UUID `gorm:"type:uuid" json:"merch_color_id"`
	MerchColor   MerchColor `gorm:"foreignKey:MerchColorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
