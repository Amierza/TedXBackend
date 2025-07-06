package entity

import "github.com/google/uuid"

type MerchSizeDetail struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"merch_size_detail_id"`

	MerchID     *uuid.UUID `gorm:"type:uuid" json:"merch_id"`
	Merch       Merch      `gorm:"foreignKey:MerchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MerchSizeID *uuid.UUID `gorm:"type:uuid" json:"merch_size_id"`
	MerchSize   MerchSize  `gorm:"foreignKey:MerchSizeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
