package entity

import "github.com/google/uuid"

type Speaker struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"speaker_id"`
	Name  string    `gorm:"not null" json:"speaker_name"`
	Image string    `gorm:"not null" json:"speaker_image"`

	TimeStamp
}
