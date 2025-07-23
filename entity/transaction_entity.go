package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                uuid.UUID  `gorm:"type:uuid;primaryKey" json:"transaction_id"`
	OrderID           string     `json:"order_id"`
	ItemType          ItemType   `json:"item_type"`
	ReferalCode       string     `json:"referal_code"`
	TransactionStatus string     `json:"transaction_status"`
	PaymentType       string     `json:"payment_type"`
	SignatureKey      string     `json:"signature_key"`
	Acquire           string     `json:"acquire"`
	SettlementTime    *time.Time `json:"settlement_time"`
	GrossAmount       float64    `json:"gross_amount"`

	UserID *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TicketID *uuid.UUID `gorm:"type:uuid" json:"ticket_id"`
	Ticket   Ticket     `gorm:"foreignKey:TicketID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	BundleID *uuid.UUID `gorm:"type:uuid" json:"bundle_id"`
	Bundle   Bundle     `gorm:"foreignKey:BundleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TicketForms []TicketForm `gorm:"foreignKey:TransactionID"`

	TimeStamp
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
		}
	}()

	if !IsValidItemType(t.ItemType) {
		return errors.New("invalid item type")
	}
	return nil
}
