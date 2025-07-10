package migrations

import (
	"github.com/Amierza/TedXBackend/entity"
	"gorm.io/gorm"
)

func Rollback(db *gorm.DB) error {
	tables := []interface{}{
		&entity.BundleItem{},
		&entity.Ticket{},
		&entity.Bundle{},

		&entity.Merch{},
		&entity.MerchImage{},
		&entity.MerchImageDetail{},
		&entity.MerchColor{},
		&entity.MerchColorDetail{},
		&entity.MerchSize{},
		&entity.MerchSizeDetail{},

		&entity.GuestAttendance{},
		&entity.TicketForm{},
		&entity.Transaction{},
		&entity.User{},

		&entity.Speaker{},
		&entity.Sponsorship{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}
