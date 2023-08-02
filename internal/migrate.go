package internal

import (
	"github.com/GauravMakhijani/ecommerce-products/internal/domain"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.Product{})
	return err
}
