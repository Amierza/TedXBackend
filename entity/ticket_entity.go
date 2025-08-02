package entity

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Type        TicketType `gorm:"default:'main-event'" json:"type"`
	Price       float64    `gorm:"not null;default:0" json:"price"`
	Image       string     `gorm:"not null" json:"image"`
	Quota       int        `gorm:"not null;default:0" json:"quota"`
	Description string     `json:"description"`
	EventDate   time.Time  `gorm:"not null" json:"event_date"`

	Transactions []Transaction `gorm:"foreignKey:TicketID"`

	TimeStamp
}
