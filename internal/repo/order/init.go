package order

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	// MIGRATE THE SCHEMA
	db.AutoMigrate(
		&Order{},
		&ProductOrder{},
	)

	return Repository{
		db: db,
	}
}
