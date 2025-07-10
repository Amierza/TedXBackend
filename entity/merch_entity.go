package entity

import "github.com/google/uuid"

type Merch struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey" json:"merch_id"`
	Name          string        `gorm:"not null" json:"merch_name"`
	Stock         int           `gorm:"not null;default:0" json:"merch_stock"`
	Description   string        `gorm:"not null" json:"merch_desc"`
	MerchCategory MerchCategory `gorm:"not null;default:'t-shirt'" json:"merch_category"`

	BundleItems       []BundleItem       `gorm:"foreignKey:MerchID"`
	MerchColorDetails []MerchColorDetail `gorm:"foreignKey:MerchID"`
	MerchImageDetails []MerchImageDetail `gorm:"foreignKey:MerchID"`
	MerchSizeDetails  []MerchSizeDetail  `gorm:"foreignKey:MerchID"`

	TimeStamp
}
