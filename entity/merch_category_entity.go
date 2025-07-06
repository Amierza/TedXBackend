package entity

import "github.com/google/uuid"

type MerchCategory struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"merch_cat_id"`
	Name string    `gorm:"not null" json:"merch_cat_name"`

	Merchs []Merch `gorm:"foreignKey:MerchCategoryID"`

	TimeStamp
}
