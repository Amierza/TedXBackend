package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	SessionToken string     `json:"token"`
	Expires      *time.Time `json:"expires"`

	UserID *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
