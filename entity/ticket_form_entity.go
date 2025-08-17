package entity

import (
	"errors"

	"github.com/Amierza/TedXBackend/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketForm struct {
	ID           uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`
	AudienceType AudienceType `gorm:"default:'regular'" json:"audience_type"`
	Instansi     Instansi     `gorm:"default:'unair'" json:"instansi"`
	Email        string       `gorm:"not null" json:"email"`
	FullName     string       `gorm:"not null" json:"full_name"`
	PhoneNumber  string       `gorm:"not null" json:"phone_number"`
	LineID       string       `json:"line_id"`

	GuestAttendances []GuestAttendance `gorm:"foreignKey:TicketFormID"`

	TransactionID *uuid.UUID  `gorm:"type:uuid" json:"transaction_id"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}

func (tf *TicketForm) BeforeCreate(tx *gorm.DB) error {
	var err error

	tf.PhoneNumber, err = helpers.StandardizePhoneNumber(tf.PhoneNumber)
	if err != nil {
		return err
	}

	if !IsValidAudienceType(tf.AudienceType) {
		return errors.New("invalid audience type")
	}

	return nil
}
