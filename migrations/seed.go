package migrations

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// err := SeedFromJSON[entity.MerchCategory](db, "./migrations/json/merch_categories.json", entity.MerchCategory{}, "Name")
	// if err != nil {
	// 	return err
	// }

	return nil
}
