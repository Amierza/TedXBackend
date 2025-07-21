package entity

import "github.com/google/uuid"

type Ticket struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"ticket_id"`
	Name  string    `gorm:"not null" json:"ticket_name"`
	Price float64   `gorm:"not null;default:0" json:"ticket_price"`
	Image string    `gorm:"not null" json:"ticket_image"`
	Quota int       `gorm:"not null;default:0" json:"ticket_quota"`

	BundleItems  []BundleItem  `gorm:"foreignKey:TicketID"`
	Transactions []Transaction `gorm:"foreignKey:TicketID"`

	TimeStamp
}
