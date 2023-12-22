package userservice

import (
	gormpostgre "github.com/aryamaahendra/mesiakad/domains/account/gorm_postgre"
	"gorm.io/gorm"
)

type ServiceConfig func(os *UserService) error

func WithGormPostgeReposiotry(db *gorm.DB) ServiceConfig {
	return func(os *UserService) error {
		os.account = gormpostgre.New(db)
		return nil
	}
}
