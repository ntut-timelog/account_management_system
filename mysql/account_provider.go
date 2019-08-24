package mysql

import (
	"github.com/jinzhu/gorm"

	"github.com/SSL/account_management_system/model"
)

type AccountProvider interface {
	Register(user *model.User) error
}

type accountProvider struct {
	db *gorm.DB
}

func NewAccountProvider(db *gorm.DB) AccountProvider {
	return &accountProvider{
		db: db,
	}
}

func (provider *accountProvider) Register(user *model.User) error {
	return provider.db.Create(user).Error
}
