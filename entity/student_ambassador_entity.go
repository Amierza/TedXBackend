package entity

import "github.com/google/uuid"

type StudentAmbassador struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	ReferalCode string    `gorm:"not null" json:"referal_code"`
	Discount    float64   `gorm:"type:numeric(10,2)" json:"discount"`
	MaxReferal  int       `json:"max_referal"`

	TimeStamp
}
