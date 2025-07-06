package entity

import "github.com/google/uuid"

type Ticket struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"ticket_id"`
	Name  string    `gorm:"not null" json:"ticket_name"`
	Price float64   `gorm:"not null" json:"ticket_price"`

	BundleItems []BundleItem `gorm:"foreignKey:TicketID"`

	TimeStamp
}
