package database

import (
	"micro-ocr-service/database/seeders"
	"micro-ocr-service/pkg/entities"

	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{},
		&entities.MerchantCategoryCode{},
		&entities.Merchant{},
	)

	// SeedMerchantCategoryCodes()
}

func MigrateDown(db *gorm.DB) {
	db.Migrator().DropTable(
		&entities.User{},
		&entities.MerchantCategoryCode{},
		&entities.Merchant{},
	)
}

func SeedAll(db *gorm.DB) {

	// list of seeder tables
	seeders.MerchantCategoryCodeSeeder(db)
	seeders.UserSeeder(db)
	seeders.MerchantSeeder(db)

}
