package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect_DB() (db *gorm.DB, err error) {

	db, err = gorm.Open(postgres.Open("host=localhost user=postgres password=Josh@123 dbname=ecommerce port=5432 sslmode=disable"), &gorm.Config{})
	return
}
