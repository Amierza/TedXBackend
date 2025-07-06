package entity

import "github.com/google/uuid"

type MerchImageDetail struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"merch_image_detail_id"`

	MerchID      *uuid.UUID `gorm:"type:uuid" json:"merch_id"`
	Merch        Merch      `gorm:"foreignKey:MerchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MerchImageID *uuid.UUID `gorm:"type:uuid" json:"merch_image_id"`
	MerchImage   MerchImage `gorm:"foreignKey:MerchImageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
