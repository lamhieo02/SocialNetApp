package authenandpoststorage

import (
	"gorm.io/gorm"
)

type authenAndPostStorage struct {
	db *gorm.DB
}

func NewAuthenticateAndPostStorage(db *gorm.DB) *authenAndPostStorage {
	return &authenAndPostStorage{db: db}
}
