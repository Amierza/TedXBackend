package entity

import (
	"errors"
	"time"

	"github.com/Amierza/TedXBackend/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"user_id"`
	Name        string     `gorm:"not null" json:"user_name"`
	Email       string     `gorm:"unique;not null" json:"user_email"`
	VerifiedAt  *time.Time `json:"verified_at"`
	Password    string     `gorm:"not null" json:"user_password"`
	Image       string     `json:"user_image"`
	PhoneNumber string     `gorm:"not null" json:"user_phone_number"`
	Role        Role       `gorm:"not null;default:'guest'" json:"user_role"`

	GuestAttendances []GuestAttendance `gorm:"foreignKey:CheckedBy"`
	Accounts         []Account         `gorm:"foreignKey:UserID"`
	Sessions         []Session         `gorm:"foreignKey:UserID"`
	Transactions     []Transaction     `gorm:"foreignKey:UserID"`

	TimeStamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error

	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.PhoneNumber, err = helpers.StandardizePhoneNumber(u.PhoneNumber)
	if err != nil {
		return err
	}

	if !IsValidRole(u.Role) {
		return errors.New("invalid user role")
	}

	return nil
}
