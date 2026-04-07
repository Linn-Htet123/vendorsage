package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"vendorsage/backend/internal/config"
	"vendorsage/backend/internal/models"
)

func Connect(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Email{},
		&models.Expense{},
		&models.SaaSSubscription{},
	)
	if err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}

	createIndexes(db)
}

func createIndexes(db *gorm.DB) {
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_expenses_category ON expenses(category)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_expenses_transaction_date ON expenses(transaction_date)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_saas_product_name ON saas_subscriptions(product_name)`)
}
