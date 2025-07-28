package entity

import "github.com/google/uuid"

type StudentAmbassador struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"stu_am_id"`
	Name        string    `gorm:"not null" json:"stu_am_name"`
	ReferalCode string    `gorm:"not null" json:"stu_am_referal_code"`
	Discount    float64   `gorm:"type:numeric(10,2)" json:"stu_am_discount"`
	MaxReferal  int       `json:"stu_am_max_referal"`

	TimeStamp
}
