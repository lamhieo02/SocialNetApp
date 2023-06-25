package authenandpoststorage

import (
	"context"

	"github.com/lamhieo02/socialnetapp/internal/entities"
	"gorm.io/gorm"
)

type authenAndPostStorage struct {
	db *gorm.DB
}

func NewAuthenticateAndPostStorage(db *gorm.DB) *authenAndPostStorage {
	return &authenAndPostStorage{db: db}
}

func (storage *authenAndPostStorage) CreateUser(ctx context.Context, user *entities.UserRegister) (*uint, error) {
	db := storage.db.Begin()

	if err := db.Table(user.TableName()).Create(user).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, err
	}
	return &user.ID, nil
}