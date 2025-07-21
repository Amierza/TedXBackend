package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"session_id"`
	SessionToken string     `gorm:"not null" json:"session_name"`
	ExpiredAt    *time.Time `json:"expired_at"`

	UserID *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
