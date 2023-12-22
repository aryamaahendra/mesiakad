package gormpostgre

import (
	"gorm.io/gorm"
)

type GormPostgreRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GormPostgreRepository {
	return &GormPostgreRepository{db: db}
}
