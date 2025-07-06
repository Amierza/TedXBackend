package entity

import "github.com/google/uuid"

type GuestAttendance struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"guest_att_id"`

	TicketFormID  *uuid.UUID `gorm:"type:uuid" json:"ticket_form_id"`
	TicketForm    TicketForm `gorm:"foreignKey:TicketFormID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CheckedBy     *uuid.UUID `gorm:"type:uuid" json:"checked_by"`
	CheckedByUser User       `gorm:"foreignKey:CheckedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
