package entity

import (
	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Type              string    `gorm:"not null" json:"type"`
	Provider          string    `gorm:"not null" json:"provider"`
	ProviderAccountID string    `gorm:"not null" json:"provider_account_id"`
	RefreshToken      string    `json:"refresh_token"`
	AccessToken       string    `json:"access_token"`
	ExpiresAt         int       `json:"expires_at"`
	TokenType         string    `json:"token_type"`
	Scope             string    `json:"scope"`
	IDToken           string    `json:"id_token"`
	SessionState      string    `json:"session_state"`

	UserID *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
