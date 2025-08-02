package entity

import "github.com/google/uuid"

type Speaker struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Image       string    `gorm:"not null" json:"image"`
	Description string    `gorm:"not null" json:"desc"`

	TimeStamp
}
