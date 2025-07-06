package entity

import "github.com/google/uuid"

type BundleItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"bundle_item_id"`

	MerchID  *uuid.UUID `gorm:"type:uuid" json:"merch_id"`
	Merch    Merch      `gorm:"foreignKey:MerchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TicketID *uuid.UUID `gorm:"type:uuid" json:"ticket_id"`
	Ticket   Ticket     `gorm:"foreignKey:TicketID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BundleID *uuid.UUID `gorm:"type:uuid" json:"bundle_id"`
	Bundle   Bundle     `gorm:"foreignKey:BundleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
