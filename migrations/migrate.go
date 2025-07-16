package migrations

import (
	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Sponsorship{},
		&entity.Speaker{},

		&entity.User{},
		&entity.Transaction{},
		&entity.TicketForm{},
		&entity.GuestAttendance{},

		&entity.MerchImage{},
		&entity.Merch{},

		&entity.Bundle{},
		&entity.Ticket{},
		&entity.BundleItem{},
	); err != nil {
		return err
	}

	return nil
}
