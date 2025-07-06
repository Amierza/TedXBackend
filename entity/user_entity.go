package entity

import (
	"errors"

	"github.com/Amierza/TedXBackend/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	Email       string    `gorm:"unique;not null" json:"user_email"`
	Password    string    `gorm:"not null" json:"user_password"`
	PhoneNumber string    `gorm:"not null" json:"user_phone_number"`
	Role        Role      `gorm:"not null;default:'guest'" json:"user_role"`

	GuestAttendances []GuestAttendance `gorm:"foreignKey:CheckedBy"`

	TimeStamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

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
