package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID  `gorm:"type:uuid;primaryKey" json:"account_id"`
	Type              string     `gorm:"not null" json:"account_type"`
	Provider          string     `gorm:"not null" json:"account_provider"`
	ProviderAccountID string     `gorm:"not null" json:"provider_account_id"`
	RefreshToken      string     `json:"refresh_token"`
	AccessToken       string     `json:"access_token"`
	ExpiredAt         *time.Time `json:"expired_at"`
	TokenType         string     `json:"token_type"`
	Scope             string     `json:"scope"`
	TokenID           string     `json:"token_id"`
	SessionState      string     `json:"session_state"`

	UserID *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
