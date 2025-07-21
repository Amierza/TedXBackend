package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID             uuid.UUID         `gorm:"type:uuid;primaryKey" json:"transaction_id"`
	OrderID        *uuid.UUID        `gorm:"type:uuid" json:"transaction_order_id"`
	ItemType       ItemType          `json:"transaction_item_type"`
	AudienceType   AudienceType      `gorm:"default:'regular'" json:"transaction_audience_type"`
	ReferalCode    string            `json:"transaction_referal_code"`
	Status         TransactionStatus `json:"transaction_status"`
	PaymentType    PaymentType       `json:"transaction_payment_type"`
	SignatureKey   string            `json:"transaction_signature_key"`
	Acquire        Acquire           `json:"transaction_acquire"`
	SettlementTime *time.Time        `json:"transaction_settlement_time"`
	GrossAmount    float64           `json:"transaction_gross_amount"`

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

	if !IsValidAudienceType(t.AudienceType) {
		return errors.New("invalid audience type")
	}

	if !IsValidTransactionStatus(t.Status) {
		return errors.New("invalid transaction status")
	}

	if !IsValidPaymentType(t.PaymentType) {
		return errors.New("invalid payment type")
	}

	if !IsValidItemType(t.ItemType) {
		return errors.New("invalid item type")
	}

	if t.Acquire != "" && !IsValidAcquire(t.Acquire) {
		return errors.New("invalid acquire")
	}
	return nil
}
