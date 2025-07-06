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

		&entity.MerchSizeDetail{},
		&entity.MerchSize{},
		&entity.MerchColorDetail{},
		&entity.MerchColor{},
		&entity.MerchImageDetail{},
		&entity.MerchImage{},
		&entity.Merch{},
		&entity.MerchCategory{},

		&entity.Bundle{},
		&entity.Ticket{},
		&entity.BundleItem{},
	); err != nil {
		return err
	}

	return nil
}
