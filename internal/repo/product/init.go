package product

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	// MIGRATE THE SCHEMA
	db.AutoMigrate(&Product{})

	return Repository{
		db: db,
	}
}
